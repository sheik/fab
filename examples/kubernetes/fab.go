//go:build exclude
// +build exclude

package main

import "github.com/sheik/fab"

var (
	buildContainer = fab.Container(image).Mount("$PWD", "/code").Env("CGO_ENABLED", "0")
	image          = "builder:latest"
)

var plan = fab.Plan{
	"clean": {
		Command: "rm -rf fab",
	},
	"build-container": {
		Command: "docker build . -f builder/Dockerfile -t" + image,
		Check:   fab.ImageExists(image),
	},
	"build": {
		Command: buildContainer.Run("go build ./cmd/fab"),
		Depends: "clean build-container",
	},
	"minikube": {
		Command: "minikube start",
		Check:   fab.ReturnZero("minikube status"),
	},
	"redis-cluster": {
		Command: "helm install redis-cluster bitnami/redis-cluster",
		Depends: "minikube",
		Check:   fab.ReturnZero("helm list | grep redis-cluster"),
	},
	"kafka-cluster": {
		Command: "helm install kafka-cluster --set replicaCount=3 bitnami/kafka",
		Depends: "minikube",
		Check:   fab.ReturnZero("helm list | grep kafka-cluster"),
	},
	"network": {
		Depends: "redis-cluster kafka-cluster",
	},
	"test": {
		Command: "docker ps",
		Depends: "build network",
		Default: true,
	},
}

func main() {
	fab.Run(plan)
}
