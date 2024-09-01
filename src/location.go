package main

import (
	"astrocalc/src/capitals"
	"astrocalc/src/location"
)

func getLocation(arr []string) (loc location.Location) {
	caps := capitals.Init()
	if len(arr) == 1 {
		loc = caps.GetLocation(CLI.Location[0])
	}
	if len(arr) > 1 {
		coords, err := parseCoords(CLI.Location)
		if err == nil {
			loc.Capital = "custom"
			loc.Coords.Lat = coords.Lat
			loc.Coords.Lon = coords.Lon
		}
	}
	return loc
}
