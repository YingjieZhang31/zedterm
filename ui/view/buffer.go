package view

import (
	"bufio"
	"os"
)

type buffer struct {
	lines []string

	fileName string
}

func newBuffer() *buffer {
	return &buffer{
		lines: make([]string, 0),
	}
}

func (b *buffer) loadFile(fileName string) error {
	if fileName == "" {
		return nil
	}
	b.fileName = fileName
	fd, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		b.lines = append(b.lines, scanner.Text())
	}

	return nil
}
