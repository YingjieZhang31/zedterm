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

func (b *buffer) Delete(atRow, atCol int) {
	if len(b.lines) <= atRow {
		return
	}
	text := b.lines[atRow]
	if atCol < len(text) {
		b.lines[atRow] = text[:atCol] + text[atCol+1:]
	}
}

func (b *buffer) insertChar(row, col int, r rune) {
	// insert char at new row
	if len(b.lines) == row {
		b.lines = append(b.lines, string(r))
		return
	}
	// insert char at existing row
	line := b.lines[row]
	b.lines[row] = line[:col] + string(r) + line[col:]
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
