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
    Name string
    Locations [] Location //lul, that was easy
    Port int
    Ssl Ssl // Abstract this into a struct
    Domain string
    httpServer http.Server
    running bool
    // Various other configuration types
}

type Ssl struct {
    Email string
    AcceptTos bool
}

func (walleServer WalleServer) Start() {

    if walleServer.running {
        return
    }

    mux := http.NewServeMux()

    for _, location := range walleServer.Locations {
        handlerFunction, parseError := ParseHandler(location)
        if parseError != nil {
            log.Printf(parseError.Error())
            continue
        }
        mux.HandleFunc(location.From, handlerFunction)
    }

    // TODO Add loggin middleware and so on
    mux.HandleFunc("/.well-known/acme-challenge/",  handler.HandleLetsEncryptAcme)
    mux.HandleFunc("/hello-world",  handler.HandleOkJson)
    // Todo add the custom handler here

    // could we possibly use our own Server instead of these?
    walleServer.httpServer = http.Server{
        Addr: ":" + strconv.Itoa(walleServer.Port),
        Handler: mux,
    }

    if walleServer.Ssl.AcceptTos {
        log.Fatal(walleServer.httpServer.ListenAndServeTLS("./data/" + walleServer.Domain + "/cert.pem", "./data/" + walleServer.Domain + "/key.pem"))
    } else {
        log.Fatal(walleServer.httpServer.ListenAndServe())
    }

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
