package internal

import (
	"bufio"
	"fmt"
	"os"
)

type OSInfo struct {
	ToName   string
	FromName string
	EOL      string
}

type Connection struct {
	Send    *os.File
	Receive *os.File
	Scanner *bufio.Scanner
}

// Dial handles the actual OS file opening
func Dial(cfg OSInfo) (*Connection, error) {
	toFile, err := os.OpenFile(cfg.ToName, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return nil, fmt.Errorf("failed to open write pipe: %w", err)
	}

	fromFile, err := os.OpenFile(cfg.FromName, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		toFile.Close()
		return nil, fmt.Errorf("failed to open read pipe: %w", err)
	}

	scanner := bufio.NewScanner(fromFile)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	return &Connection{
		Send:    toFile,
		Receive: fromFile,
		Scanner: scanner,
	}, nil
}
