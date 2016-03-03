package main

// -- takes routes
// -- initializes websockets
// -- walks the filepath
// -- serves files from util

import (
	"flag"
)

func main() {
	var (
		addr string = "localhost:3000"
	)

	flag.StringVar(&addr, "addr", "localhost:3000", "")
	flag.Parse()

	server := NewServer(addr)
	StartServer(server)
}
