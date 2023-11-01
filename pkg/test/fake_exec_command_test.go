/**
 * Author: Andrew Shoell
 * File: fake_exec_command_test.go
 */

package test

import (
	"reflect"
	"testing"
)

// Cleanup MockArgs after tests
func Cleanup() {
	ResetMockArgs()
}

// TestCreateFakeExecCommand tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs
func TestFakeExecCommandStruct(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	if fakeExecCommand.UniqueId != uniqueId {
		Cleanup()
		t.Errorf("FakeExecCommand.UniqueId = %s, expected %s", fakeExecCommand.UniqueId, uniqueId)
	}

	if fakeExecCommand.Command != command {
		Cleanup()
		t.Errorf("FakeExecCommand.Command = %s, expected %s", fakeExecCommand.Command, command)
	}

	if !reflect.DeepEqual(fakeExecCommand.Args, args) {
		Cleanup()
		t.Errorf("FakeExecCommand.Args = %s, expected %s", fakeExecCommand.Args, args)
	}
	Cleanup()
}

// TestCreateFakeExecCommandNoUniqueId tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestFakeExecCommandStructNoUniqueId(t *testing.T) {
	Cleanup()
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: "default",
		Command:  command,
		Args:     args,
	}

	if fakeExecCommand.UniqueId != "default" {
		Cleanup()
		t.Errorf("FakeExecCommand.UniqueId = %s, expected \"default\"", fakeExecCommand.UniqueId)
	}

	if fakeExecCommand.Command != command {
		Cleanup()
		t.Errorf("FakeExecCommand.Command = %s, expected %s", fakeExecCommand.Command, command)
	}

	if !reflect.DeepEqual(fakeExecCommand.Args, args) {
		Cleanup()
		t.Errorf("FakeExecCommand.Args = %s, expected %s", fakeExecCommand.Args, args)
	}
	Cleanup()
}

// TestCreateFakeExecCommandDuplicateUniqueId tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestFakeExecCommandStructDuplicateUniqueId(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	CreateFakeExecCommand(uniqueId)(command, args...)

	defer func() {
		if r := recover(); r == nil {
			Cleanup()
			t.Errorf("The code did not panic")
		}
	}()

	CreateFakeExecCommand(uniqueId)(command, args...)
	Cleanup()
}

// TestCreateFakeExecCommandNoCommand tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestFakeExecCommandStructNoCommand(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  "",
		Args:     args,
	}

	if fakeExecCommand.UniqueId != uniqueId {
		Cleanup()
		t.Errorf("FakeExecCommand.UniqueId = %s, expected %s", fakeExecCommand.UniqueId, uniqueId)
	}

	if fakeExecCommand.Command != "" {
		Cleanup()
		t.Errorf("FakeExecCommand.Command = %s, expected \"\"", fakeExecCommand.Command)
	}

	if !reflect.DeepEqual(fakeExecCommand.Args, args) {
		Cleanup()
		t.Errorf("FakeExecCommand.Args = %s, expected %s", fakeExecCommand.Args, args)
	}
	Cleanup()
}

// TestCreateFakeExecCommandNoArgs tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestFakeExecCommandStructNoArgs(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     nil,
	}

	if fakeExecCommand.UniqueId != uniqueId {
		Cleanup()
		t.Errorf("FakeExecCommand.UniqueId = %s, expected %s", fakeExecCommand.UniqueId, uniqueId)
	}

	if fakeExecCommand.Command != command {
		Cleanup()
		t.Errorf("FakeExecCommand.Command = %s, expected %s", fakeExecCommand.Command, command)
	}

	if fakeExecCommand.Args != nil {
		Cleanup()
		t.Errorf("FakeExecCommand.Args = %s, expected nil", fakeExecCommand.Args)
	}
	Cleanup()
}

