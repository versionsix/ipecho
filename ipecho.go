package main

import (
    "fmt"
    "strings"
    "net/http"
    "os"
    "strconv"
)

const DEFAULT_PORT = 18888

func handler(w http.ResponseWriter, req *http.Request) {
    addr := strings.Split(req.RemoteAddr, ":")
    if len(addr) > 0 {
        fmt.Fprintf(w, addr[0])
    } else {
        fmt.Fprintf(w, "unknown")
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
    fmt.Printf("ipecho listen on port: %d\n", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
