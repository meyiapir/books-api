package main

import (
	"booksApi"
	"log"
)

func main() {
	srv := new(booksApi.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("error occured while running http server: ", err.Error())
	}
}
