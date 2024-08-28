package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "astrocalc"
	appDescription = "simple astronomical calculation"
	appMainversion = "0.2"
)

var CLI struct {
	Location    []string `help:"location to use" arg:"" optional:"" default:"berlin"`
	Format      string   `help:"output format, can be: json, toml, yaml" short:"f" enum:"json,toml,yaml" default:"json"`
	VersionFlag bool     `help:"display version" short:"V"`
}

func parseArgs() {
	curdir, _ := os.Getwd()
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"curdir":  curdir,
			"logfile": path.Join(os.TempDir(), alnum(appName)+".log"),
			"output":  path.Join(curdir, "tyson.json"),
			"proc":    strconv.Itoa(runtime.NumCPU()),
			"filter":  "\\.(md|toml|yaml|json)$",
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}

func alnum(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z0-9_-]")
	return re.ReplaceAllString(s, "-")
}
