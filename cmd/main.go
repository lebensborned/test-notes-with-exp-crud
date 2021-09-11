package main

import (
	"kode-task/server"
	"log"
	"time"
)

func main() {
	log.Println("Start service")
	srv := server.NewServer()

	srv.Storage.Add("value", time.Second*5)
	srv.Storage.Add("value2", time.Second*70)
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
