package main

import "gopkg.in/bendahl/uinput.v1"
import "time"
import "flag"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	/*
		Command Args
		1. -c : No. of time repeat (0.....2^32) , -1 for infinit
		2. -d : delay between two hotkey in milliseconds
	*/

	repeatCount := flag.Int("c", 10, "Repeat Count [0 - 2^32] or [-1 for infinit]")
	hotkeyDelay := flag.Int("d", 2000, "Delay between two hotkeys in milliseconds")
	flag.Parse()

	if *hotkeyDelay <= 0 {
		*hotkeyDelay = 2000
	}

	cSignal := make(chan os.Signal, 1)
	signal.Notify(cSignal, os.Interrupt, syscall.SIGTERM)

	// initialize keyboard and check for possible errors
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("eurekaKeyboard"))
	if err != nil {
		return
	}
	// always do this after the initialization in order to guarantee that the device will be properly closed
	defer keyboard.Close()
	var ctx AppContext;
	ctx.keyboard = keyboard

	index := 0
	hotkeysCount := len(hotkeys)
	infinitCount := *repeatCount < 0

	for {
		select {
			case <-time.After(time.Duration((*hotkeyDelay)) * time.Millisecond):
				hotkey := hotkeys[index]
				hotkey.run(&ctx)
				index ++
				if index >= hotkeysCount {
					index = 0;
				}

				if !infinitCount {
					if *repeatCount --; *repeatCount <= 0 {
						fmt.Println("Repeat Completed");
						keyboard.Close()
						os.Exit(0)
					}
				}

			case <-cSignal:
				fmt.Println("Signal Recieved");
				keyboard.Close()
				os.Exit(0)
		}
	}
}
