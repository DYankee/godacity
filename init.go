package godacity

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// Information needed to correctly send and receive
// cmds based on operating system
type osInfo struct {
	toName   string
	fromName string
	eol      string
}

type Connection struct {
	send    *os.File
	receive *os.File
}

type Audacity struct {
	osInfo     osInfo
	connection Connection
	Status     bool
}

func (a *Audacity) Init() {
	// Set OS specific data
	if runtime.GOOS == "windows" {
		fmt.Println("pipe-test.go, running on windows")
		a.osInfo.toName = `\\.\pipe\ToSrvPipe`
		a.osInfo.fromName = `\\.\pipe\FromSrvPipe`
		a.osInfo.eol = "\r\n"
	} else {
		fmt.Println("pipe-test.go, running on linux or mac")
		a.osInfo.toName = fmt.Sprintf("/tmp/audacity_script_pipe.to.%d", os.Getuid())
		a.osInfo.fromName = fmt.Sprintf("/tmp/audacity_script_pipe.from.%d", os.Getuid())
		a.osInfo.eol = "\n"
	}

	// Try and connect to audacity
	err := a.Connect()
	if err != nil {
		println("Cannot connect to audacity. Attempting to start ...")
		for i := 0; i < 10; i++ {
			exec.Command("Audacity")
			err := a.Connect()
			if err == nil {
				break
			}
		}
		if err != nil {
			return
		}
	}

}

func (a *Audacity) Connect() error {
	a.Status = false
	toFile, err := os.OpenFile(a.osInfo.toName, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Println("-- Failed to open file to write to:", err)
		a.Status = false
	} else {
		a.Status = true
		log.Println("-- File to write to has been opened")
	}
	a.connection.send = toFile

	fromfile, err := os.OpenFile(a.osInfo.fromName, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Println("-- Failed to open file to read from:", err)
		a.Status = false
	} else {
		a.Status = true
		log.Println("-- File to read from has been opened")
	}
	a.connection.receive = fromfile

	if !a.Status {
		err = errors.New("cannot connect to audacity")
	} else {
		err = nil
	}
	return err
}
