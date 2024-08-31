package main

import (
	"astrocalc/src/capitals"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func runServer() {
	fmt.Printf("listening at %d", CLI.Port)
	http.HandleFunc("/", returnHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(CLI.Port), nil)
	if err != nil {
		panic(err)
	}
}

func returnHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("loc")
	if location != "" {
		now := time.Now()
		caps := capitals.Init()
		loc := caps.GetLocation(location)
		data := calc(now, loc.Coords.Lat, loc.Coords.Lon)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	} else {
		fmt.Fprintf(
			w, "location request parameter missing, add something like %s",
			"?loc=london",
		)
	}

}
