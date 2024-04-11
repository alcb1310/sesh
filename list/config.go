package list

import (
	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/session"
)

func listConfigSessions(c *config.Config, existingSessions []session.Session) (sessions []session.Session, err error) {
	var configSessions []session.Session
	sessionMap := session.Map(existingSessions)
	for _, sessionConfig := range c.SessionConfigs {
		if sessionConfig.Name != "" && !session.IsInMap(sessionMap, sessionConfig.Name) {
			configSessions = append(configSessions, session.Session{
				Src:  "config",
				Name: sessionConfig.Name,
				Path: sessionConfig.Path,
			})
		}
	}
	return configSessions, nil
}
