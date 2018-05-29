#!/bin/bash

TOPIC_NAME="my-topic"
ACK_DEADLINE=600
SUB_1_NAME="my-subscription-1"
SUB_2_NAME="my-subscription-2"

# Create sandbox SMS topic
echo "1.- Creating a topic...\n"
gcloud pubsub topics create $TOPIC_NAME
read -p "Press enter to continue"

# Create sandbox SMS topic subscriptions
echo "2.- Creating Subscriptions...\n"
gcloud pubsub subscriptions create $SUB_1_NAME --topic $TOPIC_NAME --ack-deadline=$ACK_DEADLINE
gcloud pubsub subscriptions create $SUB_2_NAME --topic $TOPIC_NAME --ack-deadline=$ACK_DEADLINE
read -p "Press enter to continue"

# Push 10 messages to topic
echo "3.- Push 10 messages to the topic...\n"
for i in {1..10}
do
   gcloud pubsub topics publish $TOPIC_NAME --message "Hello world #$i!"
done
read -p "Press enter to continue"

# Pull messages
echo "3.- Pull 10 messages 5 by each subscriber...\n"
gcloud pubsub subscriptions pull --auto-ack --limit=5 $SUB_1_NAME
gcloud pubsub subscriptions pull --auto-ack --limit=5 $SUB_2_NAME
read -p "Press enter to continue"

# Delete topic
echo "4.- Deleting the topic...\n"
gcloud pubsub topics delete $TOPIC_NAME
read -p "Press enter to continue"

# Delete subscriptions
echo "5.- Deleting the subscriptions...\n"
gcloud pubsub subscriptions delete $SUB_1_NAME
gcloud pubsub subscriptions delete $SUB_2_NAME
echo "Proccess ended!\n"
