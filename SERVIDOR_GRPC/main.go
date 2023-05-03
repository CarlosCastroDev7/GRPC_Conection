package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/carlos/grpc/api"
)

func main() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		err := api.Api_GRPC("127.0.0.1:50001")
		if err != nil {
			fmt.Println(err)
			panic(1)
		}
	}()

	<-ch
	fmt.Printf("\nsignal caught. shutting down...")

}
