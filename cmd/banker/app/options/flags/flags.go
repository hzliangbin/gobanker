package flags

import (
	"github.com/urfave/cli/v2"

	"github.com/hzliangbin/gobanker/cmd/banker/app/options"
)

var (
	BankerOpts = options.NewBankerOptions()

	Flag = []cli.Flag{
		&cli.BoolFlag{
			Name:        "ues-config-file",
			Aliases:     []string{"c,C"},
			Value:       false,
		},
	}
)