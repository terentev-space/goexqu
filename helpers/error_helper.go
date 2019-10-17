package helpers

import "log"

func CheckError(err error, message string) {
	if err != nil {
		log.Fatalf("%s\n%s", message, err)
	}
}
