package commons

import "fmt"

func IrrecoverableFailure(message string, err error) {
	panicMessage := fmt.Sprintf("%s: cause (%s)", message, err.Error())
	panic(panicMessage)
}
