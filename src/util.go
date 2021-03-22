package main

import (
	"encoding/json"
	"fmt"
)

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func check(err error) {
	if err != nil {
		fmt.Printf("An error occured: %q\n", err)
	}
}
