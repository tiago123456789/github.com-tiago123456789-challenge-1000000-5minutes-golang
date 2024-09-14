package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1"),
			CredentialsChainVerboseErrors: aws.Bool(true)},
		Profile: "tiago",
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}

	// Create a new instance of the SQS service client
	svc := sqs.New(sess)

	// Specify the SQS queue URL
	queueURL := "https://sqs.us-east-1.amazonaws.com/507403822990/challenge_1000000_5minutes" // Replace with your queue URL

	// Create a new message to send
	messageBody := "Hello from Go!"

	// Create the message parameters
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueURL),
	}

	startAt := time.Now()

	var messages []*sqs.SendMessageBatchRequestEntry
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		messages = append(messages, &sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(uuid.New().String()),
			MessageBody: &messageBody,
		})

		if len(messages) == 10 {
			fmt.Println("passed on here")

			_, err := svc.SendMessageBatch(&sqs.SendMessageBatchInput{
				Entries:  messages,
				QueueUrl: params.QueueUrl,
			})

			if err != nil {
				log.Fatalf("failed to send message, %v", err)
			}

			messages = []*sqs.SendMessageBatchRequestEntry{}
		}

	}

	fmt.Println(time.Since(startAt))
}
