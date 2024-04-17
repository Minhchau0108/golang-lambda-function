package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strings"
    "net/http"
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

        // Define the endpoint URL
        endpointURL := "https://contentful-cms-blog.vercel.app/api/webhook"

        // Prepare the payload for the POST HTTP Request
        payload, err := json.Marshal(event)
        if err != nil {
            log.Printf("Error marshalling event: %v", err)
            continue
        }
        // Print the payload
        log.Printf("Payload: %s", string(payload))
        

        // Send the HTTP POST request to endpoint URL
        _, err = http.Post(endpointURL, "application/json", strings.NewReader(string(payload)))
        if err != nil {
            log.Printf("Error sending POST request: %v", err)
            continue
        }

        log.Println("POST request sent successfully")
    }
    return nil
}

func main() {
    lambda.Start(HandleRequest)
}