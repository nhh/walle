package main

import (
    "github.com/nhh/walle/config"
    "github.com/nhh/walle/server"
    "log"
    "net/http"
)

func main() {

    walleServers := config.ParseServers()

    if len(walleServers) == 0 {
        log.Printf("Hello World", "")
    } else {
        server.Bootstrap(walleServers)
    }

    log.Printf("Management Server is startig at: %s", ":1995")

    // Todo connect each walle Server with the master server via channels
    http.ListenAndServe(":1995", nil)

}
