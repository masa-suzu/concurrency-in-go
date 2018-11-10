package chapter5

import (
	"fmt"
	"log"
	"os"
)

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v\n", key, message)
}

func Run() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		msg := "There was an expected issue; please report this as a bug."
		if _, ok := err.(IntermediateError); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)
	}
}
