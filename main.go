package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
)

func outputResults(m map[string]string) {
	maxDataWidth := 0
	for _, data := range m {
		if len(data) > maxDataWidth {
			maxDataWidth = len(data)
		}
	}

	fmt.Fprintf(os.Stdout, "Sum:     %*.*s\n", maxDataWidth, maxDataWidth, m["Sum"])
	fmt.Fprintf(os.Stdout, "Min:     %*.*s\n", maxDataWidth, maxDataWidth, m["Min"])
	fmt.Fprintf(os.Stdout, "Max:     %*.*s\n", maxDataWidth, maxDataWidth, m["Max"])
	fmt.Fprintf(os.Stdout, "Avg:     %*.*s\n", maxDataWidth, maxDataWidth, m["Avg"])
	fmt.Fprintf(os.Stdout, "Median:  %*.*s\n", maxDataWidth, maxDataWidth, m["Median"])
	fmt.Fprintf(os.Stdout, "Count:   %*.*s\n", maxDataWidth, maxDataWidth, m["Count"])

	if val, ok := m["Errors"]; ok {
		fmt.Fprintf(os.Stdout, "Errors:  %*.*s\n", maxDataWidth, maxDataWidth, val)
	}
}

func app(r io.Reader) {
	var dataL []datapoint
	sum := 0.0
	errorCount := 0
	maxDecPlaces := 0
	var min, max *datapoint

	scanner := bufio.NewScanner(r)

	// read in all lines from the input file (or stdin)
	for scanner.Scan() {
		var l = scanner.Text()
		l = strings.TrimSpace(l)

		if len(l) > 0 {
			dp, err := newDatapoint(l)
			if err != nil {
				errorCount++
				continue
			}

			dataL = append(dataL, *dp)
			sum += dp.cleanFloat
			if dp.decPlaces > maxDecPlaces {
				maxDecPlaces = dp.decPlaces
			}

			if min == nil || dp.cleanFloat < min.cleanFloat {
				min = dp
			}
			if max == nil || dp.cleanFloat > max.cleanFloat {
				max = dp
			}
		}

	}

	var m map[string]string
	if len(dataL) > 0 {
		midpoint := len(dataL) / 2
		m = map[string]string{
			"Sum":    fmt.Sprintf("%.*f", maxDecPlaces, sum),
			"Min":    fmt.Sprintf("%s", min.rawString),
			"Max":    fmt.Sprintf("%s", max.rawString),
			"Avg":    fmt.Sprintf("%.*f", maxDecPlaces, sum/float64(len(dataL))),
			"Median": fmt.Sprintf("%s", dataL[midpoint].rawString),
			"Count":  fmt.Sprintf("%d", len(dataL)),
		}
	} else {
		m = map[string]string{
			"Sum":    "0",
			"Min":    "0",
			"Max":    "0",
			"Avg":    "0",
			"Median": "0",
			"Count":  "0",
		}
	}

	if errorCount > 0 {
		m["Errors"] = fmt.Sprintf("%d", errorCount)
	}

	outputResults(m)
}

func main() {
	usage := `A command-line tool for converting a stream of numbers into total, averages, etc

Usage:
  statz [-f <infile>]
  statz -h | --help
  statz --version

Options:
  -f <infile>  The input file to be parsed [default: -]
`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/statz v0.0.1")

	if filepath, _ := opts.String("-f"); filepath == "-" {
		app(os.Stdin)
	} else {
		r, err := os.Open(filepath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer r.Close()
		app(r)
	}
}
