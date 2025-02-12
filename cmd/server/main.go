package main

import (
	"context"
	"fmt"
	"log"

	"grpc_client_streaming_sample/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------------------------server")
}
