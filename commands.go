package godacity

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DYankee/godacity/internal"
)

// ---------------------------------------------------------------- \\
// This file contains the function for executing all cmds
// along with any commands that don't fit in another file.
// ---------------------------------------------------------------- \\

// Execute a command and return the response
func (a *Audacity) ExecCommand(command string) (*internal.Response, error) {
	// lock pipes to prevent concurrent access
	a.mu.Lock()
	defer a.mu.Unlock()

	// Set a read deadline on the pipe
	a.conn.Receive.SetReadDeadline(
		time.Now().Add(10 * time.Second),
	)
	defer a.conn.Receive.SetReadDeadline(time.Time{})

	//send command
	log.Println("Send: >>> \n" + command)
	_, err := fmt.Fprint(a.conn.Send, command+a.platformConfig.EOL)
	if err != nil {
		return nil, fmt.Errorf("failed to write to pipe: %w", err)
	}

	//get response
	var sb strings.Builder
	for a.conn.Scanner.Scan() {
		text := a.conn.Scanner.Text()
		sb.WriteString(text + "\n")
		// Check if text was the last line
		if strings.HasPrefix(text, "BatchCommand finished:") {
			break
		}
	}
	if err := a.conn.Scanner.Err(); err != nil {
		return nil, fmt.Errorf("pipe read error: %w", err)
	}
	var response internal.Response
	response.Raw = sb.String()
	if !response.OK() {
		return nil, fmt.Errorf("Audacity error: %s", response.Raw)
	}

	return &response, nil
}
