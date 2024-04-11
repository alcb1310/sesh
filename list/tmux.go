package list

import (
	"fmt"

	"github.com/joshmedeski/sesh/session"
	"github.com/joshmedeski/sesh/tmux"
)

func listTmuxSessions(o Options) (sessions []session.Session, err error) {
	tmuxList, err := tmux.List(tmux.Options{
		HideAttached: o.HideAttached,
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't list tmux sessions: %q", err)
	}
	tmuxSessions := make([]session.Session, len(tmuxList))
	for i, s := range tmuxList {
		tmuxSessions[i] = session.Session{
			Src:      session.Tmux,
			Name:     s.Name,
			Path:     s.Path,
			Attached: s.Attached,
			Windows:  s.Windows,
		}
	}
	return tmuxSessions, nil
}
