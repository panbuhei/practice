# 构建镜像
```
root@k8s001:~# cd module8/src/httpServer

root@k8s001:~/module8/src/httpServer# ls
Dockerfile  go.mod  go.sum  main  Makefile

root@k8s001:~/module8/src/httpServer# make build
echo "building httpserver container"
building httpserver container
docker build -t wupanfeng035/httpserver:v3.0 .
Sending build context to Docker daemon  8.192kB
Step 1/12 : FROM golang:1.16-alpine AS base
 ---> eee5af307da8
Step 2/12 : ENV GO111MODULE=on     GOPROXY=https://goproxy.cn/,direct
 ---> Running in a68e0c8ea77d
Removing intermediate container a68e0c8ea77d
 ---> 544aa1375d3a
Step 3/12 : COPY main/main.go /go/src/project/
 ---> dbfd21ec926b
Step 4/12 : COPY go.mod /go/src/project/
 ---> 87440388af13
Step 5/12 : COPY go.sum /go/src/project/
 ---> 433e494fa7ef
Step 6/12 : WORKDIR /go/src/project/
 ---> Running in 0fed78eee859
Removing intermediate container 0fed78eee859
 ---> 16865a11edde
Step 7/12 : RUN go mod download
 ---> Running in 17dae01b945e
Removing intermediate container 17dae01b945e
 ---> 6b2836c340c2
Step 8/12 : RUN GOOS=linux GOARCH=amd64 go build -o httpserver ./main.go
 ---> Running in 5c0d9f49552c
Removing intermediate container 5c0d9f49552c
 ---> cdb30801094a
Step 9/12 : FROM alpine
 ---> c059bfaa849c
Step 10/12 : COPY --from=base /go/src/project/httpserver /httpserver
 ---> 33b919fd0973
Step 11/12 : EXPOSE 80
 ---> Running in bfbe301be1c1
Removing intermediate container bfbe301be1c1
 ---> 63af7491910b
Step 12/12 : ENTRYPOINT ["/httpserver"]
 ---> Running in eda2730be072
Removing intermediate container eda2730be072
 ---> 63ac8f26dc01
Successfully built 63ac8f26dc01
Successfully tagged wupanfeng035/httpserver:v3.0
```

# 上传镜像
记得登陆 hub.docker.com
```
root@k8s001:~/module8/src/httpServer# make push
echo "pushing wupanfeng035/httpserver"
pushing wupanfeng035/httpserver
docker push wupanfeng035/httpserver:v3.0
The push refers to repository [docker.io/wupanfeng035/httpserver]
b43e87e8772b: Pushed 
8d3ac3489996: Layer already exists 
v3.0: digest: sha256:10cdb39c1ba440000f77829c311ad9e619eeaffa4079fdb24202a8d5349b56cb size: 739
```

# 创建 httpserver 服务
## 首先，通过 deployment 来创建 pod
```
root@k8s001:~/module8/src/httpServer# cd ../../config/

root@k8s001:~/module8/config# ls
deployment.yaml  ingress-nginx.yml  secret.yml  service.yml

root@k8s001:~/module8/config# k apply -f deployment.yaml 
deployment.apps/httpserver created
```
```
root@k8s001:~/module8/config# k get pod
NAME                         READY   STATUS    RESTARTS   AGE
httpserver-fd6d69cf5-br464   1/1     Running   0          6m58s
httpserver-fd6d69cf5-g8l79   1/1     Running   0          6m58s
httpserver-fd6d69cf5-hq89t   1/1     Running   0          6m58s
```

## 接着，构建 service
```
root@k8s001:~/module8/config# k apply -f service.yml 
service/httpserver created

root@k8s001:~/module8/config# k get svc 
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
httpserver   ClusterIP   10.96.253.94   <none>        80/TCP    9s
kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP   20d

### 验证一下 service 是否可以正常访问
root@k8s001:~/module8/config# curl 10.96.253.94
hello [stranger]
===================Details of the http request header:============
User-Agent=[curl/7.68.0]
Accept=[*/*]
```

### 然后，创建 TLS 证书
```
### 生成 crt 和 key
# openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout httpserver.key -out httpserver.crt -subj "/CN=panbuhei.com/O=cncamp"

### 通过 base64 进行转码后，填写至 secret 中对应的 tls.key 和 tls.crt 位置
# cat httpserver.key | base64 -w 0
# cat httpserver.crt | base64 -w 0

### 创建 secret
root@k8s001:~/module8/config# k create -f secret.yml 
secret/httpserver-tls created

root@k8s001:~/module8/config# k get secret
NAME                  TYPE                                  DATA   AGE
default-token-l5wsz   kubernetes.io/service-account-token   3      20d
httpserver-tls        kubernetes.io/tls                     2      6s
```

### 最后，创建 ingress
```
root@k8s001:~/module8/config# k apply -f ingress-nginx.yml 
ingress.networking.k8s.io/httpserver-gateway created

root@k8s001:~/module8/config# k get ingress
NAME                 CLASS    HOSTS          ADDRESS      PORTS     AGE
httpserver-gateway   <none>   panbuhei.com   10.0.30.24   80, 443   34s
```

# 通过浏览器访问 httpserver 服务
```
### 查看 ingress 的 svc
root@k8s001:~/module8/config# k get svc -n ingress-nginx
NAME                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             NodePort    10.109.238.20   <none>        80:30752/TCP,443:30020/TCP   5h29m
ingress-nginx-controller-admission   ClusterIP   10.108.47.16    <none>        443/TCP                      5h29m
```

![httpserver](https://user-images.githubusercontent.com/83450378/143858775-17ed5711-03c2-40a1-866c-96904c51297c.png)

