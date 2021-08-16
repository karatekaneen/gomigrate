package main

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

func main() {
	readzip("ansokan3.zip")
}

func readzip(archive string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	for _, file := range reader.File {

		if file.FileInfo().IsDir() {
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		contentBuf := new(strings.Builder)

		_, err = io.Copy(contentBuf, fileReader)
		if err != nil {
			return err
		}

		targetPath := getTargetPath(file.Name)
		fmt.Println("*********")
		fmt.Printf("Original path: %s \t Target path: %s\n", file.Name, targetPath)
		fmt.Printf("Content: %s \n", contentBuf.String())
		fmt.Println("*********")

	}

	return nil
}

func getTargetPath(fullPath string) string {
	parts := strings.Split(fullPath, "/")
	return fmt.Sprintf("%s/%s", parts[0], parts[2])
}
