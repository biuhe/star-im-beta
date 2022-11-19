# star-im
基于 Golang 的即时通信系统（IM）

## 核心功能
- 短信发送与接受
  - 文字、表情、图片、语音、视频等
- 访客模式
- 点对点私聊、群聊、广播、机器人等
- 心跳检测下线
- 快捷回复
- 撤回记录
- 拉黑等

## 需要技术栈

### 前端

- H5 Ajax 获取音频
- Websocket 发送消息
- Js/Vue 单页面App
- Mui/css3 等

### 后端

- websocket 组件转发消息
  github：https://github.com/gorilla/websocket
- channel 管道/goroutine 协程 提高并发性
- Gin 框架
  github：https://github.com/gin-gonic/gin

- template, swagger
- Viper
  配置管理库
  github：https://github.com/spf13/viper

- GORM
  ORM 框架
  官网：https://gorm.io/zh_CN/
  github：https://github.com/go-gorm/gorm

- swagger
  github：https://github.com/swaggo/swag

- logrus auth, logger,
- govalidator
  github：https://github.com/asaskevich/govalidator

- JWT
  github：https://github.com/golang-jwt/jwt/

- Set
  github：https://github.com/fatih/set

- Redis
  github：https://github.com/go-redis/redis
- SQL, NoSQL, MQ 等
-

## 架构划分
- 客户端：H5
- 接入层：TCP、HTTP/HTTPS、WebSocket
- 逻辑层：鉴权、登录、消息接收/发送、单聊/群聊/广播、心跳检测、用户/关系管理
- 存储层：SQL（MySQL）、NoSQL（Redis）、MQ



## 网络结构

App/浏览器/接口 - Websocket调用接口 > 应用服务器 - 推送、获取> 数据库



## 发送流程

``` mermaid
游客->鉴权
 用户->登录->鉴权->获取websocket连接->发送消息->根据消息类型推送到不同用户
```

## 网站

- 依赖包下载： https://pkg.go.dev/
- swagger：http://localhost:8081/swagger/index.html

## 步骤

```shell
# 进入项目
cd ./star-im

# 更新依赖
go mod tidy

```

## 其他

swagger
进入到 main.go所在目录执行

``` shell
swag init
```
