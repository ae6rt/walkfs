package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	root = flag.String("root-dir", ".", "root path")
)

func main() {
	flag.Parse()
	if err := os.Chdir(*root); err != nil {
		log.Fatalf("%v\n", err)
	}
	files := make([]string, 0)
	markFn := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && strings.Count(path, "/") > 1 {
			return filepath.SkipDir
		}
		if err != nil {
			return err
		}
		if info.Name() == "config.xml" {
			files = append(files, path)
		}
		return nil
	}
	err := filepath.Walk(".", markFn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("files: %d\n", len(files))
	fmt.Printf("files: %v\n", (files))
}
