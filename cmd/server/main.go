package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"os"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	if len(os.Getenv("HTTP2")) != 0 {
		H2CServerPrior()
	} else {
		H21ServerPrior()
	}
}

func H2CServerPrior() {
	server := http2.Server{}

	l, err := net.Listen("tcp", "0.0.0.0:1010")
	checkErr(err, "while listening")

	fmt.Printf("HTTP2 Listening [0.0.0.0:1010]...\n")
	for {
		conn, err := l.Accept()
		checkErr(err, "during accept")

		go func() {
			server.ServeConn(conn, &http2.ServeConnOpts{
				Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				}),
			})
		}()
	}
}
func hello(w http.ResponseWriter, req *http.Request) {}
func H21ServerPrior() {
	fmt.Printf("HTTP1 Listening [0.0.0.0:1010]...\n")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":1010", nil)
}
