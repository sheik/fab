package main

import (
	_ "embed"
	"github.com/sheik/fab"
	"github.com/sheik/fab/pkg/log"
	"os"
	"strings"
)

//go:embed templates/fab.go
var initFile []byte

func main() {
	var args string

	if len(os.Args) > 1 {
		args = strings.Join(os.Args[1:], " ")

		switch os.Args[1] {
		case "update":
			fab.Run(fab.Plan{"update": fab.UpdateStep})
			return
		case "init":
			err := os.WriteFile("fab.go", initFile, 0644)
			if err != nil {
				log.Error("unable to create fab.go: %s", err)
				return
			}
			fab.Exec("go mod tidy")
			return
		}

	}

	fab.InteractiveCommand("go run fab.go " + args)
}
