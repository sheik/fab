//go:build exclude
// +build exclude

package main

import (
	"fmt"
	"github.com/sheik/fab"
)

var (
	buildContainer = fab.Container(image).Mount("$PWD", "/code").Env("CGO_ENABLED", "0")
	image          = "builder:latest"
	currentTag     = fab.GetVersion()
	nextTag        = fab.IncrementMinorVersion(currentTag)
)

var plan = fab.Plan{
	"clean": {
		Command: "rm -rf fab",
	},
	"build-container": {
		Command: "docker build . -f builder/Dockerfile -t" + image,
		Check:   fab.Check{fab.ImageExists, image},
	},
	"build": {
		Command: buildContainer.Run("go build ./cmd/fab"),
		Depends: "clean build-container",
	},
	"test": {
		Command: "docker ps",
		Depends: "build",
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
