package handler

import (
    "io"
    "net/http"
    "net/url"
    "time"
)

func NewProxyHandler(target url.URL) http.HandlerFunc {
    return func(w http.ResponseWriter, request *http.Request) {
        tr := &http.Transport {
            MaxIdleConns:       10,
            IdleConnTimeout:    30 * time.Second,
            DisableCompression: true,
        }

        client := &http.Client{Transport: tr}

        req := http.Request{
            Method: request.Method,
            RemoteAddr: request.RemoteAddr,
            URL: &target,
            ContentLength: request.ContentLength,
            GetBody: request.GetBody,
        }

        response, error := client.Do(&req)

        // Todo Move this into an own middleware handler
        w.Header().Add("X-Server-Used", "W.A.L.L.E")

        if(error != nil) {
            w.WriteHeader(500)
            return
        }

        // Very limited, we actually need to set the complete response header here.
        //w.Header()
        w.WriteHeader(response.StatusCode)
        io.Copy(w, response.Body)
    }
}
