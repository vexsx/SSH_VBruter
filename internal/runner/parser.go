package runner

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
)

// Options declare its options
type Options struct {
	concurrent int
	wordlist   string
	retries    int
	verbose    bool
	timeout    time.Duration
	output     string
	port       int
	list       io.ReadCloser
	file       *os.File

	user string
	host string
}

// Parse arguments
func Parse() *Options {
	opt := &Options{
		timeout: 30 * time.Second, // Example default value for the timeout.
	}
	green := "\033[32m"
	reset := "\033[0m"

	showBanner()
	//showUsage()

	fmt.Printf("%sEnter the port (default: 22)%s : ", green, reset)
	fmt.Scanln(&opt.port)

	fmt.Printf("%sEnter the number of retries (default: 1)%s : ", green, reset)
	fmt.Scanln(&opt.retries)

	fmt.Printf("%sEnter the number of concurrent connections (default: 1)%s : ", green, reset)
	fmt.Scanln(&opt.concurrent)

	var verboseInput string
	fmt.Printf("%sEnter verbose mode (true/false)%s : ", green, reset)
	fmt.Scanln(&verboseInput)
	opt.verbose, _ = strconv.ParseBool(verboseInput)

	fmt.Printf("%sEnter the output file path%s : ", green, reset)
	fmt.Scanln(&opt.output)

	fmt.Printf("%sEnter the wordlist file path :%s ", green, reset)
	fmt.Scanln(&opt.wordlist)

	var timeoutInput string
	fmt.Printf("%sEnter the timeout duration (in seconds) :%s ", green, reset)
	fmt.Scanln(&timeoutInput)
	if timeout, err := strconv.Atoi(timeoutInput); err == nil {
		opt.timeout = time.Duration(timeout) * time.Second
	}

	if err := opt.validate(); err != nil {
		fmt.Printf("Error! %s.\n", err.Error())
		os.Exit(1)
	}

	return opt
}

func showBanner() {
	_, err := fmt.Fprint(os.Stderr, aurora.Bold(aurora.Cyan(banner)))
	if err != nil {
		log.Fatal(err)
	}
}

func showUsage() {
	_, err := fmt.Fprint(os.Stderr, aurora.Bold(aurora.Cyan(usage)))
	if err != nil {
		log.Fatal(err)
	}

}
