package main

import (
    "github.com/nhh/walle/config"
    "github.com/nhh/walle/handler"
    "log"
    "net/http"
    "strconv"
)

func main() {

	for _, conf := range config.Parse() {

		// Todo move the "server generation" in its own package/builder
        handler.Build(conf)
        configuration := conf

		go func() {
            mux := http.NewServeMux()
            mux.HandleFunc("/.well-known/acme-challenge/",  handler.HandleLetsEncryptAcme)
            // Todo add the custom handler here

            server := http.Server{
                Addr: ":" + strconv.Itoa(configuration.Port),
                Handler: mux,
            }

            log.Fatal(server.ListenAndServe())
        }()
	}

	// this is the management server, may hold the references to the other server?
    http.ListenAndServe(":1995", nil)
}
