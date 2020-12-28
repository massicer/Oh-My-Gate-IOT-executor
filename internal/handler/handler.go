package handler

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/entities"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/iot_adapter"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/logger"
)

type Handler interface {
	Handle(message []byte) error
}

type BaseHandler struct {
	Adapter iot_adapter.Iot_adapter
	Logger logger.Logger
}

func (h *BaseHandler) Handle(message []byte) error {

	h.Logger.Infof("Preparig to handle message: %s", string(message))

    var msg entities.OpenMessage
	err := json.Unmarshal(message, &msg)

	if err != nil {
		h.Logger.Errorf("Error: %s during parsing message: %s", string(message), err)
		return err
	}
	h.Logger.Infof("Successfully parsed message: %s", msg)

	if !msg.Is_open_action() {
		var error_string = fmt.Sprintf("Error: message: %v is not open action but is: %s", msg, msg.Action)
		h.Logger.Error(error_string)
		return errors.New(error_string)
	}

	h.Logger.InfoF("Message: %s is open action. Preparing to trigger open", msg)
	err = h.Adapter.Open(msg)
	if err != nil {
		var error_string = fmt.Sprintf("Error: %s during open for message: %v", err, msg,)
		h.Logger.Error(error_string)
		return errors.New(error_string)
	}
	h.Logger.InfoF("Opened successfully for msg: %s", msg)

	return nil
}