package main

import (
	"github.com/projects/vue-go-1st/api/models/passwords"
	"log"
	"os"
)

func main() {
	password := os.Args[1]
	log.Println(password)
	hash, err := passwords.Encrypt(password)
	if err != nil {
		panic(err)
	}
	log.Println(hash)

}
