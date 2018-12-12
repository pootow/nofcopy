package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

// GetSHA1Hash compute sha1 hash of a file
func GetSHA1Hash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func StringSliceChunk(slice []string, chunkSize int) [][]string {
	var divided [][]string

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		divided = append(divided, slice[i:end])
	}
	return divided
}
