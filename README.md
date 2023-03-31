# statz

A command-line tool for converting a stream of numbers into total, averages, etc

## Installation

`statz` will run on most Linux, MacOS and Windows systems.

To install it, just `cd` into the directory in which you wish to install it and then copy-paste the appropriate one-liner from below.

### Linux (32-bit)

```
curl -s -L -o statz https://github.com/alasdairmorris/statz/releases/latest/download/statz-linux-386 && chmod +x statz
```

### Linux (64-bit)

```
curl -s -L -o statz https://github.com/alasdairmorris/statz/releases/latest/download/statz-linux-amd64 && chmod +x statz
```

### Mac OS X (Intel)

```
curl -s -L -o statz https://github.com/alasdairmorris/statz/releases/latest/download/statz-darwin-amd64 && chmod +x statz
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o statz https://github.com/alasdairmorris/statz/releases/latest/download/statz-darwin-arm64 && chmod +x statz
```

### Windows (32-bit)

```
curl -s -L -o statz.exe https://github.com/alasdairmorris/statz/releases/latest/download/statz-windows-386.exe
```

### Windows (64-bit)

```
curl -s -L -o statz.exe https://github.com/alasdairmorris/statz/releases/latest/download/statz-windows-amd64.exe
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

## License

[MIT](LICENSE)
