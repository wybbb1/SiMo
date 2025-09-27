# 项目定位

SiMo 是一款专为思维记录而设计的自托管平台，以生活状态区分出了**以“沉浸专注”为主的写作空间**和**以“放松”为主的动态空间**，编写方便，分享快捷，写作沉浸等特点让你随时随地放飞思维，分享想法。

# 项目亮点
- 轻量高性能：
- 一人部署多用户使用：单机多用户，支持多人同时在线
- 数据隐私：所有数据本地保存
- 沉浸写作：支持多种写作模式，帮助用户更好地沉浸于创作中
- 分享便捷：主页点击其他用户头像一键跳转查看其动态，快速了解其他用户的想法

# 开发须知

## 后端
- go 1.25.1
- 数据存储：sqlite
- 缓存：Ristretto

## 前端
- vue

## 数据库表设计
```go
type Article struct {
    Id string
    Type uint 
    Content string

    // 以下也可以合并为多媒体链接media []string
    // 图片链接
    Images []string
    // 视频与音乐链接（可能取消）
    Medias []string

    Status uint
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
    draft = 0
    // 已发布
    publish = 1
    // 设为隐私
    private = 2 // 如果设置private则所有人（包括超级管理员）无法看见你的内容，除非超级管理员在部署的机器上直接查看
    // 以上问题也可以通过加密解决
    // 已删除
    delete = 3
)

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
```

# 权限管理
- 超级管理员：admin，拥有所有权限，包括封禁用户、审核用户、删除用户**已发布**的文章和白板等
- 普通用户：注册即为普通用户，拥有发布文章、说说、点赞等权限
- 游客：只能浏览以及点赞公开的文章和说说

# 关于OAuth登录
SiMo 支持 GitHub OAuth 登录，但是用户需要在 GitHub 上注册 OAuth 应用，并将 ClientID 和 ClientSecret 以及callback URL填入 SiMo 的配置文件中。
