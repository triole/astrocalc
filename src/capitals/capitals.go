package capitals

import (
	"astrocalc/src/location"
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed capitals.json
var capitalsEmbed string

// Self holds the capitals class
type Self struct {
	Capitals tCapitals
}

type tCapitals []location.Location

// Init does what it says
func Init() (s Self) {
	s.Capitals = readJSON(capitalsEmbed)
	return s
}

func readJSON(content string) (capitals tCapitals) {
	err := json.Unmarshal([]byte(content), &capitals)
	if err != nil {
		panic(err)
	}
	return
}

func (self Self) GetLocation(s string) (loc location.Location) {
	for _, cap := range self.Capitals {
		if strings.ToLower(s) == strings.ToLower(cap.Capital) {
			loc = cap
			break
		}
	}
	return
}