// TestMockArgs tests that MockArgs is a list of FakeExecCommand
func TestMockArgs(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if !reflect.DeepEqual(MockArgs[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("MockArgs[0] = %s, expected %s", MockArgs[0], fakeExecCommand)
	}
	Cleanup()
}

// TestCreateFakeExecCommand tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs
func TestCreateFakeExecCommand(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	CreateFakeExecCommand(uniqueId)(command, args...)

	if !reflect.DeepEqual(MockArgs[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("MockArgs[0] = %s, expected %s", MockArgs[0], fakeExecCommand)
	}
	Cleanup()
}

// TestCreateFakeExecCommandNilMockArgs tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs
func TestCreateFakeExecCommandNilMockArgs(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	CreateFakeExecCommand(uniqueId)(command, args...)

	if !reflect.DeepEqual(MockArgs[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("MockArgs[0] = %s, expected %s", MockArgs[0], fakeExecCommand)
	}
	Cleanup()
}

// TestCreateFakeExecCommandDefaultUniqueId tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestCreateFakeExecCommandDefaultUniqueId(t *testing.T) {
	Cleanup()
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: "default",
		Command:  command,
		Args:     args,
	}

	CreateFakeExecCommand("")(command, args...)

	if !reflect.DeepEqual(MockArgs[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("MockArgs[0] = %s, expected %s", MockArgs[0], fakeExecCommand)
	}
	Cleanup()
}

// TestCreateFakeExecCommandDuplicateUniqueId tests that CreateFakeExecCommand returns a function that appends a FakeExecCommand to MockArgs with a default uniqueId
func TestCreateFakeExecCommandDuplicateUniqueId(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	CreateFakeExecCommand(uniqueId)(command, args...)

	defer func() {
		if r := recover(); r == nil {
			Cleanup()
			t.Errorf("The code did not panic")
		}
	}()

	CreateFakeExecCommand(uniqueId)(command, args...)
	Cleanup()
}

// TestResetMockArgs tests that ResetMockArgs resets MockArgs
func TestResetMockArgs(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if !reflect.DeepEqual(MockArgs[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("MockArgs[0] = %s, expected %s", MockArgs[0], fakeExecCommand)
	}

	ResetMockArgs()

	if len(MockArgs) != 0 {
		Cleanup()
		t.Errorf("len(MockArgs) = %d, expected 0", len(MockArgs))
	}
	Cleanup()
}

// TestGetMockArgs tests that GetMockArgs returns MockArgs
func TestGetMockArgs(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if !reflect.DeepEqual(GetMockArgs()[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("GetMockArgs()[0] = %s, expected %s", GetMockArgs()[0], fakeExecCommand)
	}
	Cleanup()
}

// TestGetMockArgsByUniqueId tests that GetMockArgsByUniqueId returns MockArgs filtered by uniqueId
func TestGetMockArgsByUniqueId(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if !reflect.DeepEqual(GetMockArgsByUniqueId(uniqueId)[0], fakeExecCommand) {
		Cleanup()
		t.Errorf("GetMockArgsByUniqueId(uniqueId)[0] = %s, expected %s", GetMockArgsByUniqueId(uniqueId)[0], fakeExecCommand)
	}
	Cleanup()
}

// TestGetMockArgsByUniqueIdEmpty tests that GetMockArgsByUniqueId returns an empty list when no uniqueId is found
func TestGetMockArgsByUniqueIdBad(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if len(GetMockArgsByUniqueId("notUniqueId")) != 0 {
		Cleanup()
		t.Errorf("len(GetMockArgsByUniqueId(\"notUniqueId\")) = %d, expected 0", len(GetMockArgsByUniqueId("notUniqueId")))
	}
	Cleanup()
}

// TestGetMockArgsByUniqueIdEmpty tests that GetMockArgsByUniqueId returns an empty list when uniqueId is empty
func TestGetMockArgsByUniqueIdEmpty(t *testing.T) {
	Cleanup()
	uniqueId := "uniqueId"
	command := "command"
	args := []string{"arg1", "arg2"}

	fakeExecCommand := FakeExecCommand{
		UniqueId: uniqueId,
		Command:  command,
		Args:     args,
	}

	MockArgs = append(MockArgs, fakeExecCommand)

	if len(GetMockArgsByUniqueId("")) != 0 {
		Cleanup()
		t.Errorf("len(GetMockArgsByUniqueId(\"\")) = %d, expected 0", len(GetMockArgsByUniqueId("")))
	}
	Cleanup()
}
