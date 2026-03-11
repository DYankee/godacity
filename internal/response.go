package internal

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Response struct {
	Raw string
}

func (r *Response) OK() bool {
	return strings.Contains(r.Raw, "BatchCommand finished: OK")
}

// This function returns only the JSON portion of the response
func (r *Response) JSON() string {
	lines := strings.Split(strings.TrimSpace(r.Raw), "\n")
	var jsonLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "BatchCommand finished:") {
			break
		}
		jsonLines = append(jsonLines, line)
	}
	return strings.Join(jsonLines, "\n")
}

func (r *Response) Unmarshal(target any) error {
	if !r.OK() {
		return fmt.Errorf("command failed: %s", r.Raw)
	}
	jsonStr := r.JSON()
	if jsonStr == "" {
		return fmt.Errorf("empty JSON in response")
	}
	return json.Unmarshal([]byte(r.JSON()), target)
}
