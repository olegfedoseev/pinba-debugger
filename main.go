package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/olegfedoseev/pinba"
)

var (
	in_addr  = flag.String("in", "0.0.0.0:30002", "incoming socket")
	raw      = flag.Bool("raw", false, "dump raw request object")
	rawBytes = flag.Bool("bytes", false, "dump raw request bytes in hex format")
)

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp4", *in_addr)
	if err != nil {
		log.Fatalf("Can't resolve address: '%v'", err)
	}

	sock, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatalf("Can't open UDP socket: '%v'", err)
	}

	log.Printf("Start listening on udp://%v\n", *in_addr)

	defer sock.Close()

	for {
		var buf = make([]byte, 65536)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf("Error on sock.ReadFrom, %v", err)
		}
		if rlen == 0 {
			continue
		}

		if *rawBytes {
			fmt.Printf("%#v\n", buf[0:rlen])
		}

		request, err := pinba.NewRequest(buf[0:rlen])
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		if *raw {
			fmt.Printf("%#v\n", request)
		}

		log.Printf("%s %s: %3.2f %s\n",
			request.ServerName,
			request.ScriptName,
			request.RequestTime,
			request.Tags.String(),
		)
		fmt.Printf("\t%s\n", request.Timers.String())
	}
}
