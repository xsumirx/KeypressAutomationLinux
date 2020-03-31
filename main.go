package main

import "gopkg.in/bendahl/uinput.v1"
import "time"
import "flag"


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

	// initialize keyboard and check for possible errors
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("eurekaKeyboard"))
	if err != nil {
		return
	}
	// always do this after the initialization in order to guarantee that the device will be properly closed
	defer keyboard.Close()
	var ctx AppContext;
	ctx.keyboard = keyboard

	for *repeatCount != 0  {
		for _,hotkey := range hotkeys {
			time.Sleep(time.Duration((*hotkeyDelay)) * time.Millisecond)
			hotkey.run(&ctx)
		}

		if *repeatCount > 0 {
			*repeatCount --
		}

	}
	
}
