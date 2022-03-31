package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"
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
		fmt.Println("HTTP2 bench start")
		HttpClientExample()
	} else {
		fmt.Println("HTTP1 bench start")
		Http1ClientExample()
	}
}

const url = "http://127.0.0.1:1010"

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

func Http1ClientExample() {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
	}
	var cnt int64
	for i := 0; i < 16; i++ {
		go func() {
			for {
				_, err := client.Get(url)
				checkErr(err, "during get")
				atomic.AddInt64(&cnt, 1)
			}
		}()
	}
	for {
		prev := atomic.LoadInt64(&cnt)
		time.Sleep(5 * time.Second)
		now := atomic.LoadInt64(&cnt)
		fmt.Printf("avgQps %d\n", (now-prev)/5)
	}
}

func HttpClientExample() {
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	var cnt int64
	for i := 0; i < 16; i++ {
		go func() {
			for {
				_, err := client.Get(url)
				checkErr(err, "during get")
				atomic.AddInt64(&cnt, 1)
			}
		}()
	}
	for {
		prev := atomic.LoadInt64(&cnt)
		time.Sleep(5 * time.Second)
		now := atomic.LoadInt64(&cnt)
		fmt.Printf("avgQps %d\n", (now-prev)/5)
	}
}
