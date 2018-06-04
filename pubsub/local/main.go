package main

import (
	"fmt"

	pbs "gitlab.rnd.saltosystems.com/cloud-architecture/messaging/pkg/pubsub/server"
	"github.com/lopezator/go_sandbox/pubsub/local/pkg"
)

var pubSubServer *pbs.Server

func main() {
	//Init vars
	pubSubServer, _ = pbs.NewServer("pubsub/local/config.yaml")
	_  = pubSubServer.PrepareServer()

	fmt.Println("Topic's mandanga...")
	pkg.CreateTopic()
	waitForEnter()
	pkg.GetTopic()
	waitForEnter()
	pkg.ListTopics()
	waitForEnter()
	pkg.ListTopicSubscriptions()
	waitForEnter()
	pkg.PublishTopic()
	waitForEnter()
	pkg.DeleteTopic()
	fmt.Println("Subscriptions's mandanga...")
}

func waitForEnter() {
	fmt.Println("Press the Enter Key to continue...")
	fmt.Scanln()
}
