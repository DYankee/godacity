//go:build integration

package godacity_test

import "testing"

func TestLabels(t *testing.T) {
	// Setup: Select a region first
	err := a.Tracks.SelectRegion(0.0, 0.5)
	if err != nil {
		t.Fatalf("setup selection failed: %v", err)
	}

	t.Run("AddLabel", func(t *testing.T) {
		err := a.Labels.Add()
		if err != nil {
			t.Errorf("Labels.Add failed: %v", err)
		}
	})

	t.Run("SetLabel", func(t *testing.T) {
		err := a.Labels.Set(0, "Refactored-Label")
		if err != nil {
			t.Errorf("Labels.Set failed: %v", err)
		}
	})

	t.Run("SetLabel_InvalidIndex", func(t *testing.T) {
		err := a.Labels.Set(999, "fail")
		if err == nil {
			t.Error("expected error for invalid label index")
		}
	})
}
