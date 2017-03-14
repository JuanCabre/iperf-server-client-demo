package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	dbg "github.com/JuanCabre/go-debug"
)

var debugT = dbg.Debug("Timer")
var debugAddr = dbg.Debug("Address")

// Choose the minimun and maximun value for the random timer that starts the
// iperf client
var (
	min      int
	max      int
	minTimer time.Duration
	maxTimer time.Duration
)

func init() {
	flag.IntVar(&min, "min", 5,
		"Min value in seconds of the timer for calling the iperf client")
	flag.IntVar(&max, "max", 7,
		"Max value in seconds of the timer for calling the iperf client")
	flag.Parse()
	if min <= 0 || max <= 0 || min > max {
		log.Fatal("The given values for the timer were invalid")
	}

	minTimer = time.Duration(min) * time.Second
	maxTimer = time.Duration(max) * time.Second
}

// Fill the map of addresses with the data
var addresses = map[string]string{
	"0,0,0": "127.0.0.1",
	"0,0,1": "127.0.0.1",
	"0,3,1": "127.0.0.1",
	"0,1,1": "127.0.0.1",
}

func main() {

	// Seed the rand
	rand.Seed(time.Now().UnixNano())

	// Start the Iperf server
	go startIperfServer()

	timer := time.NewTimer(randomDuration())
	// Start an Iperf client conecting randomly to an address
	for {
		select {
		case <-timer.C:

			// Start iperf client
			startIperfClient(randomAddress())

			// Reset timer
			timer.Reset(randomDuration())
		}
	}
}

func startIperfServer() {
	cmd := exec.Command("iperf", "-s")
	// Attach the Stdout and Stderr of the iperf server to the os
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error starting iperf server:", err)
	}
}

func startIperfClient(addr string) {
	cmd := exec.Command("iperf", "-c", addr)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error starting iperf client: ", err)
	}
}

// Choose a random duration between minTimer and maxTimer
func randomDuration() time.Duration {
	r := (rand.Int63n(maxTimer.Nanoseconds()-minTimer.Nanoseconds()) +
		minTimer.Nanoseconds())
	debugT("Reseting timer to %v", time.Duration(r))
	return time.Duration(r)
}

// Choose a random address from the list
func randomAddress() string {
	var k string
	i := rand.Intn(len(addresses))
	for k = range addresses {
		if i == 0 {
			break
		}
		i--
	}
	debugAddr("Random address chosen: %v -- %v", k, addresses[k])
	return addresses[k]
}
