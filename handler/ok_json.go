package handler

import (
    "encoding/json"
    "net/http"
)

type ok struct {
    Status string // lul again :D
}

func HandleOkJson(w http.ResponseWriter, request *http.Request) {
    status := ok{"200 OK"}

    // I reeeally like json handling
    response, err := json.Marshal(status)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(response)

}
