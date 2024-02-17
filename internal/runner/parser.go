package runner

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
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
	opt = &Options{}
	opt.timeout = timeout

	flag.IntVar(&opt.port, "p", 22, "")
	flag.IntVar(&opt.retries, "r", 1, "")
	flag.IntVar(&opt.concurrent, "c", 100, "")
	flag.BoolVar(&opt.verbose, "v", true, "")
	flag.StringVar(&opt.output, "o", "", "")
	flag.StringVar(&opt.wordlist, "w", "", "")
	flag.DurationVar(&opt.timeout, "t", opt.timeout, "")

	flag.Usage = func() {
		showBanner()
		showUsage()
		_, fprint := fmt.Fprint(os.Stderr, usage)
		if fprint != nil {
			return
		}

	}
	flag.Parse()

	showBanner()
	showUsage()
	fmt.Scanln()

	if err := opt.validate(); err != nil {
		gologger.Fatalf("Error! %s.", err.Error())
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
