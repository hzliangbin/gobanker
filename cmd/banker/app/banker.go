package app

import (
	"github.com/urfave/cli/v2"

	"github.com/hzliangbin/gobanker/cmd/banker/app/options"
	"github.com/hzliangbin/gobanker/cmd/banker/app/options/flags"
	"github.com/hzliangbin/gobanker/pkg/apiserver"
	"github.com/hzliangbin/gobanker/pkg/blog"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

var (
	Before = func(c *cli.Context) error {
		if !c.Bool("uesConfigFile") {
			return nil
		}

		var err error
		flags.BankerOpts, err = options.LoadConfigFromDisk()
		if err != nil {
			return err
		}

		return nil
	}

	Start = func(c *cli.Context) error {
		if errs := flags.BankerOpts.Validate(); len(errs) > 0 {
			// TODO we should not import k8s dependency
			return utilerrors.NewAggregate(errs)
		}

		run(flags.BankerOpts)

		return nil
	}
)

func run(opts * options.BankerOptions) {
	blog.InitLoggerWithOpts(opts.LoggerOpts)

	apiserver.NewAPIServerWithOpts(opts.APIServerOpts).Run()
}