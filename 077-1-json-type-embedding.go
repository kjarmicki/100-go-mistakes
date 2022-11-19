package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
 * When embedding a field, beware of fields that implement json.Marshaler interface.
 * Because embedding lifts everything up to the parent level, the entire struct will be marshaled with the embedded field implementation.
 */

type JsonEvent struct {
	ID int
	time.Time
}

// here, it won't be {ID: 1234, Time: "2022-11-18T00:00:00Z"} but instead just "2022-11-18T00:00:00Z"
// because time.Time is embedded and it implements json.Marshaler so it will hijack the marshaling
func JsonMarshalEmbeddedType() {
	t, _ := time.Parse("2006-01-02", "2022-11-18")
	event := JsonEvent{
		ID:   1234,
		Time: t,
	}

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(b))
}

// possible solutions:
// - put time under Time field, making it no longer embedded
// - implement json.Marshaler on the JsonEvent level
