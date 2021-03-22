package location

type Location struct {
	Capital string
	Country string
	Coords  Coords
}

type Coords struct {
	Lat float64
	Lon float64
}
