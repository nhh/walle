package server

import (
    "fmt"
    "net/http"
    "net/url"
)

func ParseHandler(vhost Location) http.HandlerFunc {

	endpoint, error := url.Parse(vhost.From)

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
			return nil
		}
		case "file": {
			// This case would be a static file server
			// func(writer http.ResponseWriter, request *http.Request) {}
		    return nil
		}
		case "tcp": {
            return nil
		}
		default: return nil//func(writer http.ResponseWriter, request *http.Request) {}
	}

}
