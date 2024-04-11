package list

import (
	"fmt"

	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/session"
)

type Options struct {
	HideAttached bool
	Icons        bool
	ShowConfig   bool
	ShowTmux     bool
	ShowZoxide   bool
}

func List(o Options, config *config.Config) (s []session.Session, err error) {
	var sessions []session.Session
	anySrcs := o.ShowTmux || o.ShowConfig || o.ShowZoxide

	if !anySrcs || o.ShowTmux {
		tmuxSessions, err := listTmuxSessions(o)
		if err != nil {
			return nil, fmt.Errorf("failed to list tmux sessions: %v", err)
		}
		sessions = append(sessions, tmuxSessions...)
	}

	if !anySrcs || o.ShowConfig {
		configSessions, err := listConfigSessions(config, sessions)
		if err != nil {
			return nil, fmt.Errorf("failed to list config sessions: %v", err)
		}
		sessions = append(sessions, configSessions...)
	}

	if !anySrcs || o.ShowZoxide {
		zoxideSessions, err := listZoxideSessions(sessions)
		if err != nil {
			return nil, fmt.Errorf("failed to list zoxide sessions: %v", err)
		}
		sessions = append(sessions, zoxideSessions...)
	}

	return sessions, nil
}
