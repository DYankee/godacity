package io

import (
	"fmt"

	"github.com/DYankee/godacity/internal"
)

// ---------------------------------------------------------------- \\
// This file contains all functions relating to importing and
// exporting audio.
// ---------------------------------------------------------------- \\

type CommandExecutor interface {
	ExecCommand(command string) (*internal.Response, error)
}

type IOService struct {
	api CommandExecutor
}

func NewIOService(api CommandExecutor) *IOService {
	return &IOService{api: api}
}

// Load an audio file into the current project
func (s *IOService) ImportAudio(filePath string) error {
	cmd := fmt.Sprintf(`Import2: Filename="%s"`, filePath)
	_, err := s.api.ExecCommand(cmd)
	if err != nil {
		return fmt.Errorf("ImportAudio: %w", err)
	}
	return nil
}

// Export selected audio
func (s *IOService) ExportAudio(destination string, fileName string) error {
	cmd := fmt.Sprintf(`Export2: Filename="%s/%s" NumChannels="2"`, destination, fileName)
	_, err := s.api.ExecCommand(cmd)
	if err != nil {
		return fmt.Errorf("ExportAudio: %w", err)
	}
	return nil
}
