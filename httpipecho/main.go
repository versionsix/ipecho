package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_PORT = 18888

func handler(w http.ResponseWriter, req *http.Request) {
	addr, e := net.ResolveTCPAddr("tcp", req.RemoteAddr)
	if e != nil {
		fmt.Fprintf(w, "error: %s", addr)
	} else {
		fmt.Fprintf(w, addr.IP.String())
	}
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
	fmt.Printf("httpipecho listen on port: %d\n", port)
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
