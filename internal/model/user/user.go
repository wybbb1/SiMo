package model

import "gorm.io/gorm"


type User struct {
    Id string
    Username string
    Password string
    Status uint
    Avater string
    Description string // 个性签名
    Email string // 个人邮箱
    Link string // 个人链接
    Online bool // 是否在线
    gorm.Model
}

const (
    // id
    NotUser = "0"
    Admin = "1"

    // status
    // 待审核
    Pending = 0
    // 正常
    Normal = 1
    // 封禁
    Banned = 2

    // online
    Offline = false
    Online = true
)