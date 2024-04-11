package list

import (
	"testing"

	"github.com/joshmedeski/sesh/session"
	realTmux "github.com/joshmedeski/sesh/tmux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockTmux struct {
	mock.Mock
}

func (m *MockTmux) List() ([]realTmux.TmuxSession, error) {
	args := m.Called()
	return args.Get(0).([]realTmux.TmuxSession), args.Error(1)
}

func TestTmuxList(t *testing.T) {
	t.Run("List tmux sessions", func(t *testing.T) {
		tmux := new(MockTmux)
		tmux.On("List").Return([]realTmux.TmuxSession{}, nil)

		actual, _ := listTmuxSessions(Options{})
		expected := []session.Session{
			{
				Src:      "Tmux",
				Name:     "dotfiles",
				Path:     "/Users/joshmedeski/c/dotfiles",
				Attached: 0,
				Windows:  1,
			},
		}
		require.Equal(t, expected, actual)
	})
}
