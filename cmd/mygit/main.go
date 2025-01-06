package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Usage: mygit.sh <command> <arg1> <arg2> ...
func main() {
	fmt.Fprintf(os.Stderr, "Programming is Running\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		// e.g : "mygit init"
		for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			}
		}

		headFileContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		}

		fmt.Println("Initialized git directory")

	case "cat-file":
		// construct file path with arguments e.g : "mygit cat-file -p a1g2h345h6"
		filePath := fmt.Sprintf(".git/objects/%s/%s", os.Args[3][:2], os.Args[3][2:])

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening .git/objects file: %s\n", err)
			return
		}
		defer file.Close()

		// zlib reader
		r, err := zlib.NewReader(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating zlib reader: %s\n", err)
			return
		}
		defer r.Close()

		// read all data from zlib reader
		w, err := io.ReadAll(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading zlib reader: %s\n", err)
			return
		}

		// split data by null byte and print second half
		parts := bytes.Split(w, []byte("\x00"))
		if len(parts) > 1 {
			fmt.Print(string(parts[1]))
		} else {
			fmt.Fprintln(os.Stderr, "Error: invalid object")
		}

	case "hash-object":
		// read file contents e.g : "mygit hash-object -w file.txt"
		file, err := os.ReadFile(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
			return
		}
		fmt.Fprintf(os.Stderr, "File contents: %s\n", file)

		// get file stats for file size
		stats, err := os.Stat(os.Args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting file stats: %s\n", err)
			return
		}

		// content with header (blob size + null byte + file contents)
		content := string(file)
		contentAndHeader := fmt.Sprintf("blob %d\x00%s", stats.Size(), content)

		// SHA-1 hash
		sha := sha1.Sum([]byte(contentAndHeader))
		hash := fmt.Sprintf("%x", sha)

		// creating object path based on hash (like first two characters as folder)
		blobName := []rune(hash)
		blobPath := ".git/objects/"
		for i, v := range blobName {
			blobPath += string(v)
			if i == 1 {
				blobPath += "/"
			}
		}

		// compress content using zlib
		var buffer bytes.Buffer
		z := zlib.NewWriter(&buffer)
		_, err = z.Write([]byte(contentAndHeader))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing zlib: %s\n", err)
			return
		}
		z.Close()

		// create directory structure
		if err := os.MkdirAll(filepath.Dir(blobPath), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			return
		}

		// create file and write compressed content
		f, err := os.Create(blobPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %s\n", err)
			return
		}
		defer f.Close()
		_, err = f.Write(buffer.Bytes())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
			return
		}

		fmt.Print(hash)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
