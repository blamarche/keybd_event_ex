package keybd_event_ex_test

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func ExampleNewKeyBonding() {
	kb, err := keybd_event_ex.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event_ex.VK_A, keybd_event_ex.VK_B)

	// Set shift to be pressed
	kb.HasSHIFT(true)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

	// Here, the program will generate "ABAB" as if they were pressed on the keyboard.

	// Output:
}
