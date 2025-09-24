package model

import "gorm.io/gorm"

type Article struct {
    Id      string
    Type    uint
    Content string

    // 媒体链接
    Medias []string

    Status    uint
    LikeCount uint

    Createby string
    gorm.Model
}