package execwrap

import (
	"os/exec"
	"strings"
)

type Executor interface {
	// Path looks up the path of the provided executable and returns the
	// full path if the executable is found, otherwise it returns an error.
	Path(executable string) (string, error)
	// Run executes the command and returns the combined output of stdout and stderr.
	Run(cmd string, args ...string) (string, error)
	// List executes the command and returns a slice of strings.
	List(cmd string, args ...string) ([]string, error)
}

type CmdExecutor struct {
	// Commander is an interface that defines methods for executing commands.
	Commander interface {
		LookPath(file string) (string, error)
		Command(name string, arg ...string) *exec.Cmd
	}
}

// OSCommander is a struct that implements the Commander interface using the os/exec package.
type OSCommander struct{}

func (c *OSCommander) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (c *OSCommander) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

func (c *CmdExecutor) Path(executable string) (string, error) {
	path, err := c.Commander.LookPath(executable)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *CmdExecutor) List(name string, args ...string) ([]string, error) {
	cmd := c.Commander.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outputStr := string(output)
	list := strings.Split(outputStr, "\n")
	return list, nil
}
