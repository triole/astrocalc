package main

import (
	"astrocalc/src/capitals"
	"astrocalc/src/location"
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

	var loc location.Location

	if CLI.Server {
		runServer()
	} else {
		// try to find capital, if fails assume to be coords
		caps := capitals.Init()
		capital := caps.GetLocation(CLI.Location[0])
		if capital.Capital != "" {
			loc = capital
		} else {
			coords, err := parseCoords(CLI.Location)
			if err == nil {
				loc.Capital = "custom"
				loc.Coords.Lat = coords.Lat
				loc.Coords.Lon = coords.Lon
			}
		}

		if loc.Capital == "" {
			displayErr()
		} else {
			now := time.Now()
			t := calc(now, loc.Coords.Lat, loc.Coords.Lon)
			t.Location["name"] = loc.Capital

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
	fmt.Println("Failed to parse args.")
	fmt.Println("Location required, either capital's name or lat lon coordinates")
	fmt.Println("Examples: astrocalc berlin, astrocalc 55.2 22.1")
	os.Exit(1)
}
