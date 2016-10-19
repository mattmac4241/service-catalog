package service

import (
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/unrolled/render"
)

func postRegisterHandler(formatter *render.Render, repo repository) http.HandlerFunc {
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
        services := []Service{}
        keys, _ := repo.GetAllKeys()
        for _, key := range keys {
            val, _ := repo.RedisGetValue(key)
            service := Service{Name: key, URL: val}
            services = append(services, service)
        }
        formatter.JSON(w, http.StatusOK, services)
    }
}
