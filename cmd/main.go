package main

import (
	"github.com/massicer/Oh-My-Gate-IOT-executor/internal/logger"
)

func main() {
	var logger logger.Logger = logger.Create_logger("iot-executor")
	logger.Info("Preparing to start...")
}