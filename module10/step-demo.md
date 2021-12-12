# 构建镜像
```
root@ubuntu20:~# cd module10/src/httpServer/

root@ubuntu20:~/module10/src/httpServer# ls
Dockerfile  go.mod  go.sum  main.go  Makefile  metrics

root@ubuntu20:~/module10/src/httpServer# make build
echo "building httpserver container"
building httpserver container
docker build -t wupanfeng035/httpserver:v1.0-metrics .
Sending build context to Docker daemon   25.6kB
Step 1/13 : FROM golang:1.16-alpine AS base
 ---> 065ac3b1a78f
Step 2/13 : ENV GO111MODULE=on     GOPATH=/go/src/project/     GOPROXY=https://goproxy.cn/,direct
 ---> Using cache
 ---> 2cabe6ea2b33
Step 3/13 : COPY main.go /go/src/project/github/panbuhei/practice/module10/src/httpServer/
 ---> Using cache
 ---> 9dbe377aab2f
Step 4/13 : COPY metrics/ /go/src/project/github/panbuhei/practice/module10/src/httpServer/metrics
 ---> Using cache
 ---> d5fca53a22f9
Step 5/13 : COPY go.mod /go/src/project/github/panbuhei/practice/module10/src/httpServer/
 ---> Using cache
 ---> fec54ee35628
Step 6/13 : COPY go.sum /go/src/project/github/panbuhei/practice/module10/src/httpServer/
 ---> Using cache
 ---> 82c7208cc1e6
Step 7/13 : WORKDIR /go/src/project/github/panbuhei/practice/module10/src/httpServer/
 ---> Using cache
 ---> 4c666eb3594f
Step 8/13 : RUN go mod download
 ---> Using cache
 ---> 2f67716bd8dc
Step 9/13 : RUN GOOS=linux GOARCH=amd64 go build -o httpserver .
 ---> Using cache
 ---> 7fe993841fd1
Step 10/13 : FROM alpine
 ---> c059bfaa849c
Step 11/13 : COPY --from=base /go/src/project/github/panbuhei/practice/module10/src/httpServer/httpserver /httpserver
 ---> Using cache
 ---> ae5fe03e2f30
Step 12/13 : EXPOSE 80
 ---> Using cache
 ---> f0c5163e116a
Step 13/13 : ENTRYPOINT ["/httpserver"]
 ---> Using cache
 ---> 83082ec1168c
Successfully built 83082ec1168c
Successfully tagged wupanfeng035/httpserver:v1.0-metrics
```

# 上传镜像
记得登陆 hub.docker.com
```
root@ubuntu20:~/module10/src/httpServer# make push
echo "pushing wupanfeng035/httpserver"
pushing wupanfeng035/httpserver
docker push wupanfeng035/httpserver:v1.0-metrics
The push refers to repository [docker.io/wupanfeng035/httpserver]
a47a91c0c049: Pushed 
8d3ac3489996: Layer already exists 
v1.0-metrics: digest: sha256:e9b7134194270e204c51cf716f141ac8bde64ebd0d309fb2382c242d3336a1fd size: 739
```

# 创建 httpserver 服务
## 通过 deployment 来创建 pod
```
root@node1:~/module10/config# kubectl apply -f deployment.yaml 
deployment.apps/httpserver created

root@node1:~/module10/config# kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
httpserver-75f86d7954-8p9vg   1/1     Running   0          33s
httpserver-75f86d7954-ggfd7   1/1     Running   0          33s
httpserver-75f86d7954-hv7fm   1/1     Running   0          33s
```

## 接着，构建 service
```
root@node1:~/module10/config# kubectl apply -f service.yaml 
service/httpserver created

root@node1:~/module10/config# kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpserver   ClusterIP   10.233.39.254   <none>        80/TCP    4s
kubernetes   ClusterIP   10.233.0.1      <none>        443/TCP   4d2h

### 验证一下 service 是否可以正常访问
root@node1:~/module10/config# curl 10.233.39.254/metrics
......
promhttp_metric_handler_requests_total{code="200"} 3
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

