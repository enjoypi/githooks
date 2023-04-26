package main

import (
	"os"

	"git.tap4fun.com/k2/githooks/verify-commit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
