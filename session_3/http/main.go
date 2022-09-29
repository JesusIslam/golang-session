package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// using default http server and default muxer
	go runDefaultServer()

	// using custom http server
	go runServer()

	// request data
	requestData()
}

func runDefaultServer() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Got request from default http server")
		_, err := io.WriteString(res, "OK!")
		if err != nil {
			panic(err)
		}
	})

	// this will block if we don't run it in a goroutine
	http.ListenAndServe(":8080", nil)
}

func runServer() {
	var srv http.Server

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Got request from new http server")
		_, err := io.WriteString(res, "OK2!")
		if err != nil {
			panic(err)
		}
	})

	srv.Addr = ":8181"
	srv.Handler = mux

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func requestData() {
	// using default client
	reqDefault, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(reqDefault)
	if err != nil {
		panic(err)
	}
	// response body MUST be closed or it would not be picked by GC
	defer res.Body.Close()
	fmt.Printf("%s\n", res.Body)

	// using custom client
	var cli http.Client

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8181/", nil)
	if err != nil {
		panic(err)
	}
	res2, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	defer res2.Body.Close()
	fmt.Printf("%s\n", res2.Body)
}
