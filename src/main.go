package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type tCoords struct {
	Lat float64
	Lon float64
}

func main() {
	parseArgs()

	if CLI.Server {
		runServer()
	} else {
		// try to find capital, if fails assume to be coords
		loc := getLocation(CLI.Location)
		if loc.Name == "" {
			displayErr()
		} else {
			now := time.Now()
			t := calc(now, loc)

			out := stringToJSON(t)

			if strings.EqualFold(CLI.Format, "toml") {
				out = stringToTOML(t)
			}
			if strings.EqualFold(CLI.Format, "yaml") {
				out = stringToYAML(t)
			}
			fmt.Printf("%s\n", out)
		}
	}
}

func parseCoords(arr []string) (c tCoords, err error) {
	var fl float64
	for idx, el := range arr {
		fl, err = strconv.ParseFloat(el, 64)
		if err == nil {
			if idx == 0 {
				c.Lat = fl
			}
			if idx == 1 {
				c.Lon = fl
			}
		} else {
			return tCoords{}, err
		}
	}
	return
}

func displayErr() {
	fmt.Println("can not parse args, location required")
	fmt.Println("use either capital's name or lat lon coordinates")
	fmt.Println("example: astrocalc berlin, astrocalc 55.2 22.1")
	os.Exit(1)
}

func ts() string {
	return time.Now().Format("[2006-01-02 15:04:05]")
}
