//go:build integration

package godacity_test

import "testing"

func TestClipsAndTracks(t *testing.T) {
	t.Run("SelectRegion", func(t *testing.T) {
		err := a.Tracks.SelectRegion(0.1, 0.9)
		if err != nil {
			t.Errorf("Tracks.SelectRegion failed: %v", err)
		}
	})

	t.Run("SplitNew", func(t *testing.T) {
		_ = a.Tracks.SelectRegion(0.2, 0.4)
		err := a.Tracks.SplitNew()
		if err != nil {
			t.Errorf("Tracks.SplitNew failed: %v", err)
		}
	})

	t.Run("GetClips", func(t *testing.T) {
		data, err := a.Clips.GetClips() // Assuming GetClips is in Clips service
		if err != nil {
			t.Errorf("Clips.GetClips failed: %v", err)
		}
		if data == nil {
			t.Error("expected clip data, got nil")
		}
	})
}
