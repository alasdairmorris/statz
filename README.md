# statz

A command-line tool for converting a stream of numbers into total, averages, etc

## Installation

Simply download the appropriate pre-compiled binary for your system from [the release page](https://github.com/alasdairmorris/statz/releases).

Or, if you have Go installed and prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/statz@latest
```

## Usage

```
A command-line tool for converting a stream of numbers into total, averages, etc

Usage:
  statz [-f <infile>]
  statz -h | --help
  statz --version

Options:
  -f <infile>  The input file to be parsed [default: -]
```

## Examples

```
$ echo -e "1\n2\n145" | statz
Sum:     148
Min:       1
Max:     145
Avg:      49
Median:    2
Count:     3
```
