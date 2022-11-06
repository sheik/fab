//go:build exclude
// +build exclude

package main

import "github.com/sheik/fab"

var plan = fab.Plan{
	"clean": {
		Command: "rm -rf $(ls cmd)",
		Help:    "clean binaries",
	},
	"build": {
		Command: "go build ./...",
		Depends: "clean test",
		Default: true,
		Help:    "build binaries",
	},
	"test": {
		Command: "go test ./...",
		Depends: "clean",
		Help:    "run bdd tests",
	},
}

func main() {
	fab.Run(plan)
}
