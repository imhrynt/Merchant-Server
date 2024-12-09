package util

import "log"

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(err error) {
	if err != nil {
		log.Fatalln(map[string]string{"error": err.Error()})
		return
	}
}
