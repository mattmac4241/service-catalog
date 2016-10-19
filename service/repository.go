package service

import (
    "github.com/garyburd/redigo/redis"
)

type repository interface {
    registerService(service Service) error
    addSnapshot(snapshot SnapShot) error
    GetAllKeys() ([]string, error)
    RedisGetValue(key string) (string, error)
}

type RepoHandler struct{}

func (r *RepoHandler) registerService(service Service) error {
    c := REDIS.Get()
    defer c.Close()
    _, err := c.Do("SET", service.Name, service.URL)
    return err
}

func (r *RepoHandler) addSnapshot(snapshot SnapShot) error {
    return DB.Create(&snapshot).Error
}

func (r *RepoHandler) GetAllKeys() ([]string, error) {
    c := REDIS.Get()
    defer c.Close()
    keys, err := redis.Strings(c.Do("KEYS", "*"))
    return keys, err
}

func (r *RepoHandler) RedisGetValue(key string) (string, error ){
    c := REDIS.Get()
    defer c.Close()
    var value string
    reply, err := redis.Values(c.Do("MGET", key))
    redis.Scan(reply, &value)
    return value, err
}
