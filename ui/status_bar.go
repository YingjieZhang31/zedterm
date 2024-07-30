package ui

import (
	"fmt"

	"zedterm/doc_status"
	"zedterm/terminal"
)

const (
	statusBarFmtStr = "file: %s | line: %d, col: %d"
)

type StatusBar struct {
	currentStatus *doc_status.DocStatus
}

func NewStatusBar() *StatusBar {
	return &StatusBar{}
}

func (m *StatusBar) UpdateDocStatus(s *doc_status.DocStatus) {
	m.currentStatus = s
}

func (s *StatusBar) Render() {
	_, windowHeight := terminal.Size()
	terminal.PrintLine(windowHeight-1, fmt.Sprintf(statusBarFmtStr, s.currentStatus.FileName, s.currentStatus.TextLocY, s.currentStatus.TextLocX))
}
