package main

import (
	"log"
	"fmt"
	"net"
	"os"
	"strconv"
)

const DEFAULT_PORT = 18889

func logError(e error) {
	if e != nil {
		log.Print(e)
	}
}

func noError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func serve(c *net.TCPConn) {
	ret := c.RemoteAddr().(*net.TCPAddr).IP.String()
	_, err := c.Write([]byte(ret))
	logError(err)
	err = c.Close()
	logError(err)
}

func main() {
	var port int = DEFAULT_PORT
	var err error

	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	face := fmt.Sprintf(":%d", port)
	addr, err := net.ResolveTCPAddr("tcp4", face)
	noError(err)

	lis, err := net.ListenTCP("tcp4", addr)
	noError(err)

	log.Printf("rawipecho listening on port: %d", port)

	for {
		con, err := lis.AcceptTCP()
		noError(err)

		go serve(con)
	}
}
