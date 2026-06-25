package bca

import (
	"log"

	"github.com/emersion/go-message"
)

func Inspect(entity *message.Entity) {

	log.Println("========== EMAIL ==========")

	log.Println("Content-Type :", entity.Header.Get("Content-Type"))

	log.Println("Subject      :", entity.Header.Get("Subject"))

	log.Println("From         :", entity.Header.Get("From"))

	log.Println("===========================")
}
