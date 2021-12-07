# 多段构建-构建镜像
FROM golang:1.16-alpine AS base

# 配置国内代理
ENV GO111MODULE=on \
    GOPATH=/go/src/project/ \
    GOPROXY=https://goproxy.cn/,direct

# 拷贝代码
COPY main.go /go/src/project/github/panbuhei/practice/module10/src/httpServer/
COPY metrics/ /go/src/project/github/panbuhei/practice/module10/src/httpServer/metrics
COPY go.mod /go/src/project/github/panbuhei/practice/module10/src/httpServer/
COPY go.sum /go/src/project/github/panbuhei/practice/module10/src/httpServer/
WORKDIR /go/src/project/github/panbuhei/practice/module10/src/httpServer/
# 下载依赖
RUN go mod download

# 构建镜像
RUN GOOS=linux GOARCH=amd64 go build -o httpserver .

# 多段构建-启动
FROM alpine
COPY --from=base /go/src/project/github/panbuhei/practice/module10/src/httpServer/httpserver /httpserver
EXPOSE 80
ENTRYPOINT ["/httpserver"]