package main

import (
	"encoding/json"
	"fmt"
)

/*
 * When unmarshaling a map of any that contains a number,
 * this number will be typed as float64.
 */

func JsonMarshalMap() {
	var m map[string]any
	json.Unmarshal([]byte(`
		{
			"id": 32,
			"name": "foo"
		}
	`), &m)
	fmt.Println(m)              // map[id:32 name:foo]
	fmt.Printf("%T\n", m["id"]) // float64
}
