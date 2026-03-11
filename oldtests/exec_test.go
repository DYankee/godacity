//go:build integration

package godacity_test

import "testing"

func TestExec(t *testing.T) {
	t.Run("InvalidCommand", testExecCommandInvalid)
}

func testExecCommandInvalid(t *testing.T) {
	_, err := a.ExecCommand("CompletelyFakeCommand:")
	if err == nil {
		t.Fatal("expected error for invalid command, got nil")
	}
}
