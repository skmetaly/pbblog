package main

import (
	"fmt"
	"github.com/skmetaly/pbblog/framework/server"
)

func init() {

}

func main() {
	fmt.Println("Starting http")
	// Start the HTTP server
	server.StartHTTP()
}
