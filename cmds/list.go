package cmds

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"

	"github.com/joshmedeski/sesh/config"
	"github.com/joshmedeski/sesh/icons"
	"github.com/joshmedeski/sesh/json"
	"github.com/joshmedeski/sesh/list"
)

func List() *cli.Command {
	return &cli.Command{
		Name:                   "list",
		Aliases:                []string{"l"},
		Usage:                  "List sessions",
		ArgsUsage:              "[no arguments allowed]",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "tmux",
				Aliases: []string{"t"},
				Usage:   "show tmux sessions",
			},
			&cli.BoolFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "show configured sessions",
			},
			&cli.BoolFlag{
				Name:    "zoxide",
				Aliases: []string{"z"},
				Usage:   "show zoxide results",
			},

			&cli.BoolFlag{
				Name:    "icons",
				Aliases: []string{"i"},
				Usage:   "append icons (using Nerd Font)",
			},
			&cli.BoolFlag{
				Name:    "hide-attached",
				Aliases: []string{"H"},
				Usage:   "don't show the currently attached session(s)",
			},

			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "output as json",
			},
		},
		Action: func(cCtx *cli.Context) error {
			options := list.Options{
				HideAttached: cCtx.Bool("hide-attached"),
				Icons:        cCtx.Bool("icons"),
				ShowConfig:   cCtx.Bool("config"),
				ShowTmux:     cCtx.Bool("tmux"),
				ShowZoxide:   cCtx.Bool("zoxide"),
			}
			// TODO: allow hiding attached globally with config
			config := config.ParseConfigFile(&config.DefaultConfigDirectoryFetcher{})

			sessions, err := list.List(options, &config)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return nil
			}

			useIcons := cCtx.Bool("icons")
			if useIcons {
				for i, s := range sessions {
					sessions[i].Name = icons.PrependIcon(s, config)
				}
			}

			useJson := cCtx.Bool("json")
			if useJson {
				fmt.Println(json.List(sessions))
				return nil
			}

			for _, session := range sessions {
				fmt.Println(session.Name)
			}
			return nil
		},
	}
}
