package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/h2non/bimg"
)

func main() {
	// Define the flag: -path="your/path" (defaults to current directory)
	rootPath := flag.String("path", ".", "The directory path to scan for images")
	flag.Parse()

	start := time.Now()
	count := 0

	fmt.Printf("Scanning directory: %s\n", *rootPath)

	err := filepath.Walk(*rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Filter for .jpg and ignore existing thumbnails
		if !info.IsDir() && isJPG(path) && !isThumb(path) {

			thumbPath := strings.TrimSuffix(path, filepath.Ext(path)) + "_thumb.jpg"

			// Skip if already exists
			if _, err := os.Stat(thumbPath); err == nil {
				return nil
			}

			err := createThumbnail(path, thumbPath)
			if err != nil {
				log.Printf("Error processing %s: %v\n", path, err)
			} else {
				fmt.Printf("✓ Created: %s\n", filepath.Base(thumbPath))
				count++
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Walk failed: %v", err)
	}

	fmt.Printf("\nFinished! Processed %d images in %v\n", count, time.Since(start))
}

func isJPG(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jpg" || ext == ".jpeg"
}

func isThumb(path string) bool {
	return strings.HasSuffix(path, "_thumb.jpg")
}

func createThumbnail(src string, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	imageBuf, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// bimg.Thumbnail(800) automatically performs a smart center-crop
	thumbnail, err := bimg.NewImage(imageBuf).Thumbnail(800)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, thumbnail, 0664)
}
