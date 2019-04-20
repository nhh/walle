package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Handle(w http.ResponseWriter, request *http.Request) {

	tr := &http.Transport {
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	errorChannel, responseChannel := askExampleCom(client, request)

	select {

	case e := <- errorChannel:
		fmt.Println("ERROR received", e)
		w.WriteHeader(500)
	case response := <- responseChannel:
		fmt.Println("SUCCESS received", response)
		// Todo real merging of response headers
		// Todo feature: Appending various headers per config
		w.Header().Add("X-Server-Used", "W.A.L.L.E")
		io.Copy(w, response.Body)
	}

}

func askExampleCom(client *http.Client, request *http.Request) (chan error, chan *http.Response){

	errorChannel := make(chan error)
	responseChannel := make(chan *http.Response)

	target, _ := url.Parse("https://github.com/mholt/caddy")

	go func() {
		req := http.Request{
			Method: request.Method,
			RemoteAddr: request.RemoteAddr,
			URL: target,
			ContentLength: request.ContentLength,
			GetBody: request.GetBody,
		}

		resp, err := client.Do(&req)

		if err != nil {
			errorChannel <- err
		}

		responseChannel <- resp
	}()

	return errorChannel, responseChannel

}