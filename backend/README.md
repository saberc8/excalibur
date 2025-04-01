# aquila

## 项目介绍

Aquila 是一个基于 Go 语言的后台管理系统，提供了一系列 API 接口用于管理和操作数据。

## 打包

GOOS=linux GOARCH=amd64 go build -o myapp main.go

## 运行 
go mod tidy
go run main.go

## 目录结构

```
aquila
├── README.md
├── api // api接口
├── config // 配置文件
├── constants // 常量
├── enum // 枚举
├── global // 全局变量
├── initialize // 初始化
├── log
├── middleware // 中间件
├── model // 数据库模型
├── router // 路由
├── utils // 工具
└── main.go
```

panic: listen tcp :9090: bind: address already in use

lsof -i :9090

sudo netstat -nlp | grep :9090

kill -9 <PID>