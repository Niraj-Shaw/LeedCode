package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()
	path = filepath.Dir(path)
	count := 0
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				count++
			}
		}
	}
	fmt.Print(count)
}
