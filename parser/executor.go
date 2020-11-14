package parser

import "os"

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if in == "exit" {
		os.Exit(0)
	}
}
