package service

import (
    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a server.
func NewServer() *negroni.Negroni {
    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mux := mux.NewRouter()
    repo := &RepoHandler{}
    initRoutes(mux, formatter, repo)
    n.UseHandler(mux)

    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo repository) {
    mx.HandleFunc("/services/service", postRegisterServiceHandler(formatter, repo)).Methods("POST")
    mx.HandleFunc("/services/services", getServicesHandler(formatter, repo)).Methods("GET")
}
