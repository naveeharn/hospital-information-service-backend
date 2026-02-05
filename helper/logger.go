package helper

import (
	"log"
)

func LoggerErrorPath(pc uintptr, filename string, line int, ok bool) {
	log.Printf("[ ERROR ] %s:%d", filename, line)
}
