package main

import (
	"os"
	"time"
	"strconv"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/iot_adapter"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/logger"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/message_broker"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/handler"
)

func main() {
	var logger logger.Logger = logger.Create_logger("iot-executor")
	logger.Info("Preparing to start...")

	var adapter_type = os.Getenv("ADAPTER_TYPE")
	logger.InfoF("Going to setup iot-adapter of type: %s", adapter_type)
	var adapter, error = iot_adapter.GetAdapter(adapter_type)
	handle_error(logger, error)
	logger.Info("Iot adapter configured")

	logger.Info("Going to setup handler")
	var handler = setup_handler(logger, adapter)
	logger.Info("Handler configured")
	
	logger.Info("Going to setup_broker...")
	err, broker := setup_broker(logger, handler)
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

func setup_handler(logger logger.Logger, adapter iot_adapter.Iot_adapter) *handler.BaseHandler {
	return &handler.BaseHandler{
		Adapter: adapter,
		Logger: logger,
	}
}


func setup_broker(logger logger.Logger, handler *handler.BaseHandler) (error, *message_broker.Broker){

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
		Handler: handler,
	}

	return nil, &broker
}