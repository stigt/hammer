package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

var lastGood string
var best uint64 = 0

func done(rc int) {
	fmt.Printf("\n* %s\n", lastGood)
	fmt.Printf("best run %v\n", best)
	os.Exit(rc)
}

func main() {

	urlPtr := flag.String("url", "http://4.3.2.1/img/100KB",
		"the url to use")
	startRatePtr := flag.Uint64("start", 1000, "starting rate/second")
	incrementPtr := flag.Uint64("increment", 1000,
		"increament each iteration by")
	durationPtr := flag.Uint64("duration", 10,
		"duration in seconds per attack")
	flag.Parse()

	duration := time.Duration(*durationPtr) * time.Second
	rate := *startRatePtr

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		done(1)
	}()

	var errThold uint64 = 1
	var sleepDelay time.Duration = 3 * time.Minute
	for {
		if rate == 0 {
			fmt.Fprintln(os.Stderr, "Invalid rate", rate)
			os.Exit(1)
		}

		fmt.Printf("\n# Trying rate %v, iteration %v\n",
			rate, *incrementPtr)

		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: "GET",
			URL:    *urlPtr,
		})
		attacker := vegeta.NewAttacker(vegeta.KeepAlive(false))

		var metrics vegeta.Metrics
		var errs uint64 = 0
		var good uint64 = 0
		var attempts = rate * *durationPtr

		pause := false
		for res := range attacker.Attack(targeter, rate, duration) {
			metrics.Add(res)

			if res.Code == 200 {
				good++
			} else {
				if false {
					fmt.Printf("Error: code %v [%v]\n",
						res.Code, res.Error)
				}
				errs++
				if errs >= errThold {
					attacker.Stop()
					pause = true
					break
				}
			}

		}
		metrics.Close()

		line := fmt.Sprintf("rate %5v success %4.1f%% errs %v "+
			"99th percentile: %s",
			rate, metrics.Success*100, len(metrics.Errors),
			metrics.Latencies.P99)

		if errs >= errThold {
			*incrementPtr /= 2
			rate -= *incrementPtr
			percent := float64(good) / float64(attempts)
			fmt.Printf("- Requests %v, ok %v  %4.1f%%\n",
				attempts, good, percent*100)
			fmt.Printf("  Dropping rate to %v and increment "+
				"by %v\n",
				rate, *incrementPtr)
			if *incrementPtr < 250 || rate < best {
				done(0)
			}
		} else {
			if rate > best {
				best = rate
			}
			fmt.Printf("+ %s\n", line)
			lastGood = line
			rate += *incrementPtr
		}
		if pause {
			fmt.Printf("  Pausing %v\n", sleepDelay)
			time.Sleep(sleepDelay)
			sleepDelay += 60 * time.Second
			pause = false
		}
	}
}
