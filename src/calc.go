package main

import (
	"time"

	"github.com/sixdouglas/suncalc"
)

type tResults struct {
	Time     time.Time              `json:"time" toml:"time" yaml:"time"`
	Location map[string]interface{} `json:"location" toml:"location" yaml:"location"`
	Sun      tSun                   `json:"sun" toml:"sun" yaml:"sun"`
	Moon     tMoon                  `json:"moon" toml:"moon" yaml:"moon"`
}

type tSun struct {
	Light    tSunLight `json:"light" toml:"light" yaml:"light"`
	Position tPosition `json:"position" toml:"position" yaml:"position"`
}

type tMoon struct {
	Light        tMoonLight        `json:"light" toml:"light" yaml:"light"`
	Position     tPosition         `json:"position" toml:"position" yaml:"position"`
	Illumination tMoonIllumination `json:"illumination" toml:"illumination" yaml:"illumination"`
}

type tSunLight map[string]time.Time
type tMoonLight map[string]interface{}
type tMoonIllumination map[string]interface{}
type tPosition map[string]float64

func newDataset() (ds tResults) {
	ds.Sun.Light = make(tSunLight)
	ds.Sun.Position = make(tPosition)
	ds.Moon.Light = make(tMoonLight)
	ds.Moon.Position = make(tPosition)
	ds.Moon.Illumination = make(tMoonIllumination)
	return
}

func calc(now time.Time, lat float64, lon float64) (res tResults) {
	res = newDataset()

	res.Time = now
	res.Location = make(map[string]interface{})
	res.Location["lat"] = lat
	res.Location["lon"] = lon

	// sunlight times
	arr := suncalc.GetTimes(now, lat, lon)
	for key, val := range arr {
		res.Sun.Light[string(key)] = val.Value
	}

	a := suncalc.GetPosition(now, lat, lon)
	res.Sun.Position = make(map[string]float64)
	res.Sun.Position["altitude"] = a.Altitude
	res.Sun.Position["azimuth"] = a.Azimuth

	d := suncalc.GetMoonTimes(now, lat, lon, false)
	res.Moon.Light["rise"] = d.Rise
	res.Moon.Light["set"] = d.Set
	res.Moon.Light["alwaysUp"] = d.AlwaysUp
	res.Moon.Light["alwaysDown"] = d.AlwaysDown

	b := suncalc.GetMoonPosition(now, lat, lon)
	res.Moon.Position["altitude"] = b.Altitude
	res.Moon.Position["azimuth"] = b.Azimuth
	res.Moon.Position["distance"] = b.Distance
	res.Moon.Position["parallacticAngle"] = b.ParallacticAngle

	c := suncalc.GetMoonIllumination(now)
	res.Moon.Illumination["fraction"] = c.Fraction
	res.Moon.Illumination["phase"] = c.Phase
	res.Moon.Illumination["angle"] = c.Angle

	return
}
