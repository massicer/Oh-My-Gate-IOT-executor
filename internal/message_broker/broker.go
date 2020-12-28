package message_broker

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"context"
	"time"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/logger"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/handler"
)


type BrokerConfig struct {
	Subscription string `json:"subscription"`
	Topic string `json:"topic"`
	AckTimeInSeconds time.Duration `json:"ack_time_in_seconds"`
	Project_id string `json:"project_id"`
}



type Broker struct {
	Config BrokerConfig
	subscription *pubsub.Subscription
	Logger logger.Logger
	Handler handler.Handler
}

type MessageBroker interface {
	Start() error
}

func (b *Broker) Start() error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, b.Config.Project_id)
	if err != nil {
		return err
	}

	topic := client.Topic(b.Config.Topic)
	if topic == nil {
		b.Logger.Errorf("Error during retrieving topic with name: %s", topic)
		return fmt.Errorf("Error during retrieving topic with name %s", b.Config.Topic)
	}

	sub := client.Subscription(b.Config.Subscription)

	ok, err := sub.Exists(ctx)
	if err != nil {
		b.Logger.Error("Error during checking existence of subscription")
		return err
	}
	if !ok {
		sub, err = client.CreateSubscription(ctx, b.Config.Subscription, pubsub.SubscriptionConfig{
			Topic:            topic,
			AckDeadline:       b.Config.AckTimeInSeconds * time.Second,
			ExpirationPolicy: 25 * time.Hour,
		})

		if err != nil {
			b.Logger.Error("Error during creation of subscription")
			return err
		}
	} else {
		b.Logger.Info("Subscription already exists")
	}

	
	b.subscription = sub
	b.Logger.Info("Broker started successfully")
	
	b.Logger.Info("Preparing to start listening")
	err = b.subscription.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		b.Logger.Infof("Got message: %s", m.Data)
		_ = b.Handler.Handle(m.Data)
		m.Ack()
	})
	if err != nil {
		return err
	}
	b.Logger.Info("Started to listening")
	return nil
}
