package main

import "SSH_VBruter/internal/runner"

func main() {

	opt := runner.Parse()
	runner.New(opt)

}
