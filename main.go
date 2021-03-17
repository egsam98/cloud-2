package main

import (
	"os"

	"github.com/egsam98/cloud-2/cmd"
)

func main() {
	if err := cmd.Raid.Execute(); err != nil {
		os.Exit(1)
	}
}
