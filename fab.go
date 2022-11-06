//go:build exclude
// +build exclude

package main

import (
	"fmt"
	"github.com/sheik/fab"
)

var (
	nextTag = fab.IncrementMinorVersion(fab.GetVersion())
)

var plan = fab.Plan{
	"clean": {
		Command: "rm -rf $(ls ./cmd)",
	},
	"test": {
		Command: "go test ./...",
	},
	"build": {
		Command: "go build -o . ./...",
		Depends: "clean test",
		Default: true,
	},
	"tag": {
		Command: fmt.Sprintf(`
			echo "%s" > ./cmd/fab/buildinfo.txt
			git commit ./cmd/fab/buildinfo.txt -m "updating tag to %s"
			git tag %s
			git push origin %s
		`, nextTag, nextTag, nextTag, nextTag),
		Depends: "test",
	},
}

func main() {
	fab.Run(plan)
}
