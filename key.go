package cue

// #include <windows.h>
import "C"

import (
	"sync"
	"time"
)

const (
	Down = 1
	Up   = 2
)

type Key struct {
	WindowsKeyId int
	HasLed       bool
	Led          Led
	State        int
}

var keyMonitorMutex *sync.Mutex = &sync.Mutex{}
var KeysDownChanges chan []Key
var KeysUpChanges chan []Key
var KeyboardState [256]Key

var shouldMonitorKeyChanges bool = false

// TODO: Uncomment for release

func StartKeyMonitor() {
	for c := 1; c < 255; c++ {
		// led, found := getLedFromWindowsKeyId(c)

		KeyboardState[c] = Key{
			WindowsKeyId: c,
			// HasLed:       found,
			// Led:          led,
			State: Up,
		}
	}

	KeysDownChanges = make(chan []Key)
	KeysUpChanges = make(chan []Key)

	shouldMonitorKeyChanges = true
	go monitorKeyEvents()
}

func monitorKeyEvents() {
	keyMonitorMutex.Lock()

	defer keyMonitorMutex.Unlock()

	for {
		var keysUp []Key
		var keysDown []Key

		for c := 1; c < 255; c++ {
			keyState := int16(C.GetAsyncKeyState(C.int(c)))

			if keyState&-32768 == -32768 {
				if KeyboardState[c].State == Up {
					key := KeyboardState[c]
					key.State = Down
					keysDown = append(keysDown, key)
				}
			} else {
				if KeyboardState[c].State == Down {
					key := KeyboardState[c]
					key.State = Up
					keysUp = append(keysUp, key)
				}
			}
		}

		if !shouldMonitorKeyChanges {
			close(KeysUpChanges)
			close(KeysDownChanges)
			return
		}

		if len(keysDown) > 0 {
			select {
			case KeysDownChanges <- keysDown:
				for _, key := range keysDown {
					KeyboardState[key.WindowsKeyId] = key
				}
			default:
			}
		}

		if len(keysUp) > 0 {
			select {
			case KeysUpChanges <- keysUp:
				for _, key := range keysUp {
					KeyboardState[key.WindowsKeyId] = key
				}
			default:
			}
		}

		time.Sleep(time.Millisecond * 30)
	}
}

func StopKeyMonitor() {
	shouldMonitorKeyChanges = false

	keyMonitorMutex.Lock()
	keyMonitorMutex.Unlock()
}
