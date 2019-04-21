package server

import (
    "context"
    "github.com/nhh/walle/handler"
    "log"
    "net/http"
    "strconv"
    "time"
)

type Runnable interface {
    Start()
    Stats() // Whatever
}

type WalleServer struct {
    Locations [] Location //lul, that was easy
    Port int
    Ssl bool // Abstract this into a struct
    httpServer http.Server
    running bool
    // Various other configuration types
}

func (walleServer WalleServer) Start() {
    if walleServer.running {
        return
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/.well-known/acme-challenge/",  handler.HandleLetsEncryptAcme)
    // Todo add the custom handler here

    // could we possibly use our own Server instead of these?
    walleServer.httpServer = http.Server{
        Addr: ":" + strconv.Itoa(walleServer.Port),
        Handler: mux,
    }

    // Mindfuck
    log.Fatal(walleServer.httpServer.ListenAndServe())
}

func (walleServer WalleServer) StartAsync() {
    if walleServer.running {
        return
    }
    go walleServer.Start()
}

func (walleServer WalleServer) Stop() {
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Thx stack overflow
    walleServer.httpServer.Shutdown(ctx)
}

func (walleServer WalleServer) Stats() {

}

func parseLocations() {

}
