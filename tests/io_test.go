//go:build integration

package godacity_test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIO(t *testing.T) {
	t.Run("ImportAudio", func(t *testing.T) {
		err := a.IO.ImportAudio(testAudioPath(t))
		if err != nil {
			t.Fatalf("IO.ImportAudio failed: %v", err)
		}
	})

	t.Run("ExportAudio", func(t *testing.T) {
		tmpDir := t.TempDir()
		err := a.IO.ExportAudio(tmpDir, "test_out.wav")
		if err != nil {
			t.Fatalf("IO.ExportAudio failed: %v", err)
		}

		if _, err := os.Stat(filepath.Join(tmpDir, "test_out.wav")); os.IsNotExist(err) {
			t.Error("exported file does not exist")
		}
	})
}
