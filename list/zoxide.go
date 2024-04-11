package list

import (
	"fmt"

	"github.com/joshmedeski/sesh/session"
	"github.com/joshmedeski/sesh/zoxide"
)

func listZoxideSessions(existingSessions []session.Session) (sessions []session.Session, err error) {
	results, err := zoxide.List()
	if err != nil {
		return nil, fmt.Errorf("couldn't list zoxide results: %q", err)
	}
	var zoxideSessions []session.Session
	sessionMap := session.Map(existingSessions)
	for _, result := range results {
		if !session.IsInMap(sessionMap, result.Path) {
			zoxideSessions = append(zoxideSessions, session.Session{
				Src:   "zoxide",
				Name:  result.Name,
				Path:  result.Path,
				Score: result.Score,
			})
		}
	}
	return zoxideSessions, nil
}
