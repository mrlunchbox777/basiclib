/**
 * Author: Andrew Shoell
 * File: browser_test.go
 */

package os

import (
	"os"
	"os/exec"
	"testing"

	test "github.com/mrlunchbox777/basiclib/pkg/test"
)

// TestOpenBrowser tests that OpenBrowser opens a browser
func TestOpenBrowser(t *testing.T) {
	// not resetting mock args because it should be thread safe
	// test.ResetMockArgs()
	uniqueId := "TestOpenBrowser"
	myExec = test.CreateFakeExecCommand(uniqueId)
	defer func() { myExec = exec.Command }()
	url := "https://www.google.com"
	OpenBrowser(url)
	got := test.GetMockArgsByUniqueId(uniqueId)
	if len(got) != 1 {
		// not resetting mock args because it should be thread safe
		// test.ResetMockArgs()
		t.Errorf("OpenBrowser(%s) didn't call exec.Command", url)
	}
}

// TestHelperProcess is a helper process for FakeExecCommand
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
}
