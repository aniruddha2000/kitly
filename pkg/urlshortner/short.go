package urlshortner

import (
	"bufio"
	b64 "encoding/base64"
	"log"
	"os"
	"strings"
)

type Shorter interface {
	Short(string) string
}

type FileStore struct {
	File *os.File
}

func NewFileStore(f *os.File) *FileStore {
	return &FileStore{File: f}
}

func (fs *FileStore) Short(url string) string {
	log.Println("Short called")
	val, ok := fs.find(url)
	if ok {
		return val
	}

	shortenURL := fs.encode(url)
	fs.store(url, shortenURL)
	return shortenURL
}

func (fs *FileStore) find(url string) (string, bool) {
	fileScanner := bufio.NewScanner(fs.File)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedLine := strings.Split(line, "=")
		log.Println(splittedLine)

		if splittedLine[0] == url {
			return splittedLine[1], true
		}
	}
	return "", false
}

func (fs *FileStore) store(key, val string) {
	line := key + "=" + val + "\n"

	n, err := fs.File.WriteString(line)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes written\n", n)
}

func (fs *FileStore) encode(url string) string {
	return b64.StdEncoding.EncodeToString([]byte(url))
}
