//go:build integration

package godacity_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DYankee/godacity"
)

var a *godacity.Audacity

func TestMain(m *testing.M) {
	var err error
	a, err = godacity.NewAudacity(&godacity.Config{
		AutoStart:    true,
		StartTimeout: 15 * time.Second,
	})
	if err != nil {
		panic("failed to connect to audacity: " + err.Error())
	}

	// Import fixture so every test starts with audio present
	absPath, _ := filepath.Abs("testdata/test.wav")
	err = a.IO.ImportAudio(absPath)
	if err != nil {
		fmt.Printf("Warning: setup import failed: %v\n", err)
	}

	code := m.Run()

	a.ClosePipes()
	os.Exit(code)
}

func testAudioPath(t *testing.T) string {
	t.Helper()
	absPath, err := filepath.Abs("testdata/test.wav")
	if err != nil {
		t.Fatalf("failed to resolve test audio path: %s", err)
	}
	return absPath
}
