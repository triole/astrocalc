package main

import (
	"bytes"
	"log"

	"github.com/BurntSushi/toml"
)

type tGeoConf map[string][]float64

func readTomlFile(content string) tGeoConf {
	geo := make(tGeoConf)
	if _, err := toml.Decode(string(content), &geo); err != nil {
		panic(err)
	}
	return geo
}

func toTomlString(itf interface{}) string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(itf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
