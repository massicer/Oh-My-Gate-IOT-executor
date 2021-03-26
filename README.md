# Oh My Gate IOT Executor

[![Build Status](https://travis-ci.org/massicer/Oh-My-Gate-IOT-executor.svg?branch=main)](https://travis-ci.org/massicer/Oh-My-Gate-IOT-executor)

Receive msg related to IOT actions to be executed via google [Pub/Sub](https://cloud.google.com/pubsub) and handles them.

**_important_**
This service doesn't create a topic if not exists. Please be sure to create it in GCP.
The subscription absence is handled.

## Expected message shape

```
{
	"action": "Open",
	"id": 2
}
```

## Available IOT adapters

This service is able to support the following adapters:

- `Standard out`: simply prints msg to standard_out.
- `gpio`: interacts with a gpio board (example: Raspberry Pi)
  The adapter can be configured using the environment variable: `ADAPTER_TYPE`

### Envs needed

- `GOOGLE_APPLICATION_CREDENTIALS`: Path where gcp service account is placed. [Here](https://cloud.google.com/iam/docs/creating-managing-service-accounts) you can find a way to create it.

- `GCP_PROJECT_ID`: Name of project id.

- `SUBSCRIPTION_NAME`: Name of the subscription to use to listen for msgs.

- `TOPIC_NAME`: Name of the topic to use to listen for msgs.

- `ACK_TIME_IN_SECONDS`: Ack time to use in seconds by msg broker

- `ADAPTER_TYPE`: Adapter type to use to handle messages. Available values are:
  - `standard_out`
  - `gpio`

---

## CI/CD:

This repo has a set of pipeline in order to produce artifacts in an automatic way: [Docker hub image](https://hub.docker.com/r/massicer/oh-my-gate-iot-executor])

- on branch != main an image with tag `dev` is published
- on branch == main an image with tag `latest` is pubblished
- on tag an image with `tag-name` is published

## How to publish a docker image for different architectures

Perform `make cross_build`

### How to release a new service version

- `git checkout master`
- `git tag {your tag}`
- `git push --follow-tags`
