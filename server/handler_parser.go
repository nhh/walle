package server

import (
    "errors"
    "fmt"
    "github.com/nhh/walle/handler"
    "net/http"
    "net/url"
    "time"
)

func ParseHandler(location Location) (http.HandlerFunc, error) {
	switch location.Type{
		case "proxy": {
			fmt.Println("Mounting " + location.From + " To " + location.To)
			target, parseError := url.Parse(location.To)

			if parseError != nil {
			    return nil, errors.New("Cannot parse target url in location")
            }

			tr := &http.Transport {
                MaxIdleConns:       10,
                IdleConnTimeout:    30 * time.Second,
                DisableCompression: true,
            }

            client := http.Client{Transport: tr}

			return handler.NewProxyHandler(*target, client), nil
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
