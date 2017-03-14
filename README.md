# Description

This is a program that makes a system call to start an Iperf server. It also
starts randomly an Iperf client that communicates with a random address. The
target addresses are hard-coded as a map.

# Installation

To compile the binary:

```bash
go build ./iperf-server-client-demo.go
```

If you wish, you could also install it using the go tools:

```bash
go get github.com/JuanCabre/iperf-server-client-demo.go
```

To build the binary for an Odroid, run:

```bash
GOOS=linux GOARCH=arm GOARM=7 go build -v ./server-client-demo.go
```

# Usage

Help:

```
./iperf-server-client-demo --help

Usage of ./iperf-server-client-demo:
  -max int
        Max value in seconds of the timer for calling the iperf client (default 7)
  -min int
        Min value in seconds of the timer for calling the iperf client (default 5)
```

Run with the default values:

```bash
./iperf-server-client-demo
```

Run with custom values for the min and max value of the iperf client timer:

```bash
iperf-server-client-demo.go --max 2 --min 1
```

To see the debug information, launch the program with the `DEBUG` environment
variable set to `DEBUG="*"`. For example:

```bash
env DEBUG="*" ./iperf-server-client-demo
```