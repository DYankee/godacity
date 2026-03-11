//go:build integration

package godacity_test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIO(t *testing.T) {
	t.Run("ImportAudio", testImportAudio)
	t.Run("ImportAudio_InvalidPath", testImportAudioInvalidPath)
	t.Run("ExportAudio", testExportAudio)
}

func testImportAudio(t *testing.T) {
	err := a.ImportAudio(testAudioPath(t))
	if err != nil {
		t.Fatalf("ImportAudio failed: %s", err)
	}

	clips, err := a.GetClips()
	if err != nil {
		t.Fatalf("GetClips after ImportAudio failed: %s", err)
	}
	if len(clips) == 0 {
		t.Fatal("expected at least one clip after importing audio")
	}
	t.Logf("clips after import: %+v", clips)
}

func testImportAudioInvalidPath(t *testing.T) {
	err := a.ImportAudio("/nonexistent/path/fake.wav")
	if err == nil {
		t.Fatal("expected error for invalid file path, got nil")
	}
}

func testExportAudio(t *testing.T) {
	tmpDir := t.TempDir()
	err := a.ExportAudio(tmpDir, "export_test.wav")
	if err != nil {
		t.Fatalf("ExportAudio failed: %s", err)
	}

	exportPath := filepath.Join(tmpDir, "export_test.wav")
	info, err := os.Stat(exportPath)
	if err != nil {
		t.Fatalf("exported file not found: %s", err)
	}
	if info.Size() == 0 {
		t.Error("exported file is empty")
	}
	t.Logf("exported file size: %d bytes", info.Size())
}
