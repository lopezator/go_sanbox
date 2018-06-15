package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"runtime"
)

var (
	projectID           = "sergeant-320195"
	topic               = "messaging-pub"
	subscription        = "messaging-sub"
	defaultMaxMessages  = 10000
	defaultMaxExtension = 60 * time.Second
)

// Subscriber holds GCP PubSub subscription info
type Subscriber struct {
	subscription *pubsub.Subscription
	cancel       func()
	err          error
}

// Publisher holds GCP PubSub topic info
type Publisher struct {
	topic *pubsub.Topic
}

func main() {
	// Creation of pubsub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("create client failed")
	}

	// Creation of pubsub subscription
	sub := client.Subscription(subscription)
	sub.ReceiveSettings = pubsub.ReceiveSettings{
		MaxExtension:           defaultMaxExtension,
		MaxOutstandingMessages: defaultMaxMessages,
	}

	// Creation of subscriber, and pull messages (if any)
	// Launch in a goroutine and keep listening for new messages
	s := Subscriber{
		subscription: sub,
	}
	go s.Pull()

	// Creation of a subscriber on the defined topic
	p := Publisher{
		client.Topic(topic),
	}

	// Send a test message to the defined topic
	res := p.topic.Publish(ctx, &pubsub.Message{
		Data: []byte("My cool message"),
	})
	id, err := res.Get(ctx)

	fmt.Printf("Published a message; msg ID: %v\n", id)

	runtime.Goexit()
}

// Start starts the subscriber, opening the channel.
// Waits for an explicit Stop() call to finish.
func (s *Subscriber) Start() <-chan pubsub.Message {
	output := make(chan pubsub.Message)
	go func(s *Subscriber, output chan pubsub.Message) {
		defer close(output)

		ctx := context.Background()
		ctx, s.cancel = context.WithCancel(ctx)
		err := s.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			output <- pubsub.Message{
				Data: msg.Data,
			}
		})
		if err != nil {
			s.err = err
		}
	}(s, output)

	return output
}

// Stop the subscriber, closing the channel that was returned by Start.
func (s *Subscriber) Stop() {
	if s.cancel != nil {
		// TODO(david.lopez) Remove when issue in pubsub-go has been fixed:
		// https://github.com/GoogleCloudPlatform/google-cloud-go/issues/1030
		time.Sleep(2 * time.Second)
		s.cancel()
	}
}

// Pull pulls messages from a subscription, if any
func (s *Subscriber) Pull() {
	for {
		c := s.Start()
		msg, ok := <-c
		if !ok {
			fmt.Println(s.err)
		} else {
			fmt.Println("message received", string(msg.Data))
			msg.Ack()
			s.Stop()
		}
	}
}
