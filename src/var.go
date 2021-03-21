package main

import (
	"fmt"
	"olibs/environment"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "suncalc"
	appMainversion = "0.1"
	appDescription = "suncalc command line tool"
	env            = environment.Init(appName, appMainversion, appDescription, BUILDTAGS)

	app      = kingpin.New(appName, appDescription)
	argsCity = app.Arg("city", "city").Default("Potsdam").String()
)

func argparse() {
	app.Version(env.AppInfoString)
	app.VersionFlag.Short('V')
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func check(err error) {
	if err != nil {
		fmt.Printf("An error occured: %q\n", err)
	}
}
