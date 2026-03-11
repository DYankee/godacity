package godacity

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"github.com/DYankee/godacity/clips"
	"github.com/DYankee/godacity/internal"
	"github.com/DYankee/godacity/io"
	audIO "github.com/DYankee/godacity/io"
	"github.com/DYankee/godacity/labels"
	"github.com/DYankee/godacity/tracks"
)

type Audacity struct {
	platformConfig internal.OSInfo // Use the type from internal
	conn           *internal.Connection
	mu             sync.Mutex
	Labels         *labels.LabelService
	Clips          *clips.ClipService
	IO             *audIO.IOService
	Tracks         *tracks.TrackService
}

// We define the method here so it can access the Audacity struct
func (a *Audacity) Connect() error {
	conn, err := internal.Dial(a.platformConfig)
	if err != nil {
		return err
	}
	a.conn = conn
	return nil
}

func (a *Audacity) ClosePipes() {
	if a.conn != nil {
		a.conn.Send.Close()
		a.conn.Receive.Close()
	}
}

// ... rest of your NewAudacity code ...

type Config struct {
	AutoStart    bool
	StartTimeout time.Duration
	ImportFile   string
}

func NewAudacity(cfg *Config) (*Audacity, error) {
	a := &Audacity{}
	// Set OS specific data
	if runtime.GOOS == "windows" {
		fmt.Println("pipe-test.go, running on windows")
		a.platformConfig.ToName = `\\.\pipe\ToSrvPipe`
		a.platformConfig.FromName = `\\.\pipe\FromSrvPipe`
		a.platformConfig.EOL = "\r\n"
	} else {
		fmt.Println("pipe-test.go, running on linux or mac")
		a.platformConfig.ToName = fmt.Sprintf("/tmp/audacity_script_pipe.to.%d", os.Getuid())
		a.platformConfig.FromName = fmt.Sprintf("/tmp/audacity_script_pipe.from.%d", os.Getuid())
		a.platformConfig.EOL = "\n"
	}

	// Try and connect to audacity
	err := a.Connect()
	if err != nil && cfg.AutoStart {
		println("Cannot connect to audacity. Attempting to start ...")
		err := exec.Command("audacity").Start()
		if err != nil {
			return nil, fmt.Errorf("failed to start audacity: %w", err)
		}

		// Try and connect to audacity once every second for the timeout duration
		deadline := time.Now().Add(cfg.StartTimeout)
		for time.Now().Before(deadline) {
			err = a.Connect()
			if err == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			return nil, fmt.Errorf("audacity started but pipes never appeared: %w", err)
		}
	}
	// init services
	a.Labels = labels.NewLabelService(a)
	a.IO = io.NewIOService(a)
	a.Tracks = tracks.NewTrackService(a)
	a.Clips = clips.NewClipService(a)

	// Import file if specified
	if cfg.ImportFile != "" {
		err = a.IO.ImportAudio(cfg.ImportFile)
		if err != nil {
			return nil, fmt.Errorf("failed to import file on startup: %w", err)
		}
	}
	return a, err
}
