# day1

## 工程目录

![1722758616487](C:\Users\JunLian\AppData\Roaming\Typora\typora-user-images\1722758616487.png)

golang项目标准：https://github.com/golang-standards/project-layout/blob/master/README_zh.md

## 设置go代理

设置goproxy.io代理

go env -w GOPROXY="https://goproxy.io"

设置GO111MOUDLE

go env -w GO111MODULE="on"

## 项目搭建

创建workspace：go work init;    go mod init test.com/project-user;   go work use ./project-user

创建模块webcenter

在webcenter中导入gin： go get -u github.com/gin-gonic/gin