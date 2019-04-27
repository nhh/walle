package server

import (
    "errors"
    "github.com/nhh/walle/handler"
    "log"
    "net/http"
    "net/url"
    "time"
)

func ParseHandler(location Location) (http.HandlerFunc, error) {
	switch location.Type{
		case "proxy": {
			log.Printf("Mounting %s to %s", location.From, location.To)
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

			return handler.NewProxyHandler(*target, &client), nil
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
