package service

import (
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/unrolled/render"
)

func postRegisterServiceHandler(formatter *render.Render, repo repository) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        var service Service
        payload, _ := ioutil.ReadAll(req.Body)
        err := json.Unmarshal(payload, &service)
        if err != nil || (service == Service{}) {
            formatter.JSON(w, http.StatusBadRequest, "Failed to parse service.")
            return
        }
        err = repo.registerService(service)
        if err != nil {
            formatter.JSON(w, http.StatusInternalServerError, "Failed to register service.")
            return
        }
        formatter.JSON(w, http.StatusCreated, "Service succesfully registered.")
    }
}


func getServicesHandler(formatter *render.Render, repo repository) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        services := repo.getServices()
        formatter.JSON(w, http.StatusOK, services)
    }
}
