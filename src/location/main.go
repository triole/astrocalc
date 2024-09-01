package location

type Location struct {
	Name    string
	Country string
	Coords  Coords
}

type Coords struct {
	Lat float64
	Lon float64
}
