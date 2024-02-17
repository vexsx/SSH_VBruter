package runner

import "time"

var (
	opt   *Options
	uhost []string
	vld   bool
)

const (
	version = "v0.1.2"
	banner  = `
             _
     ` + version + ` | |
     ___ ___| |__
    / __/ __| '_ \
    \__ \__ \ |_) |
    |___/___/_.__/

  Secure Shell Bruteforcer Edited By Vexsx
  infosec@kitabisa.com
  
`
	usage = `Usage:
  ssb

Options:
  port
     Port to connect to on the remote host (default 22).
  wordlist
     Path to wordlist file.
  timeout
     Connection timeout (default 30s).
  concurrent
     Concurrency/threads level (default 100).
  retries
     Specify the connection retries (default 1).
  output
     Save valid password to file.
  v
     Verbose mode.


`
	timeout = 30 * time.Second
)
