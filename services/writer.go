package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Writer struct {
	path string
}

func NewWriter(path string) *Writer {
	//path := "./path/to/fileOrDir"
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Error opening file: %s", err)
		os.Exit(1)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// IsDir is short for fileInfo.Mode().IsDir()
	if !fileInfo.IsDir() {
		panic(fmt.Sprintf(" %s is not a directory", path))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return &Writer{path: path}
}

func (w *Writer) WriteToFile(content string, filename string) {
	full := w.path + "/" + filename
	f, err := os.Create(full)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	log.Debug(l, "bytes written successfully")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (w *Writer) DeleteFilesInFolder() {
	// Open the directory
	dir, err := os.Open(w.path)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Read all files in the directory
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the files and delete each one
	for _, file := range files {
		err := os.Remove(w.path + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}
