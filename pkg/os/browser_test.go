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

// TestOpenBrowserLinux tests that OpenBrowser opens a browser
func TestOpenBrowserLinux(t *testing.T) {
	// not resetting mock args because it should be thread safe
	// test.ResetMockArgs()
	uniqueId := "TestOpenBrowserLinux"
	myExec = test.CreateFakeExecCommand(uniqueId)
	myRuntime = "linux"
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

// TestOpenBrowserWindows tests that OpenBrowser opens a browser
func TestOpenBrowserWindows(t *testing.T) {
	// not resetting mock args because it should be thread safe
	// test.ResetMockArgs()
	uniqueId := "TestOpenBrowserWindows"
	myExec = test.CreateFakeExecCommand(uniqueId)
	myRuntime = "windows"
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

// TestOpenBrowserDarwin tests that OpenBrowser opens a browser
func TestOpenBrowserDarwin(t *testing.T) {
	// not resetting mock args because it should be thread safe
	// test.ResetMockArgs()
	uniqueId := "TestOpenBrowserDarwin"
	myExec = test.CreateFakeExecCommand(uniqueId)
	myRuntime = "darwin"
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

// TestOpenBrowserUnsupported tests that OpenBrowser returns an error
func TestOpenBrowserUnsupported(t *testing.T) {
	// not resetting mock args because it should be thread safe
	// test.ResetMockArgs()
	uniqueId := "TestOpenBrowserUnsupported"
	myExec = test.CreateFakeExecCommand(uniqueId)
	myRuntime = "unsupported"
	defer func() { myExec = exec.Command }()
	url := "https://www.google.com"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("OpenBrowser(%s) didn't panic", url)
		}
	}()

	OpenBrowser(url)
}

// TestHelperProcess is a helper process for FakeExecCommand
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
}
