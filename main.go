package main

import (
    "fmt"
    "github.com/nhh/walle/config"
    "net/http"
    "strconv"
)

func main() {
    for _, server := range config.ParseServers() {
        fmt.Println("Starting server at: " + strconv.Itoa(server.Port))
        server.StartAsync()
    }
	// this is the management server, may hold the references to the other server?
    http.ListenAndServe(":1995", nil)
}
