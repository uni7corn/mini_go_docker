FROM golang:1.20.2-alpine as builder

#设置环境变量
ENV  GO111MODULE=on \
    # normal proxy https://goproxy.io
    # aliyun proxy https://mirrors.aliyun.com/goproxy/
     GOPROXY=https://goproxy.cn \
     CGO_ENABLED=0  \
     GOOS=linux  \
     GOARCH=amd64 \
     TZ=Asia/Shanghai

#声明一个工作目录，进行编译操作
WORKDIR /workspace

#复制代码到容器中
COPY . .

VOLUME /workspace

#执行go 编译命令
RUN  go mod download \
    && go build -o mini_docker cmd/main.go










