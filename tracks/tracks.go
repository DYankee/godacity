package tracks

import (
	"fmt"

	"github.com/DYankee/godacity/internal"
)

type CommandExecutor interface {
	ExecCommand(command string) (*internal.Response, error)
}

type TrackService struct {
	api CommandExecutor
}

func NewTrackService(api CommandExecutor) *TrackService {
	return &TrackService{api: api}
}

func (s *TrackService) SplitNew() error {
	_, err := s.api.ExecCommand("SplitNew:")
	if err != nil {
		return fmt.Errorf("split: %w", err)
	}
	return nil
}

// Select a region
func (s *TrackService) SelectRegion(startTime float64, endTime float64) error {
	// Check function parameters are valid
	if endTime <= startTime {
		return fmt.Errorf("selectRegion: startTime must be greater than endTime")
	}

	// Format command
	cmd := fmt.Sprintf(`Select: End="%f" RelativeTo="ProjectStart" Start="%f"`, endTime, startTime)
	_, err := s.api.ExecCommand(cmd)
	if err != nil {
		return fmt.Errorf("select region: %w", err)
	}
	return nil
}
