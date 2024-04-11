package session

func Map(sessions []Session) map[string]bool {
	sessionMap := make(map[string]bool, len(sessions))
	for _, session := range sessions {
		sessionMap[session.Name] = true
	}
	return sessionMap
}

func IsInMap(sessionMap map[string]bool, name string) bool {
	_, exists := sessionMap[name]
	return exists
}
