package main

import (
	"bytes"
	"encoding/json"
	"log"

	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
)

func stringToJSON(i interface{}) []byte {
	s, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func stringToTOML(itf interface{}) []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(itf); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func stringToYAML(i interface{}) []byte {
	s, err := yaml.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
