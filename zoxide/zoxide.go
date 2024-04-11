package zoxide

import (
	"github.com/joshmedeski/sesh/execwrap"
	"github.com/joshmedeski/sesh/models"
)

type Zoxide interface {
	Add(path string) error
	Remove(path string) error
	Query(path string) (models.ZoxideResult, error)
	List() ([]models.ZoxideResult, error)
}

type ZoxideExecutor struct {
	Exec execwrap.CmdExecutor
}

func NewZoxide(e execwrap.Executor) *Zoxide {
	return &Zoxide{Exec: execwrap.RealExecCommand}
}

func Add(path string) error {
	cmd, err := execwrap.Command("zoxide", "add", path)
	return nil
}
