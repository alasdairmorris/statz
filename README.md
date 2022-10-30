# statz

A command-line tool for converting a stream of numbers into total, averages, etc

## Installation

`statz` will run on most Linux and Mac OS X systems.

To install it, just find the appropriate one-liner below - based on the destination O/S and architecture - and copy-paste it into your terminal.

Feel free to change the install dir - `$HOME/bin` in the examples below - to be something more appropriate for your needs.

### Linux (32-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/statz/releases/latest/download/statz-linux-386.tar.gz | tar -zxf - -C $HOME/bin
```

### Linux (64-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/statz/releases/latest/download/statz-linux-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Intel)

```
curl -s -L -o - https://github.com/alasdairmorris/statz/releases/latest/download/statz-darwin-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o - https://github.com/alasdairmorris/statz/releases/latest/download/statz-darwin-arm64.tar.gz | tar -zxf - -C $HOME/bin
```

### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

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
