package main

import (
	"strconv"
	"strings"
)

// represents a single data point i.e. a single line in the incoming data stream
type datapoint struct {
	rawString  string
	cleanFloat float64
	decPlaces  int
}

// create a new 'datapoint' struct based on the given line from the input stream
func newDatapoint(s string) (*datapoint, error) {
	decPlaces := 0
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	words := strings.Split(s, ".")
	if len(words) > 1 {
		decPlaces = len(words[len(words)-1])
	}
	return &datapoint{
		rawString:  s,
		cleanFloat: f,
		decPlaces:  decPlaces,
	}, nil
}
