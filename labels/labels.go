package labels

import (
	"fmt"

	"github.com/DYankee/godacity/internal"
)

type CommandExecutor interface {
	ExecCommand(command string) (*internal.Response, error)
}

type LabelService struct {
	api CommandExecutor
}

func NewLabelService(api CommandExecutor) *LabelService {
	return &LabelService{api: api}
}

// Add calls the AddLabel command using any executor (like your Audacity client)
func (s *LabelService) Add() error {
	_, err := s.api.ExecCommand("AddLabel:")
	if err != nil {
		return fmt.Errorf("labels.Add: %w", err)
	}
	return nil
}

func (s *LabelService) Set(id int, text string) error {
	cmd := fmt.Sprintf(`SetLabel: Label="%d" Text="%s"`, id, text)
	_, err := s.api.ExecCommand(cmd)
	return err
}
