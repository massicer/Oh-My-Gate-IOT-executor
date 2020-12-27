package main

import (
	"os"
	"time"
	"strconv"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/logger"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/message_broker"
)

func main() {
	var logger logger.Logger = logger.Create_logger("iot-executor")
	logger.Info("Preparing to start...")

	
	logger.Info("Going to setup_broker...")
	err, broker := setup_broker(logger)
	handle_error(logger, err)
	logger.Info("Broker created")

	logger.Info("Going to start listening")
	err = broker.Start()
	handle_error(logger, err)
	logger.Info("Listening started")
	
	
}

func handle_error(logger logger.Logger, err error){
	if err != nil {
		logger.Panicf("Error starting, %s", err)
	}
}


func setup_broker(logger logger.Logger) (error, *message_broker.Broker){

	var ack_in_seconds, err = strconv.ParseInt(
		os.Getenv("ACK_TIME_IN_SECONDS"), 10, 64,
	)

	if err != nil {
		logger.Error("Cannot parse ack in seconds")
		return err, nil
	}

	var broker_config = message_broker.BrokerConfig{
		Subscription: os.Getenv("SUBSCRIPTION_NAME"),
		Topic: os.Getenv("TOPIC_NAME"),
		AckTimeInSeconds: time.Duration(ack_in_seconds),
		Project_id: os.Getenv("GCP_PROJECT_ID"),
	}

	var broker = message_broker.Broker{
		Config: broker_config,
		Logger: logger,
	}

	return nil, &broker
}