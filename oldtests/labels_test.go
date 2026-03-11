//go:build integration

package godacity_test

import "testing"

func TestLabels(t *testing.T) {
	t.Run("AddLabel", testAddLabel)
	t.Run("AddLabel_Multiple", testAddLabelMultiple)
	t.Run("SetLabel", testSetLabel)
	t.Run("SetLabel_InvalidIndex", testSetLabelInvalidIndex)
}

func testAddLabel(t *testing.T) {
	err := a.SelectRegion(0.0, 0.5)
	if err != nil {
		t.Fatalf("SelectRegion (setup) failed: %s", err)
	}

	err = a.AddLabel()
	if err != nil {
		t.Fatalf("AddLabel failed: %s", err)
	}
}

func testAddLabelMultiple(t *testing.T) {
	positions := []struct {
		start float64
		end   float64
	}{
		{0.5, 1.0},
		{1.0, 1.5},
	}

	for i, pos := range positions {
		err := a.SelectRegion(pos.start, pos.end)
		if err != nil {
			t.Fatalf("SelectRegion for label %d failed: %s", i, err)
		}
		err = a.AddLabel()
		if err != nil {
			t.Fatalf("AddLabel %d failed: %s", i, err)
		}
	}
}

func testSetLabel(t *testing.T) {
	err := a.SelectRegion(0.0, 0.5)
	if err != nil {
		t.Fatalf("SelectRegion (setup) failed: %s", err)
	}

	err = a.AddLabel()
	if err != nil {
		t.Fatalf("AddLabel (setup) failed: %s", err)
	}

	err = a.SetLabel(0, "test-label")
	if err != nil {
		t.Fatalf("SetLabel failed: %s", err)
	}
}

func testSetLabelInvalidIndex(t *testing.T) {
	err := a.SetLabel(9999, "should-fail")
	if err == nil {
		t.Fatal("expected error for invalid label index, got nil")
	}
}
