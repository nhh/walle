package main

import (
	"fmt"
	"github.com/nhh/walle/config"
	"github.com/nhh/walle/handler"
	"log"
	"net/http"
)

func main() {

	defer fmt.Println("Server shutdown")

	for _, conf := range config.Parse() {
		handler.Build(conf)
	}

	http.HandleFunc("/.well-known/acme-challenge/", handler.HandleLetsEncryptAcme)
	// We are actually ignoring the fact that our rproxy could start with multiple ports and so on.

	log.Fatal(http.ListenAndServe(":8080", nil))

}
