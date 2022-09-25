package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `A command-line tool for converting a stream of numbers into total, averages, etc

Usage:
  statz
  statz -h | --help
  statz --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/statz v0.0.1")
	fmt.Println(opts)
}
