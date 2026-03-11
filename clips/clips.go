package clips

import (
	"fmt"
	"log"

	"github.com/DYankee/godacity/internal"
)

type ClipInfo struct {
	Track int     `json:"track"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Color int     `json:"color"`
	Name  string  `json:"name"`
}

type CommandExecutor interface {
	ExecCommand(command string) (*internal.Response, error)
}

type ClipService struct {
	api CommandExecutor
}

func NewClipService(api CommandExecutor) *ClipService {
	return &ClipService{api: api}
}

// Get all clips in a project
func (s *ClipService) GetClips() ([]ClipInfo, error) {
	// Send command
	cmd := `GetInfo: Format="JSON" Type="Clips"`
	res, err := s.api.ExecCommand(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command GetClips: %w", err)
	}

	// log response to console
	log.Println("Get Info result")
	log.Println(res)

	// Unmarshal and return response
	var clips []ClipInfo
	err = res.Unmarshal(&clips)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command GetClips: %w", err)
	}
	return clips, nil
}
