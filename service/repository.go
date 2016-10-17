package service

type repository interface {
    registerService(service Service) error
}

type repoHandler struct{}

func (r *repoHandler) registerService(service Service) error {
    return REDIS.Set(service.Name, service.URL, 0).Err()
}
