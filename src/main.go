package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed geo.toml
var geoConf string

func main() {
	argparse()

	geo := readTomlFile(geoConf)
	now := time.Now()
	coords := geo[*argsCity]

	t := calc(now, coords[0], coords[1])
	t.Location["name"] = *argsCity

	fmt.Printf("%s\n", toTomlString(t))
}
