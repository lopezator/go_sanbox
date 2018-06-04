package pkg

import (
	"fmt"

	"gitlab.rnd.saltosystems.com/cloud-architecture/messaging/pkg/pubsub/server"
)

var topicSrv = server.TopicServer{}

func CreateTopic() {
	fmt.Println("Creating topic...")
	topic, err := topicSrv.Create("test-topic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topic)
}

func DeleteTopic() {
	fmt.Println("Deleting topic...")
	topic, err := topicSrv.Delete("test-topic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topic)
}

func GetTopic() {
	fmt.Println("Getting topic...")
	topic, err := topicSrv.Get("test-topic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topic)
}

func ListTopics() {
	fmt.Println("Listing topics...")
	topics, err := topicSrv.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topics)
}

func ListTopicSubscriptions() {
	fmt.Println("Listing topics subscriptions...")
	topics, err := topicSrv.ListSubscription("test-topic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topics)
}

func PublishTopic() {
	fmt.Println("Publish topic...")
	topics, err := topicSrv.Publish("test-topic", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(topics)
}