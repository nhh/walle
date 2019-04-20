package main

import (
	"fmt"
	"github.com/nhh/walle/config"
	"github.com/nhh/walle/handler"
	"log"
	"net/http"
)


// Todo (dont forget)
// 1. Basic Configuration & Request Mapping (Toml, yaml, json, env)
// 2. SSL & Letsencrypt
// 3. Middlewares
// 4. Advanced Features (Caching, Client config, Server config)

func main() {

	defer fmt.Println("Server shutdown")

	for _, conf := range config.Parse() {
		handler.Build(conf)
	}

	// We are actually ignoring the fact that our rproxy could start with multiple ports and so on.

	log.Fatal(http.ListenAndServe(":8080", nil))

}
