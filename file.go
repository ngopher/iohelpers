package iohelpers

import (
	"bufio"
	"io"
	"os"
)

// File is a wrapper of os.File.
type File struct {
	*os.File
}

func NewFile(file *os.File) *File {
	return &File{file}
}

// ReadLines reads the whole file into memory and returns a slice of its lines.
func (file *File) ReadLines() ([]string, error) {
	reader := bufio.NewReader(file)
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
