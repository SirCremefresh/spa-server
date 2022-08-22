package main

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//go:embed public/*
var embeddedFs embed.FS

const dirPrefix = "public"
const indexFileName = "/index.html"
const configFileName = "/config.json"

type loadedFile struct {
	file []byte
	mime string
}

func getenvString(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getenvUint(key string, fallback uint64) uint64 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	valueUint, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatalf("Could not convert value from env to uint64. key: %s, value: %s err: %v", key, value, err)
	}
	return valueUint
}

func loadFilesFromEmbeddedFs() (map[string]loadedFile, error) {
	var files = make(map[string]loadedFile)

	err := fs.WalkDir(embeddedFs, dirPrefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		file, err := embeddedFs.ReadFile(path)
		if err != nil {
			return err
		}
		path = strings.TrimLeft(path, dirPrefix)
		// apply base href correction
		if path == indexFileName {
			file = bytes.Replace(
				file,
				[]byte("<base href=\"/\""),
				[]byte(fmt.Sprint("<base href=\"", getenvString("BASE_HREF", "/"), "\"")),
				-1)
		}
		files[path] = loadedFile{
			file: file,
			mime: mime.TypeByExtension(filepath.Ext(path)),
		}
		log.Printf("Loading file from embeded filessystem. file %s\n", path)
		return nil
	})

	files[configFileName] = loadedFile{
		file: []byte(getenvString("CONFIG_JSON", "{}")),
		mime: mime.TypeByExtension(filepath.Ext(configFileName)),
	}

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	port := getenvString("PORT", "8080")
	addr := getenvString("ADDRESS", "0.0.0.0")

	readTimeout := getenvUint("READ_TIMEOUT_SECONDS", 5)
	writeTimeout := getenvUint("WRITE_TIMEOUT_SECONDS", 10)
	idleTimeout := getenvUint("IDLE_TIMEOUT_SECONDS", 120)

	files, err := loadFilesFromEmbeddedFs()

	if err != nil {
		log.Fatalf("Could not load files from embedded filesystem. err: %v", err)
	}

	indexFile, indexFileFound := files[indexFileName]
	if !indexFileFound {
		log.Fatalln("Could not find index.html")
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", addr, port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			loadedFile, exists := files[req.URL.Path]
			if !exists {
				loadedFile = indexFile
			}
			w.Header().Add("Content-Type", loadedFile.mime)
			switch req.URL.Path {
			case indexFileName:
				fallthrough
			case configFileName:
				w.Header().Add("Cache-Control", "public, max-age: 60") // refresh every 1 minute to ensure fresh-ness
			default:
				w.Header().Add("Cache-Control", "public, max-age: 604800, immutable")
			}

			_, err = w.Write(loadedFile.file)
			if err != nil {
				log.Printf("Could not send loadedFile to client. file: %s", req.URL.Path)
			}
		}),
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	log.Printf("Starting server on Addr: %s:%s", addr, port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server. err: %v", err)
	}

	log.Println("Stopping Server")
}
