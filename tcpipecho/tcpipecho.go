package main

import "fmt"
import "net"
import "strings"
import "os"
import "strconv"

const DEFAULT_PORT = 18889

func serve(c *net.TCPConn) {
	fields := strings.Split(c.RemoteAddr().String(), ":")

	var ret string
	if len(fields) > 0 {
		ret = fields[0]
	} else {
		ret = "<unknown>"
	}

	_, err := c.Write([]byte(ret))
	if err != nil {
		fmt.Println(err)
	}
	err = c.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
    var port int = DEFAULT_PORT
    var err error
    
    if len(os.Args) > 1 {
        port, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Println(err)
            return
        }
    }

	face := fmt.Sprintf(":%d", port)
	addr, err := net.ResolveTCPAddr("tcp4", face)
	if err != nil {
		panic(err)
	}
	lis, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		panic(err)
	}

    fmt.Printf("rawipecho listening on port: %d\n", port)

	for {
		con, err := lis.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go serve(con)
	}
}
