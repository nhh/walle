package handler

import (
	"fmt"
	"github.com/nhh/walle/config" // Now i got this idea...
	"net/http"
	"net/url"
)

func Build(vhost config.VirtualHost) {

	endpoint, error := url.Parse(vhost.Location["from"])

	if error != nil {
		panic("Cannot read endpoint! It must be a valid URI => https://tools.ietf.org/html/rfc3986")
	}

	switch endpoint.Scheme {
		case "https": {
			// This case would be a SSL Proxy
			fmt.Println("Mounting " + endpoint.Path)
			http.HandleFunc(endpoint.Path, func(writer http.ResponseWriter, request *http.Request) {
				fmt.Fprintln(writer, endpoint.Path)
			})
			return
		}
		case "file": {
			// This case would be a static file server
			// func(writer http.ResponseWriter, request *http.Request) {}
		}
		case "tcp": {

		}
		default: //func(writer http.ResponseWriter, request *http.Request) {}
	}

}
