package iot_adapter

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/entities"
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	STANDARD_OUT_ADAPTER  = "standard_out"
	GPIO_ADAPTER          = "gpio"
	SLEEP_TIME_IN_SECONDS = 3 * time.Second
)

var ACTION_MUTEX = &sync.Mutex{}

type Iot_adapter interface {
	Open(msg entities.OpenMessage) error
}

type Sdout_iot_adapter struct {
	W io.Writer
}

type GpioAdapter struct {
}

func (s *Sdout_iot_adapter) Open(msg entities.OpenMessage) error {
	defer ACTION_MUTEX.Unlock()
	ACTION_MUTEX.Lock()
	fmt.Fprintf(s.W, "Going to open iot device with id: %d", msg.Id)
	return nil
}

func (s *GpioAdapter) Open(msg entities.OpenMessage) error {
	defer ACTION_MUTEX.Unlock()
	defer rpio.Close()

	ACTION_MUTEX.Lock()
	err := rpio.Open()

	if err != nil {
		return err
	}

	pin := rpio.Pin(msg.Id)
	pin.Output()
	pin.High()
	time.Sleep(SLEEP_TIME_IN_SECONDS)
	pin.Low()
	return nil
}

func GetAdapter(adapter_type string) (Iot_adapter, error) {

	switch adapter_type {
	case STANDARD_OUT_ADAPTER:
		return &Sdout_iot_adapter{W: os.Stdout}, nil
	case GPIO_ADAPTER:
		return &GpioAdapter{}, nil

	default:
		return nil, errors.New("Adapter type not found")
	}
}
