package utils

import "github.com/inancgumus/screen"

func ValidateOp(message string) (bool, byte) {
	op := message[0] - '0'
	invalidInput := true
	for i := 1; i <= 4; i++ {
		if op == byte(i) {
			invalidInput = false
			break
		}
	}

	if invalidInput || len(message) != 2 {
		return false, op
	}

	return true, op
}

func ResetScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}
