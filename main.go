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
	fileName = flag.String("filename", "", "file to search for")
	root     = flag.String("root-dir", ".", "root path")
	maxDepth = flag.Int("max-depth", 1, "max depth")
)

func main() {
	flag.Parse()
	if *fileName == "" {
		log.Fatalf("need a filename\n")
	}

	files := make([]string, 0)
	markFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Count(path, "/") > *maxDepth-1 {
			return filepath.SkipDir
		}
		if info.Mode().IsRegular() && info.Name() == *fileName {
			files = append(files, path)
		}
		return nil
	}

	if err := os.Chdir(*root); err != nil {
		log.Fatal(err)
	}

	err := filepath.Walk(".", markFn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("files: %d\n", len(files))
}
