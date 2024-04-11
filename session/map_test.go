package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var sessionsMap = Map([]Session{
	{
		Src:  Config,
		Name: "dotfiles",
		Path: "/Users/joshmedeski/c/dotfiles",
	},
})

func TestMap(t *testing.T) {
	t.Run("Map by name", func(t *testing.T) {
		expected := map[string]bool{"dotfiles": true}
		require.Equal(t, expected, sessionsMap)
	})

	t.Run("IsInMap true", func(t *testing.T) {
		require.Equal(t, true, IsInMap(sessionsMap, "dotfiles"))
	})

	t.Run("IsInMap false", func(t *testing.T) {
		require.Equal(t, false, IsInMap(sessionsMap, "home"))
	})
}
