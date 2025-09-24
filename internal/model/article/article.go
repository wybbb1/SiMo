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

const (
    // type
    Markdown = 0
    Talk = 1
    // 白板暂定复用Article表，但是只使用Id, Type, Content（存储白板路径），status, CreateBy, gorm.Model字段
    Whiteboard = 2

    // status
    // 草稿
    Draft = 0
    // 已发布
    Publish = 1
    // 设为隐私
    Private = 2 // 如果设置private则所有人（包括超级管理员）无法看见你的内容，除非超级管理员在部署的机器上直接查看
    // 以上问题也可以通过加密解决
    // 已删除
    Delete = 3
)