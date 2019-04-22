package main

import (
    "fmt"
    "github.com/nhh/walle/config"
    server2 "github.com/nhh/walle/server"
    "net/http"
    "strconv"
)

func main() {

    // Setup letsencrypt



    for _, server := range config.ParseServers() {
        fmt.Println("Starting server at: " + strconv.Itoa(server.Port))
        server2.IssueCertificate(server)
        server.StartAsync()
    }
	// this is the management server, may hold the references to the other server?
    http.ListenAndServe(":1995", nil)
}
