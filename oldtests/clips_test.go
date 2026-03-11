//go:build integration

package godacity_test

import "testing"

func TestClips(t *testing.T) {
	t.Run("GetClips", testGetClips)
	t.Run("SelectRegion", testSelectRegion)
	t.Run("SelectRegion_InvalidRange", testSelectRegionInvalidRange)
	t.Run("SelectRegion_EqualTimes", testSelectRegionEqualTimes)
	t.Run("SplitNew", testSplitNew)
}

func testGetClips(t *testing.T) {
	clips, err := a.GetClips()
	if err != nil {
		t.Fatalf("GetClips failed: %s", err)
	}
	if len(clips) == 0 {
		t.Fatal("expected at least one clip after importing audio")
	}
	t.Logf("clips: %+v", clips)
}

func testSelectRegion(t *testing.T) {
	err := a.SelectRegion(0.0, 1.0)
	if err != nil {
		t.Fatalf("SelectRegion failed: %s", err)
	}
}

func testSelectRegionInvalidRange(t *testing.T) {
	err := a.SelectRegion(3.0, 1.0)
	if err == nil {
		t.Fatal("expected error when startTime > endTime, got nil")
	}
}

func testSelectRegionEqualTimes(t *testing.T) {
	err := a.SelectRegion(1.0, 1.0)
	if err == nil {
		t.Fatal("expected error when startTime == endTime, got nil")
	}
}

func testSplitNew(t *testing.T) {
	err := a.SelectRegion(0.5, 1.0)
	if err != nil {
		t.Fatalf("SelectRegion (setup for Split) failed: %s", err)
	}

	err = a.SplitNew()
	if err != nil {
		t.Fatalf("Split failed: %s", err)
	}

	clips, err := a.GetClips()
	if err != nil {
		t.Fatalf("GetClips after Split failed: %s", err)
	}
	if len(clips) < 2 {
		t.Errorf("expected at least 2 clips after split, got %d", len(clips))
	}
	t.Logf("clips after split: %+v", clips)
}
