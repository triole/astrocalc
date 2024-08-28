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
type tCapital struct {
	Capitals tCapitals
}

type tCapitals []location.Location

// Init does what it says
func Init() (cap tCapital) {
	cap.Capitals = readJSON(capitalsEmbed)
	return cap
}

func readJSON(content string) (capitals tCapitals) {
	err := json.Unmarshal([]byte(content), &capitals)
	if err != nil {
		panic(err)
	}
	return
}

func (cap tCapital) GetLocation(s string) (loc location.Location) {
	for _, cap := range cap.Capitals {
		if strings.EqualFold(s, cap.Capital) {
			loc = cap
			break
		}
	}
	return
}
