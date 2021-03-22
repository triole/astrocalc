package main

import (
	"bytes"
	"log"

	"github.com/BurntSushi/toml"
)

func toTomlString(itf interface{}) string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(itf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
