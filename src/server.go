package main

import (
	"astrocalc/src/location"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func runServer() {
	fmt.Printf("%s listening at %d\n", ts(), CLI.Port)
	http.HandleFunc("/", returnHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(CLI.Port), nil)
	if err != nil {
		panic(err)
	}
}

func returnHandler(w http.ResponseWriter, r *http.Request) {
	paramLoc := r.URL.Query().Get("loc")
	paramLat := r.URL.Query().Get("lat")
	paramLon := r.URL.Query().Get("lon")

	if paramLoc == "" && paramLat == "" && paramLon == "" {
		fmt.Fprintf(
			w,
			"location request parameter missing, e.g. %q or %q",
			"?loc=london", "?lat=51.50216&lon=-0.109",
		)
	} else {

		var loc location.Location

		if paramLoc != "" {
			loc = getLocation([]string{paramLoc})
		} else if paramLat != "" && paramLon != "" {
			loc = getLocation([]string{paramLat, paramLon})
		}
		now := time.Now()
		data := calc(now, loc.Coords.Lat, loc.Coords.Lon, loc.Name)
		fmt.Printf("%s respond for location %+v\n", ts(), data.Location)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}
