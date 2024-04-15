package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
    Name string `json:"name"`
}

func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
    for _, message := range sqsEvent.Records {
        var event MyEvent
        err := json.Unmarshal([]byte(message.Body), &event)
        if err != nil {
            log.Printf("Could not decode body: %v", err)
            continue
        }
        if &event == nil {
            log.Println("Received nil event")
            continue
        }
        message := fmt.Sprintf("Hello %s!", event.Name)
        log.Printf("Generated message: %s", message)
    }
    return nil
}

func main() {
    lambda.Start(HandleRequest)
}