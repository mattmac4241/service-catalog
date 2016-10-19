package service

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"

	"github.com/unrolled/render"
)

var (
    formatter = render.New(render.Options{
        IndentJSON: true,
    })
)

type testRepo struct {
    redis           map[string]string
    services        []Service
    snapShots       []SnapShot
}

func (t *testRepo) registerService(service Service) error {
    t.services = append(t.services, service)
    return nil
}

func (t *testRepo) addSnapshot(snapshot SnapShot) error {
    t.snapShots = append(t.snapShots, snapshot)
    return nil
}

func (t *testRepo) GetAllKeys() ([]string, error) {
    var keys []string
    for k := range t.redis {
        keys = append(keys, k)
    }

    return keys, nil
}

func (t *testRepo) RedisGetValue(key string) (string, error) {
    return t.redis[key], nil
}

func (t *testRepo) getServices() []Service {
    return t.services
}

func TestPostRegisterHandlerInvalidJSON(t *testing.T) {
    repo := &testRepo{}
    client := &http.Client{}

    server := httptest.NewServer(http.HandlerFunc(postRegisterServiceHandler(formatter, repo)))
    defer server.Close()

    body := []byte("this is not valid json")
    req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
    if err != nil {
        t.Errorf("Error in creating POST request for createMatchHandler: %v", err)
    }
    res, err := client.Do(req)
    if err != nil {
        t.Errorf("Error in POST to createMatchHandler: %v", err)
    }
    defer res.Body.Close()
    if res.StatusCode != http.StatusBadRequest {
        t.Error("Sending invalid JSON should result in a bad request from server.")
    }
}

func TestPostRegisterHandlerNotService(t *testing.T) {
    repo := &testRepo{}
    client := &http.Client{}

    server := httptest.NewServer(http.HandlerFunc(postRegisterServiceHandler(formatter, repo)))
    defer server.Close()

    body := []byte("{\"test\":\"Not user.\"}")
    req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
    if err != nil {
        t.Errorf("Error in creating second POST request for invalid data on create match: %v", err)
    }
    req.Header.Add("Content-Type", "application/json")
    res, _ := client.Do(req)
    defer res.Body.Close()
    if res.StatusCode != http.StatusBadRequest {
        t.Error("Sending valid JSON but with incorrect or missing fields should result in a bad request and didn't.")
    }
}

func TestPostRegisterHandlerService(t *testing.T) {
    repo := &testRepo{}
    client := &http.Client{}

    server := httptest.NewServer(http.HandlerFunc(postRegisterServiceHandler(formatter, repo)))
    defer server.Close()

    body := []byte("{\"name\":\"test\",\"url\":\"localhost:300\"}")
    req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
    if err != nil {
        t.Errorf("Error in creating second POST request for invalid data on create match: %v", err)
    }
    req.Header.Add("Content-Type", "application/json")
    res, _ := client.Do(req)
    defer res.Body.Close()
    if res.StatusCode == http.StatusBadRequest {
        t.Error("Sent valid service something went wrong")
    }

    if len(repo.services) != 1 {
        t.Errorf("Expected services length 1 instead got %d", repo.services)
    }

    service := repo.services[0]
    if service.Name != "test" && service.URL != "localhost:3000" {
        t.Error("Service not correct")
    }
}

func TestGetServiceHandler(t *testing.T) {
    repo := &testRepo{}
    client := &http.Client{}

    service1 := Service{Name: "test1", URL: "localhost:3000"}
    service2 := Service{Name: "test2", URL: "localhost:3001"}
    repo.registerService(service1)
    repo.registerService(service2)

    server := httptest.NewServer(http.HandlerFunc(getServicesHandler(formatter, repo)))
    defer server.Close()

    req, err := http.NewRequest("GET", server.URL, nil)
    if err != nil {
        t.Errorf("Error in creating second POST request for invalid data on create match: %v", err)
    }
    req.Header.Add("Content-Type", "application/json")
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected statusOK instead got %d", resp.StatusCode)
    }


    var services []Service
    payload, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        t.Error("Failed to read response from server", err)
    }
    err = json.Unmarshal(payload, &services)
    if err != nil {
        t.Error(err)
    }
    if len(services) != 2 {
        t.Errorf("Expected two services got %d", len(services))
    }
}
