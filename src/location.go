package main

import (
	"astrocalc/src/capitals"
	"astrocalc/src/location"
)

func getLocation(arr []string) (loc location.Location) {
	caps := capitals.Init()
	if len(arr) == 1 {
		loc = caps.GetLocation(arr[0])
	}
	if len(arr) > 1 {
		coords, err := parseCoords(arr)
		if err == nil {
			loc.Coords.Lat = coords.Lat
			loc.Coords.Lon = coords.Lon
			loc.Name = "custom"
		}
	}
	return loc
}
