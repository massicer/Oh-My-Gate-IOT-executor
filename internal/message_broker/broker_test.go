package message_broker

import (
	"testing"
	"time"
)

func TestBrokerConfig(t *testing.T) {

	var subscription string = "My-subscription"
	var topic string = "my-topic"
	var ack_in_seconds = time.Duration(10)
	var project_id = "my-project"

	var config = BrokerConfig{
		Subscription: subscription,
		Topic: topic,
		AckTimeInSeconds: ack_in_seconds,
		Project_id: project_id,
	}

	if config.Subscription != subscription{
		t.Errorf("Subscription mismatch. Wanted %s, got %s", config.Subscription, subscription)
	}

	if config.Topic != topic{
		t.Errorf("Topic mismatch. Wanted %s, got %s", config.Topic, topic)
	}

	if config.AckTimeInSeconds != ack_in_seconds{
		t.Errorf("ack_in_seconds mismatch. Wanted %s, got %s", config.AckTimeInSeconds, ack_in_seconds)
	}

	if config.Project_id != project_id{
		t.Errorf("project_id mismatch. Wanted %s, got %s", config.Project_id, project_id)
	}
}