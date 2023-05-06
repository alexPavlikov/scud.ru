package algorithms

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
