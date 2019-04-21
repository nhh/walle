package handler

import (
    "net/http"
)

func HandleLetsEncryptAcme(w http.ResponseWriter, request *http.Request) {
    w.Header().Add("X-Server-Used", "W.A.L.L.E")
    w.Header().Add("Content-Type", "text/plain")
    w.WriteHeader(200)
}
