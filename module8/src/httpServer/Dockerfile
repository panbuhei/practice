# 多段构建-构建镜像
FROM golang:1.16-alpine AS base

# 配置国内代理
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn/,direct

# 拷贝代码
COPY main/main.go /go/src/project/
COPY go.mod /go/src/project/
COPY go.sum /go/src/project/
WORKDIR /go/src/project/
# 下载依赖
RUN go mod download

# 构建镜像
RUN GOOS=linux GOARCH=amd64 go build -o httpserver ./main.go

# 多段构建-启动
FROM alpine
COPY --from=base /go/src/project/httpserver /httpserver
EXPOSE 80
ENTRYPOINT ["/httpserver"]
