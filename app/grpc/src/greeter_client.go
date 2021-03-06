package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/http2"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var conn *grpc.ClientConn
var address string

func init() {
	var err error
	flag.StringVar(&address, "address", "localhost:50051", "address: localhost:50051")
	flag.Parse()
	if strings.TrimSpace(address) == "" {
		panic("address is nil")
	}

	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	c := pb.NewGreeterClient(conn)
	name := time.Now().Unix()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: strconv.FormatInt(name, 10)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	responseBody := fmt.Sprintf("Greeting: %s", resp.Message)
	w.Write([]byte(responseBody))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello h2c")
	})
	mux.HandleFunc("/hello", helloHandler)
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}
	http2.ConfigureServer(s, s2)
	l, _ := net.Listen("tcp", ":8080")
	defer l.Close()
	for {
		rwc, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: s})

	}
}
