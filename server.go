package main

import (
	"flag"
	"fmt"
	"github.com/JingBh/crypto-learn/internal"
)

func main() {
	publishPtr := flag.Bool("publish", false, "Publish the server to the internet")
	flag.Parse()

	go func() {
		fmt.Println("Please visit http://localhost:8206 to view the web UI.")
	}()

	internal.StartServer(*publishPtr)
}
