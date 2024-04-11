package session

import (
	"fmt"
	"path/filepath"

	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/name"
	"github.com/joshmedeski/sesh/tmux"
	"github.com/joshmedeski/sesh/zoxide"
)

func tmuxSession(choice string) *Session {
	sessions, _ := tmux.List(tmux.Options{
		HideAttached: false,
	})
	for _, s := range sessions {
		if s.Name == choice {
			return &Session{
				Src:  "tmux",
				Name: s.Name,
				Path: s.Path,
			}
		}
	}
	return nil
}

func configByNameSession(choice string, config *config.Config) *Session {
	for _, s := range config.SessionConfigs {
		if s.Name == choice {
			return &Session{
				Src:  "config",
				Name: s.Name,
				Path: s.Path,
			}
		}
	}
	return nil
}

func configByPathSession(choice string, config *config.Config) *Session {
	for _, s := range config.SessionConfigs {
		if s.Path == choice {
			return &Session{
				Src:  "config",
				Name: name.DetermineName(choice, s.Path),
				Path: s.Path,
			}
		}
	}
	return nil
}

func zoxideSession(choice string) *Session {
	result, _ := zoxide.Query(choice)
	if result != nil {
		return &Session{
			Src:  "zoxide",
			Name: name.DetermineName(result.Name, result.Path),
			Path: result.Path,
		}
	}
	return nil
}

func pathSession(choice string) *Session {
	path, err := filepath.Abs(choice)
	if err == nil {
		return &Session{
			Src:  "path",
			Name: name.DetermineName(choice, path),
			Path: path,
		}
	}
	return nil
}

func Determine(choice string, config *config.Config) (s *Session, err error) {
	tmuxSession := tmuxSession(choice)
	if tmuxSession != nil {
		return tmuxSession, nil
	}

	configByNameSession := configByNameSession(choice, config)
	if configByNameSession != nil {
		return configByNameSession, nil
	}

	configByPathSession := configByPathSession(choice, config)
	if configByPathSession != nil {
		return configByPathSession, nil
	}

	pathSession := pathSession(choice)
	if pathSession != nil {
		return pathSession, nil
	}

	zoxideSession := zoxideSession(choice)
	if zoxideSession != nil {
		return zoxideSession, nil
	}

	return nil, fmt.Errorf("session could not be determined from choice: %s", choice)
}
