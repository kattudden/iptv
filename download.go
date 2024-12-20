package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download(url string, filepath string) error {
	createFolderIfNotExists(filepath)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Download of channelfile failed!")
	}

	_, err = io.Copy(out, resp.Body)
	return nil
}

func createFolderIfNotExists(filePath string) {
	folderPath := filepath.Dir(filePath)
	if folderPath == "." {
		return
	}
	err := os.MkdirAll(folderPath, 0774)
	if err != nil {
		fmt.Println(err)
		return
	}
}
