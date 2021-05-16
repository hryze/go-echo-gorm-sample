package main

import (
	"log"

	"github.com/paypay3/go-echo-gorm-sample/infrastructure/router"
)

func main() {
	if err := router.Run(); err != nil {
		log.Fatalf("%+v", err)
	}
}
