package operations

import (
	"testing"
)

func TestMessageOK(t *testing.T) {

	var (
		messageok = [][]string{
			{"esto", "es", "un", "mensaje", "muy", "secreto"},
			{"esto", "es", "un", "mensaje", "muy", "secreto"},
			{"esto", "es", "un", "mensaje", "muy", "secreto"},
		}
	)

	answer := GetMessage(messageok...)

	if answer != "esto es un mensaje muy secreto" {
		t.Errorf("No fue posible obtener mensaje")
	}
}

func TestMessageWithLackOK(t *testing.T) {

	var (
		messageok = [][]string{
			{"esto", "", "un", "mensaje", "muy", "secreto"},
			{"", "es", "un", "", "muy", ""},
			{"esto", "", "un", "", "muy", "secreto"},
		}
	)

	answer := GetMessage(messageok...)

	if answer != "esto es un mensaje muy secreto" {
		t.Errorf("It was not possible to retrieve message")
	}
}

func TestMessageInvalid(t *testing.T) {

	var (
		messageok = [][]string{
			{"", "", "un", "mensaje", "muy", "secreto"},
			{"", "es", "un", "", "muy", ""},
			{"", "", "un", "", "muy", "secreto"},
		}
	)

	answer := GetMessage(messageok...)

	if answer != "" {
		t.Errorf("Error: Invalid message")
	}
}

func TestMessageInvalidLengths(t *testing.T) {

	var (
		messageok = [][]string{
			{"", "", "un", "mensaje", "muy", "secreto", "", ""},
			{"", "es", "un", "", "muy", ""},
			{"", "", "un", "", "muy", "secreto"},
		}
	)

	answer := GetMessage(messageok...)

	if answer != "" {
		t.Errorf("Error: the message information from every satellite should be the same length")
	}
}
