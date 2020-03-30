package receiver

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	cepubsub "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
)

// Model is published data struct.
type Model struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

// Receiver parse payload to Model
func Receiver(ctx context.Context, msg *pubsub.Message) error {
	fmt.Printf("[INFO] message received: %v\n", msg)
	fmt.Printf("[INFO] message.Attributes: %s\n", msg.Attributes)
	fmt.Printf("[INFO] message.Data: %s\n", msg.Data)

	// convert pubsub.Message to cepubsub.Message
	var cemsg cepubsub.Message
	cemsg.Data = msg.Data
	cemsg.Attributes = msg.Attributes
	fmt.Printf("[INFO] CloudEventsVersion: %s\n", cemsg.CloudEventsVersion())

	// get data(type Model) from cepubsub.Message
	data := &Model{}
	json.Unmarshal([]byte(cemsg.Data), &data)
	fmt.Printf("[INFO] Model.Sequence: %d\n", data.Sequence)
	fmt.Printf("[INFO] Model.Message: %s\n", data.Message)

	return nil
}
