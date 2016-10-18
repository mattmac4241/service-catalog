package service

import (
    "github.com/jinzhu/gorm"
)

type SnapShot struct {
    gorm.Model
    Name            string
    URL             string
    Active          bool
    ResponseStatus  int
}

type Service struct {
    Name        string      `json:"name"`
    URL         string      `json:"url"`
}
