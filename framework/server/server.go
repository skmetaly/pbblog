package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores the hostname and port number
/*
type Server struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}
*/

type Server struct {
	Hostname  string
	UseHTTP   bool
	UseHTTPS  bool
	HTTPPort  int
	HTTPSPort int
	CertFile  string
	KeyFile   string
}

func NewServer() Server {
	server := Server{}
	server.Hostname = ""
	server.HTTPPort = 9999

	return server
}

func StartServer(s Server, router http.Handler) {
	startHTTP(s, router)
}

// startHTTP starts the HTTP listener
func startHTTP(s Server, router http.Handler) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(s))

	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), router))
}

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}
