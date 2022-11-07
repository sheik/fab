//go:build exclude
// +build exclude

package main

import "github.com/sheik/fab"

var plan = fab.Plan{
	"clean": {
		Command: "docker rmi -f producer",
	},
	"build": {
		Command: "docker build . -t producer",
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
	"producer": {
		Command: "helm install producer ./helm/producer",
		Depends: "minikube",
		Check:   fab.ReturnZero("helm list | grep producer"),
	},
	"network": {
		Depends: "redis-cluster kafka-cluster producer",
	},
	"network-down": {
		Command: "helm uninstall redis-cluster kafka-cluster producer",
	},
	"test": {
		Command: "kubectl get pods",
		Depends: "build network",
		Default: true,
	},
}

func main() {
	fab.Run(plan)
}
