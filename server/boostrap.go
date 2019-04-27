package server

import (
    "log"
    "strconv"
)

func Bootstrap(serverList []WalleServer) {
    for _, s := range serverList {
        go obtainCertOrRunServer(s)
    }
}

func obtainCertOrRunServer(server WalleServer) {
    log.Printf("%s Starting at: %s", server.Name, strconv.Itoa(server.Port))

    _, issueError := IssueOrSkipCertificate(server)

    if issueError != nil {
        log.Fatal(issueError)
        return
    }

    server.StartAsync()
}

