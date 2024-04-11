package icons

import (
	"fmt"

	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/session"
)

var Icons = map[session.SrcType]string{
	session.Zoxide: "",
	session.Tmux:   "",
	session.Config: "",
}

var icon_colors = map[session.SrcType]int{
	session.Zoxide: 36, // cyan
	session.Tmux:   34, // blue
	session.Config: 90, // gray
}

// TODO: add to config to allow for custom icons
func ansiString(code int, s string) string {
	return fmt.Sprintf("\033[%dm%s\033[39m", code, s)
}

func icon(src session.SrcType, config config.Config) string {
	if config.Icons.Tmux != "" && src == "tmux" {
		return config.Icons.Tmux
	}
	if config.Icons.Config != "" && src == "config" {
		return config.Icons.Config
	}
	if config.Icons.Zoxide != "" && src == "zoxide" {
		return config.Icons.Zoxide
	}
	return Icons[src]
}

func PrependIcon(s session.Session, c config.Config) string {
	icon := icon(s.Src, c)
	return fmt.Sprintf("%s %s", ansiString(icon_colors[s.Src], icon), s.Name)
}
