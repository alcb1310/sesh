package execwrap

import "os/exec"

type MockExecCommand struct {
	LookPathFunc func(file string) (string, error)
	CommandFunc  func(name string, arg ...string) *exec.Cmd
}

func (m *MockExecCommand) LookPath(file string) (string, error) {
	return m.LookPathFunc(file)
}

func (m *MockExecCommand) Command(name string, arg ...string) *exec.Cmd {
	return m.CommandFunc(name, arg...)
}

type MockExecutor struct {
	RunFunc  func(cmd string, args ...string) (string, error)
	ListFunc func(cmd string, args ...string) ([]string, error)
}

func (m *MockExecutor) Run(cmd string, args ...string) (string, error) {
	if m.RunFunc != nil {
		return m.RunFunc(cmd, args...)
	}
	return "", nil
}

func (m *MockExecutor) List(cmd string, args ...string) ([]string, error) {
	if m.ListFunc != nil {
		return m.ListFunc(cmd, args...)
	}
	return nil, nil
}
