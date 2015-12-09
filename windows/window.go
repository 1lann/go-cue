package windows

// #include <windows.h>
import "C"

func GetForegroundWindowTitle() string {
	title := make([]C.CHAR, 2048)
	activeWindow := C.GetForegroundWindow()
	C.GetWindowText(activeWindow, &title[0], 2048)
	titleString := ""

	for _, char := range title {
		if byte(char) == 0 {
			break
		}

		titleString = titleString + string(rune(char))
	}

	return titleString
}
