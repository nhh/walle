package handler

import (
    "net/http"
)


// This is the base letsencrypt handler which is used to the certificate request validation
// Not sure what the server should answer on /.well-known/acme-challenge/ yet.
// Maybe we can find a spec or rfc.

func HandleLetsEncryptAcme(w http.ResponseWriter, request *http.Request) {
    w.Header().Add("X-Server-Used", "W.A.L.L.E")
    w.Header().Add("Content-Type", "text/plain")
    w.Header().Add("X-Lets-Encrypt", "Renew")
    w.WriteHeader(200)
}
