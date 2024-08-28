package main

import (
	"bytes"
	"encoding/json"
	"log"

	toml "github.com/pelletier/go-toml"
)

func stringToToml(itf interface{}) []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(itf); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func stringToJSON(i interface{}) []byte {
	s, _ := json.MarshalIndent(i, "", "  ")
	return s
}
