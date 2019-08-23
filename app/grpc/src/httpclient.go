package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://grpc-client:8080/hello")
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
}
