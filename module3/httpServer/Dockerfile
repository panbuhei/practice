FROM golang:1.16-alpine AS base
COPY main/main.go /go/src/project/
WORKDIR /go/src/project/
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o httpserver ./main.go

FROM alpine
COPY --from=base /go/src/project/httpserver /httpserver
EXPOSE 80
ENTRYPOINT ["/httpserver"]
