package execwrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	// 	t.Run("returns list of strings", func(t *testing.T) {
	// 		ec := &MockExecCommand{
	// 			CommandFunc: func(name string, arg ...string) *exec.Cmd {
	// 				return &exec.Cmd{
	// 					Output func(): []byte(`402.0 /Users/joshmedeski/c/sesh
	// 338.0 /Users/joshmedeski/c/dotfiles
	// 74.0 /Users/joshmedeski/c/dotfiles/.config/nvim
	// 25.0 /Users/joshmedeski/c/dotfiles/.config/tmux`),
	// 				}
	// 			},
	// 		}
	// 		exec := &CmdExecutor{ExecCommander: ec}
	// 		list, err := exec.List("zoxide", "query", "--list", "--score")
	// 		expectedList := []string{
	// 			"402.0 /Users/joshmedeski/c/sesh",
	// 			"138.0 /Users/joshmedeski/c/dotfiles",
	// 			" 74.0 /Users/joshmedeski/c/dotfiles/.config/nvim",
	// 			"  5.0 /Users/joshmedeski/c/dotfiles/.config/tmux",
	// 		}
	// 		require.Equal(t, expectedList, list)
	// 		require.Nil(t, err)
	// 	})

	t.Run("returns list of strings", func(t *testing.T) {
		ec := &RealExecCommand{}
		exec := &CmdExecutor{ExecCommander: ec}
		list, err := exec.List("zoxide", "query", "--list", "--score")
		expectedList := []string{
			"402.0 /Users/joshmedeski/c/sesh",
			"338.0 /Users/joshmedeski/c/dotfiles",
			"74.0 /Users/joshmedeski/c/dotfiles/.config/nvim",
			"25.0 /Users/joshmedeski/c/dotfiles/.config/tmux",
		}
		require.Equal(t, nil, err)
		require.Equal(t, expectedList, list)
	})
}
