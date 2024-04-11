package session

type (
	SrcType string

	Session struct {
		Src      SrcType
		Name     string  // The display name
		Path     string  // The absolute directory path
		Score    float64 // The score of the session (from Zoxide)
		Attached int     // Whether the session is currently attached
		Windows  int     // The number of windows in the session
	}

	Srcs struct {
		Config bool
		Tmux   bool
		Zoxide bool
	}
)

const (
	Tmux   SrcType = "tmux"
	Zoxide SrcType = "zoxide"
	Config SrcType = "config"
)
