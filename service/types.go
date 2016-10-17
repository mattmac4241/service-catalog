package service

import (
    "github.com/jinzhu/gorm"
)

type SnapShot struct {
    gorm.Model
    Name       string
    URL        string
    Active     bool
}

type Service struct {
    Name        string      `json:"name"`
    URL         string      `json:"url"`
}
