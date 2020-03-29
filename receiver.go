package receiver

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	cepubsub "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
)

// published data struct.
type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

func Receiver(ctx context.Context, msg *pubsub.Message) error {
	fmt.Printf("[INFO] message received: %v\n", msg)
	fmt.Printf("[INFO] message.Attributes: %s\n", msg.Attributes)
	fmt.Printf("[INFO] message.Data: %s\n", msg.Data)

	// cloudevents pubsubパッケージ内のMessage型に変換する
	var mess cepubsub.Message
	mess.Data = msg.Data
	mess.Attributes = msg.Attributes
	fmt.Printf("[INFO] CloudEventsVersion: %s\n", mess.CloudEventsVersion())

	data := &Example{}
	json.Unmarshal([]byte(mess.Data), &data)
	fmt.Printf("[INFO] Example.Sequence: %d\n", data.Sequence)
	fmt.Printf("[INFO] Example.Message: %s\n", data.Message)

	return nil
}
