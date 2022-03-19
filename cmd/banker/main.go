package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	banker "github.com/hzliangbin/gobanker/cmd/banker/app"
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
	app.Before = banker.Before
	app.Action = banker.Start

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
