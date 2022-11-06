package main

import (
	"github.com/sheik/fab"
	"os"
	"strings"
)

func main() {
	var args string

	if len(os.Args) > 1 {
		args = strings.Join(os.Args[1:], " ")
		if os.Args[1] == "update" {
			fab.Run(fab.Plan{"update": fab.UpdateStep})
			return
		}
	}

	fab.InteractiveCommand("go run main.go " + args)
}
