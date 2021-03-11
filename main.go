package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
)

//go:embed public/*
var embeddedFs embed.FS

const dirPrefix = "public"
const indexFile = "/index.html"

type loadedFile struct {
	file []byte
	mime string
}

func loadFilesFromEmbeddedFs() (map[string][]byte, error) {
	var files = make(map[string][]byte)

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
		files[path] = file
		fmt.Printf("Loading file from embeded filessystem. file %s\n", path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	files, err := loadFilesFromEmbeddedFs()

	if err != nil {
		log.Fatalln("Could not load files from embedded filesystem. err: ", err)
	}

	indexFile, indexFileFound := files[indexFile]
	if !indexFileFound {
		log.Fatalln("Could not find index.html")
	}

	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			file, exists := files[req.URL.Path]
			if !exists {
				file = indexFile
			}

			_, err = w.Write(file)
			if err != nil {
				log.Printf("Could not send file: {%s} to client", req.URL.Path)
			}
		}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println(srv.ListenAndServe())
}
