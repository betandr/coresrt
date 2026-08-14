package main

import (
	"flag"
	"log"

	"coresrt/receiver"
)

func main() {
	port := flag.Int("port", 9999, "UDP port to listen on")
	addr := flag.String("addr", "0.0.0.0", "IP address to bind to")
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Printf("starting SRT receiver on %s:%d", *addr, *port)

	opts := receiver.Options{
		Debug: *debug,
	}

	if opts.Debug {
		log.Println("debug mode on: will dump incoming packets")
	}

	receiver.Start(*port, *addr, opts)
}
