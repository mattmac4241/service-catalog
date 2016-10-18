package service

type repository interface {
    registerService(service Service) error
    addSnapshot(snapshot SnapShot) error
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
