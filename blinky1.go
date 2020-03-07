package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"time"

	"github.com/tinygo-org/tinygo-stm32/machine/gpio"
)

func main() {

	led1 := gpio.PA5
	led2 := gpio.PA0

	led1.Configure(gpio.PinConfig{Mode: gpio.PinOutputPP, Pull: gpio.PinNoPull})
	led2.Configure(gpio.PinConfig{Mode: gpio.PinOutputPP, Pull: gpio.PinNoPull})

	go func() {
		for {
			led1.High()
			time.Sleep(500 * time.Millisecond)
			led1.Low()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		led2.High()
		time.Sleep(50 * time.Millisecond)
		led2.Low()
		time.Sleep(50 * time.Millisecond)
	}
}
