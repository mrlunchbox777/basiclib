/*
 * Author : Andrew Shoell
 * File   : fake_exec_command.go
 */

package test

import (
	"os"
	"os/exec"
)

// FakeExecCommand is a mock exec.Command
type FakeExecCommand struct {
	UniqueId string
	Command  string
	Args     []string
}

// MockArgs is a list of FakeExecCommand.
// This is used to check that the correct commands were run
var MockArgs []FakeExecCommand

// CreateFakeExecCommand mocks exec.Command with a unique id and appends it to MockArgs so it can be checked later
func CreateFakeExecCommand(uniqueId string) func(command string, args ...string) *exec.Cmd {
	if uniqueId == "" {
		uniqueId = "default"
	}
	if len(GetMockArgsByUniqueId(uniqueId)) > 0 {
		panic("uniqueId already exists")
	}
	return func(command string, args ...string) *exec.Cmd {
		MockArgs = append(MockArgs, FakeExecCommand{
			UniqueId: uniqueId,
			Command:  command,
			Args:     args,
		})
		return createFakeExecCommand(command, args...)
	}
}

// ResetMockArgs resets MockArgs
func ResetMockArgs() {
	MockArgs = make([]FakeExecCommand, 0)
}

// GetMockArgs returns MockArgs
func GetMockArgs() []FakeExecCommand {
	return MockArgs
}

// GetMockArgsByUniqueId returns MockArgs filtered by uniqueId
func GetMockArgsByUniqueId(uniqueId string) []FakeExecCommand {
	var filteredArgs []FakeExecCommand
	for _, arg := range MockArgs {
		if arg.UniqueId == uniqueId {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	return filteredArgs
}

// FakeExecCommand mocks exec.Command
func createFakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}
