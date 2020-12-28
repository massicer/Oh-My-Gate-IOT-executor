package iot_adapter

import (
	"fmt"
	"io"
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/entities"
)


type Iot_adapter interface {
	Open(msg entities.OpenMessage) error
}

type Sdout_iot_adapter struct {
	W io.Writer
}

func (s *Sdout_iot_adapter) Open(msg entities.OpenMessage) error {
	fmt.Fprintf(s.W, "Going to open iot device with id: %d", msg.Id)
    return nil
}