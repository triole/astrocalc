package main

import (
	"time"

	"github.com/sixdouglas/suncalc"
)

func calc(now time.Time, lat float64, lon float64) (m map[string]suncalc.DayTime) {
	arr := suncalc.GetTimes(now, lat, lon)
	m = make(map[string]suncalc.DayTime)
	for key, val := range arr {
		m[string(key)] = val
	}
	return
}
