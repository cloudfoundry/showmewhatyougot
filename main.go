package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/masters-of-cats/showmewhatyougot/statedetector"
)

func main() {
	state := *flag.String("state", "D", "Type of state to detect")
	pollingInterval := *flag.Duration("polling-interval", 10*time.Second, "Interval between process state checks")
	alertIntervalThreshold := *flag.Int("alert-interval-threshold", 15, "Number of checks before a process is considered in a persistent state")
	tracingEnabled := *flag.Bool("tracing-enabled", false, "Enable XFS Kernel tracing")
	processStateCounterBinaryPath := *flag.String("process-state-counter", "", "State process counter binary path")
	processStateReporterBinaryPath := *flag.String("process-state-reporter", "", "State process reporter binary path")
	xfsTraceBinaryPath := *flag.String("xfs-trace-path", "", "XFS Trace binary path")
	pidFilePath := *flag.String("pid-file-path", "", "Path to write out this process's pid file")

	if pidFilePath != "" {
		if err := ioutil.WriteFile(pidFilePath, []byte(strconv.Itoa(os.Getpid())), 0600); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write pid to '%s': %s\n", pidFilePath, err.Error())
			os.Exit(1)
		}
	}

	processStateCounter := statedetector.NewBinaryProcessStateCounter(processStateCounterBinaryPath)
	processStateReporter := statedetector.NewBinaryProcessStateReporter(processStateReporterBinaryPath)

	xfsTracer := statedetector.NewDummyXfsTracer()
	if tracingEnabled {
		xfsTracer = statedetector.NewBinaryXfsTracer(xfsTraceBinaryPath)
	}

	currentStateDetector := statedetector.NewCurrentStateDetector(state)
	persistentStateDetector := statedetector.NewPersistentStateDetector(alertIntervalThreshold, currentStateDetector)

	showMeWhatYouGot := statedetector.NewShowMeWhatYouGot(processStateCounter, processStateReporter, xfsTracer, persistentStateDetector)
	for true {
		err := showMeWhatYouGot.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		}
		time.Sleep(pollingInterval)
	}
}