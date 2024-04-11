package execwrap

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	t.Run("returns path when found", func(t *testing.T) {
		mockCmd := &MockExecCommand{
			LookPathFunc: func(file string) (string, error) {
				return "/opt/homebrew/bin/zoxide", nil
			},
		}
		exec := &CmdExecutor{ExecCommander: mockCmd}
		path, err := exec.Path("zoxide")
		require.Equal(t, "/opt/homebrew/bin/zoxide", path)
		require.Equal(t, nil, err)
	})

	t.Run("returns error when not found", func(t *testing.T) {
		mockCmd := &MockExecCommand{
			LookPathFunc: func(file string) (string, error) {
				return "", errors.New("exec: \"zoxide\": executable file not found in $PATH")
			},
		}
		exec := &CmdExecutor{ExecCommander: mockCmd}
		path, err := exec.Path("zoxide")
		require.Equal(t, "", path)
		require.Equal(t, "exec: \"zoxide\": executable file not found in $PATH", err.Error())
	})
}
