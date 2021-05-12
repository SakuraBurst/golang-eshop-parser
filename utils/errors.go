package utils

import "log"

func CheckErrorAndShotDownIfItIs(err error, additionalMessage string) {
	if err != nil {
		log.Fatal(err, additionalMessage)
	}
}
