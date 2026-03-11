//go:build integration

package godacity_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/DYankee/godacity"
)

var a *godacity.Audacity

// testAudioPath returns the absolute path to the shared test fixture.
func testAudioPath(t *testing.T) string {
	t.Helper()
	absPath, err := filepath.Abs("testdata/test.wav")
	if err != nil {
		t.Fatalf("failed to resolve test audio path: %s", err)
	}
	return absPath
}

// importTestAudio imports the shared fixture into the current project.
func importTestAudio(t *testing.T) {
	t.Helper()
	_, err := a.ExecCommand(
		`Import2: Filename="` + testAudioPath(t) + `"`,
	)
	if err != nil {
		t.Fatalf("failed to import test audio: %s", err)
	}
}

func TestMain(m *testing.M) {
	var err error
	a, err = godacity.NewAudacity(&godacity.Config{
		AutoStart:    true,
		StartTimeout: 10 * time.Second,
	})
	if err != nil {
		panic("failed to connect to audacity: " + err.Error())
	}

	// Import fixture so every test starts with audio present
	absPath, _ := filepath.Abs("testdata/test.wav")
	_, err = a.ExecCommand(`Import2: Filename="` + absPath + `"`)
	if err != nil {
		panic("failed to import test audio: " + err.Error())
	}

	code := m.Run()

	a.ClosePipes()
	os.Exit(code)
}
