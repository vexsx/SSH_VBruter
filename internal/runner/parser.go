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

	showBanner()
	showUsage()

	fmt.Println("Enter the port:")
	fmt.Scanln(&opt.port)

	fmt.Println("Enter the number of retries:")
	fmt.Scanln(&opt.retries)

	fmt.Println("Enter the number of concurrent connections:")
	fmt.Scanln(&opt.concurrent)

	var verboseInput string
	fmt.Println("Enter verbose mode (true/false):")
	fmt.Scanln(&verboseInput)
	opt.verbose, _ = strconv.ParseBool(verboseInput)

	fmt.Println("Enter the output file path:")
	fmt.Scanln(&opt.output)

	fmt.Println("Enter the wordlist file path:")
	fmt.Scanln(&opt.wordlist)

	var timeoutInput string
	fmt.Println("Enter the timeout duration (in seconds):")
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
