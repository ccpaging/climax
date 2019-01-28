package commands

import (
	"fmt"
	"github.com/ccpaging/climax"
	"strings"
)

var JoinCmd = &climax.Command{
	Name:  "join",
	Brief: "merges the strings given",
	Usage: `[-s=] "a few" distinct strings`,
	Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

	Flags: []*climax.Flag{
		&climax.Flag{
			Name:     "separator",
			Short:    "s",
			Usage:    `--separator="."`,
			Help:     `Put some separating string between all the strings given.`,
			Variable: true,
		},
	},

	Examples: []*climax.Example{
		&climax.Example{
			Usecase:     `-s . "google" "com"`,
			Description: `Results in "google.com"`,
		},
	},

	Handle: func(ctx *climax.Context) int {
		var separator string
		if sep, ok := ctx.Get("separator"); ok {
			separator = sep
		}

		fmt.Println(strings.Join(ctx.Args, separator))

		return 0
	},
}

func init() {
	App.AddCommand(JoinCmd)
}
