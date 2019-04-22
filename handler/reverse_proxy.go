package handler

import (
    "io"
    "net/http"
    "net/url"
)

func NewProxyHandler(target url.URL, client *http.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, request *http.Request) {
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
        defer response.Body.Close()

        if error != nil {
            w.WriteHeader(500)
            return
        }

        // Very limited, we actually need to set the complete response header here.
        //w.Header()
        w.WriteHeader(response.StatusCode)
        io.Copy(w, response.Body)
    }
}
