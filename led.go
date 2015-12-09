package cue

// #include "CUESDK.h"
import "C"

import (
	"image/color"
	"unicode"
	"unsafe"
)

type LedPosition struct {
	Top    float64
	Left   float64
	Height float64
	Width  float64
}

type LedId int

type Led struct {
	Id       LedId
	KeyName  rune
	Position LedPosition
}

type LedColor struct {
	Led *Led
	R   int
	G   int
	B   int
}

var numLeds int
var allLeds map[LedId]Led

type internalPositions C.CorsairLedPositions
type internalPosition C.CorsairLedPosition
type internalLedColor C.CorsairLedColor

func (l Led) SetColor(col color.Color) error {
	r, g, b, _ := col.RGBA()

	internalColor := internalLedColor{
		ledId: C.CorsairLedId(l.Id),
		r:     C.int(r / (1 << 8)),
		g:     C.int(g / (1 << 8)),
		b:     C.int(b / (1 << 8)),
	}

	if C.CorsairSetLedsColors(C.int(1),
		(*C.CorsairLedColor)(unsafe.Pointer(&internalColor))) == 1 {
		return nil
	} else {
		return getLastError()
	}
}

func (l Led) GetMultiLedColor(col color.Color) LedColor {
	r, g, b, _ := col.RGBA()

	return LedColor{
		Led: &l,
		R:   int(r / (1 << 8)),
		G:   int(g / (1 << 8)),
		B:   int(b / (1 << 8)),
	}
}

func SetMultiLedColors(ledColors []LedColor) error {
	internalLedColors := make([]internalLedColor, len(ledColors))
	for i, ledColor := range ledColors {
		internalLedColors[i] = internalLedColor{
			ledId: C.CorsairLedId(ledColor.Led.Id),
			r:     C.int(ledColor.R),
			g:     C.int(ledColor.G),
			b:     C.int(ledColor.B),
		}
	}

	if C.CorsairSetLedsColors(C.int(len(internalLedColors)),
		(*C.CorsairLedColor)(unsafe.Pointer(&internalLedColors[0]))) == 1 {
		return nil
	} else {
		return getLastError()
	}
}

func GetLedById(id LedId) (Led, bool) {
	if numLeds == 0 {
		panic("cue: no leds found (did you call cue.InitializeLeds()?)")
	}

	if led, found := allLeds[id]; found {
		return led, true
	} else {
		return Led{}, false
	}
}

func GetLedByKey(key rune) (Led, bool) {
	if numLeds == 0 {
		panic("cue: no leds found (did you call cue.InitializeLeds()?)")
	}

	key = unicode.ToUpper(key)

	for _, led := range allLeds {
		if led.KeyName == key {
			return led, true
		}
	}

	return Led{}, false
}

func GetLedAtPosition(top float64, left float64) (Led, bool) {
	if numLeds == 0 {
		panic("cue: no leds found (did you call cue.InitializeLeds()?)")
	}

	for _, led := range allLeds {
		if (top-led.Position.Top) < led.Position.Height &&
			(left-led.Position.Left) < led.Position.Width {
			return led, true
		}
	}

	return Led{}, false
}

func GetAllLeds() []Led {
	if numLeds == 0 {
		panic("cue: no leds found (did you call cue.InitializeLeds()?)")
	}

	allLedsSlice := make([]Led, numLeds)

	i := 0
	for _, led := range allLeds {
		allLedsSlice[i] = led
		i++
	}

	return allLedsSlice
}

// Returns the number of leds found, and any errors.
// Returns 0 leds and no errors (nil) if there is no device connected which
// has leds.
// You must call Initialize() first before you may call InitializeLeds().
func InitializeLeds() (int, error) {
	allLeds = make(map[LedId]Led)

	positionsPtr := unsafe.Pointer(C.CorsairGetLedPositions())
	if positionsPtr == nil {
		return 0, getLastError()
	}

	ledPositions := *(*internalPositions)(positionsPtr)
	numLeds = int(ledPositions.numberOfLed)

	if numLeds == 0 {
		return 0, nil
	}

	internalLeds := (*[1 << 24]internalPosition)(unsafe.Pointer(
		ledPositions.pLedPosition))[:numLeds:numLeds]

	for _, led := range internalLeds {
		allLeds[LedId(led.ledId)] = Led{
			Id: LedId(led.ledId),
			Position: LedPosition{
				Top:    float64(led.top),
				Left:   float64(led.left),
				Height: float64(led.height),
				Width:  float64(led.width),
			},
		}
	}

	for char := 'A'; char <= 'Z'; char++ {
		ledId := C.CorsairGetLedIdForKeyName(C.char(char))
		if LedId(ledId) == CLI_Invalid {
			return 0, getLastError()
		}

		currentLed := allLeds[LedId(ledId)]
		currentLed.KeyName = char
		allLeds[LedId(ledId)] = currentLed
	}

	return numLeds, nil
}
