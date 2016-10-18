package service

import (
    "fmt"
    "net/http"
)

type pingClient interface {
    Ping()
}

type PingWebClient struct {
    RootURL string
}

func (client PingWebClient) Ping(name string, repo RepoHandler) {
    httpclient := &http.Client{}
    fmt.Println("CALLED")
    fmt.Println(client.RootURL)
    url := fmt.Sprintf("http://%s/%s", client.RootURL, "ping")
    fmt.Println(url)

    _, err := httpclient.Get(url)
    statusCode := 200
    snapShot := SnapShot{Name: name, URL: client.RootURL, ResponseStatus: statusCode, Active: true}
    if err != nil {
        snapShot.Active = false
        snapShot.ResponseStatus = 400
    }
    repo.addSnapshot(snapShot)
}
