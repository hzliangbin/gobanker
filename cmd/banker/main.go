package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/hzliangbin/gobanker/cmd/banker/app/options/flags"
)

var version = "1.0.0"

func main() {
	rand.Seed(time.Now().UnixNano())
	app := cli.NewApp()
	app.Name = "Banker"
	app.Usage = "Banker is indicates A share temper"
	app.Version = version
	app.Compiled = time.Now()
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " Banker"

	app.Flags = flags.Flag
	app.Before = banker

}
