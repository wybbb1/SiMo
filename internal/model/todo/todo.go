package model

import "gorm.io/gorm"

type Todo struct {
    Id string
    Status uint
    CreateBy string
    gorm.Model
}

const (
    Undone = 0
    Done = 1
)