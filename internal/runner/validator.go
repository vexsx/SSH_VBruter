package runner

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func (opt *Options) validate() error {
	var uhost []string

	green := "\033[32m"
	reset := "\033[0m"

	// Prompt user for host information
	fmt.Printf("%sEnter host like root@localhost :%s ", green, reset)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input) // Trim newline and any leading/trailing spaces

	if input == "" {
		return errors.New("Please define a target server")
	} else {
		uhost = strings.SplitN(input, "@", 2)
	}

	if len(uhost) < 2 {
		usr, err := user.Current()
		if err != nil {

			opt.user = "root"
		} else {
			opt.user = usr.Username
		}
		opt.host = uhost[0]
	} else {
		opt.user = uhost[0]
		opt.host = uhost[1]
	}

	if opt.wordlist != "" {
		f, err := os.Open(opt.wordlist)
		if err != nil {
			return err
		}
		opt.list = f
	} else {
		return errors.New("No wordlist file provided")
	}

	if opt.output != "" {
		opt.file, err = os.OpenFile(opt.output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func (opt *Options) listScanner() *bufio.Scanner {
	return bufio.NewScanner(opt.list)
}

func (opt *Options) Close() {
	if opt.list != nil {
		_ = opt.list.Close()
	}

	if opt.file != nil {
		_ = opt.file.Close()
	}
}
