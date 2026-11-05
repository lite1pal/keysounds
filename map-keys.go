package main

func EventToKey(key int) string {
	var textKey string
	switch key {
	case 16:
		textKey = "KeyQ"
	case 17:
		textKey = "KeyW"
	case 18:
		textKey = "KeyE"
	case 19:
		textKey = "KeyR"
	case 20:
		textKey = "KeyT"
	case 21:
		textKey = "KeyY"
	case 22:
		textKey = "KeyU"
	case 23:
		textKey = "KeyI"
	case 24:
		textKey = "KeyO"
	case 25:
		textKey = "KeyP"
	case 26:
		textKey = "BracketLeft"
	case 27:
		textKey = "BracketRight"
	case 14:
		textKey = "Backspace"
	case 30:
		textKey = "KeyA"
	case 31:
		textKey = "KeyS"
	case 32:
		textKey = "KeyD"
	case 33:
		textKey = "KeyF"
	case 34:
		textKey = "KeyG"
	case 35:
		textKey = "KeyH"
	case 36:
		textKey = "KeyJ"
	case 37:
		textKey = "KeyK"
	case 38:
		textKey = "KeyL"
	case 39:
		textKey = "Semicolon"
	case 40:
		textKey = "Quote"
	case 43:
		textKey = "Backslash"
	case 44:
		textKey = "KeyZ"
	case 45:
		textKey = "KeyX"
	case 46:
		textKey = "KeyC"
	case 47:
		textKey = "KeyV"
	case 48:
		textKey = "KeyB"
	case 49:
		textKey = "KeyN"
	case 50:
		textKey = "KeyM"
	case 51:
		textKey = "Comma"
	case 52:
		textKey = "Period"
	case 53:
		textKey = "Slash"
	case 57:
		textKey = "Space"
	case 15:
		textKey = "Tab"
	case 58:
		textKey = "CapsLock"
	case 1:
		textKey = "Escape"
	case 28:
		textKey = "Enter"
	case 42:
		textKey = "ShiftLeft"
	case 54:
		textKey = "ShiftRight"
	case 2:
		textKey = "Digit1"
	case 3:
		textKey = "Digit2"
	case 4:
		textKey = "Digit3"
	case 5:
		textKey = "Digit4"
	case 6:
		textKey = "Digit5"
	case 7:
		textKey = "Digit6"
	case 8:
		textKey = "Digit7"
	case 9:
		textKey = "Digit8"
	case 10:
		textKey = "Digit9"
	case 11:
		textKey = "Digit0"
	case 13:
		textKey = "Equal"
	case 12:
		textKey = "Minus"
	}

	return textKey
}
