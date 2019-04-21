package server

import (
    "errors"
    "fmt"
    "github.com/nhh/walle/handler"
    "net/http"
    "net/url"
)

func ParseHandler(location Location) (http.HandlerFunc, error) {
	switch location.Type{
		case "proxy": {
			fmt.Println("Mounting " + location.From + " To " + location.To)
			target, error := url.Parse(location.To)
			if (error != nil) {
			    return nil, errors.New("Cannot parse target url in location")
            }
			return handler.NewProxyHandler(*target), nil
		}
		case "file": {
			// This case would be a static file server
			// func(writer http.ResponseWriter, request *http.Request) {}
		    return nil, errors.New("Not implemted")
		}
		case "tcp": {
            return nil, errors.New("Not implemted")
		}
		default: {
            return nil, errors.New("Not implemted")
        }
	}
}
