package connect

import (
	"fmt"
	"strings"

	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/icons"
	"github.com/joshmedeski/sesh/session"
	"github.com/joshmedeski/sesh/tmux"
	"github.com/joshmedeski/sesh/zoxide"
)

var prefixes = []string{
	icons.Icons[session.Tmux],
	icons.Icons[session.Zoxide],
	icons.Icons[session.Config],
}

func Connect(
	choice string,
	alwaysSwitch bool,
	command string,
	config *config.Config,
) error {
	for _, prefix := range prefixes {
		if strings.HasPrefix(choice, prefix) {
			choice = choice[4:]
			break
		}
	}

	session, err := session.Determine(choice, config)
	if err != nil {
		return fmt.Errorf("couldn't determine session for '%q': %w", choice, err)
	}

	if err = zoxide.Add(session.Path); err != nil {
		return fmt.Errorf("couldn't add to zoxide for '%q': %w", choice, err)
	}

	return tmux.Connect(tmux.TmuxSession{
		Name: session.Name,
		Path: session.Path,
	}, alwaysSwitch, command, session.Path, config)
}
