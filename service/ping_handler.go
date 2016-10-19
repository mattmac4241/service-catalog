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
    url := fmt.Sprintf("http://%s/%s", client.RootURL, "ping")

    _, err := httpclient.Get(url)
    statusCode := 200
    snapShot := SnapShot{Name: name, URL: client.RootURL, ResponseStatus: statusCode, Active: true}
    if err != nil {
        snapShot.Active = false
        snapShot.ResponseStatus = 400
    }
    repo.addSnapshot(snapShot)
}
