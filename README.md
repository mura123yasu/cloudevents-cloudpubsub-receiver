# cloudevents-cloudpubsub-receiver
Receive cloud events from cloud pubsub

## Deploy app to Cloud Functions
```sh
gcloud functions deploy <YOUR FUNCTION NAME> --project <YOUR GCP PROJECT> \
  --entry-point Receiver \
  --trigger-topic <YOUR PUBSUB TOPIC> \
  --runtime go113
```