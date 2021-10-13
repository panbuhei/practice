# 1. 构建 httpserver 镜像
```
root@ubuntu20:~# cd module3/httpServer/

root@ubuntu20:~/module3/httpServer# ls
Dockerfile  main

root@ubuntu20:~/module3/httpServer# docker build -t httpserver:v1.0 .
Sending build context to Docker daemon  4.608kB
Step 1/10 : FROM golang:1.16-alpine AS base
 ---> 1b35785aa3c4
Step 2/10 : COPY main/main.go /go/src/project/
 ---> 8dba123a650e
Step 3/10 : WORKDIR /go/src/project/
 ---> Running in 8ea5296d8431
Removing intermediate container 8ea5296d8431
 ---> 3f86b5f31679
Step 4/10 : ENV GOOS=linux
 ---> Running in a7ace5209889
Removing intermediate container a7ace5209889
 ---> a3cffde1dc59
Step 5/10 : ENV GOARCH=amd64
 ---> Running in d344e63d6f68
Removing intermediate container d344e63d6f68
 ---> a7d956fc031b
Step 6/10 : RUN go build -o httpserver ./main.go
 ---> Running in a30fd32f22c0
Removing intermediate container a30fd32f22c0
 ---> a06962873de4
Step 7/10 : FROM alpine
 ---> 14119a10abf4
Step 8/10 : COPY --from=base /go/src/project/httpserver /httpserver
 ---> a310612f6cf3
Step 9/10 : EXPOSE 80
 ---> Running in 65ed4e01d954
Removing intermediate container 65ed4e01d954
 ---> a36861c0f43e
Step 10/10 : ENTRYPOINT ["/httpserver"]
 ---> Running in c4d0ebaa6deb
Removing intermediate container c4d0ebaa6deb
 ---> faa3d17c6744
Successfully built faa3d17c6744
Successfully tagged httpserver:v1.0
```

# 2. 查看生成的镜像 httpserver
```
root@ubuntu20:~/module3/httpServer# docker images
REPOSITORY   TAG            IMAGE ID       CREATED          SIZE
httpserver   v1.0           faa3d17c6744   33 minutes ago   11.9MB
<none>       <none>         a06962873de4   33 minutes ago   308MB
golang       1.16-alpine    1b35785aa3c4   5 days ago       302MB
alpine       latest         14119a10abf4   6 weeks ago      5.6MB
```
镜像 a06962873de4 是多段构建的第0个构建(base)，最后生成的 httpserver 镜像只有 11.9MB


# 3. 将镜像推送至 Docker 官方镜像仓库
## 3.1 创建一个镜像仓库
注册连接：[https://hub.docker.com/](https://hub.docker.com/)，已注册账号，可[直接登陆](https://id.docker.com/login/?next=%2Fid%2Foauth%2Fauthorize%2F%3Fclient_id%3D43f17c5f-9ba4-4f13-853d-9d0074e349a7%26nonce%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0M2YxN2M1Zi05YmE0LTRmMTMtODUzZC05ZDAwNzRlMzQ5YTciLCJleHAiOiIyMDIxLTEwLTEzVDEwOjM2OjEzLjUzMTAwMDc3NFoiLCJpYXQiOiIyMDIxLTEwLTEzVDEwOjMxOjEzLjUzMTAwMTc3MVoiLCJyZnAiOiI1V2h1ZGtVd0huMFQwQ19hV2tJend3PT0iLCJ0YXJnZXRfbGlua191cmkiOiJodHRwczovL2h1Yi5kb2NrZXIuY29tIn0._ptan42erloKiofEFH7JwRuJAZfjd1Xgv70oYmaSpmo%26redirect_uri%3Dhttps%253A%252F%252Fhub.docker.com%252Fsso%252Fcallback%26response_type%3Dcode%26scope%3Dopenid%26state%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0M2YxN2M1Zi05YmE0LTRmMTMtODUzZC05ZDAwNzRlMzQ5YTciLCJleHAiOiIyMDIxLTEwLTEzVDEwOjM2OjEzLjUzMTAwMDc3NFoiLCJpYXQiOiIyMDIxLTEwLTEzVDEwOjMxOjEzLjUzMTAwMTc3MVoiLCJyZnAiOiI1V2h1ZGtVd0huMFQwQ19hV2tJend3PT0iLCJ0YXJnZXRfbGlua191cmkiOiJodHRwczovL2h1Yi5kb2NrZXIuY29tIn0._ptan42erloKiofEFH7JwRuJAZfjd1Xgv70oYmaSpmo)

![image](https://user-images.githubusercontent.com/83450378/137117357-ac4bb1ab-9670-4e95-b0e4-cdaf2584b4f7.png)

![image](https://user-images.githubusercontent.com/83450378/137117697-6ffb889d-1452-42f4-9c03-255a77e95f60.png)

## 3.2 为 httpserver 镜像打标签
```
docker tag faa3d17c6744 xxxxxxxxxx035/cncamp101/httpserver:v1.0
```

docker tag 命令示例：
```
docker tag 0e5574283393 hub.docker.com/cncamp/httpserver:v1.0
```
- hub.docker.com：镜像仓库地址，如果不填，默认为 hub.docker.co
- cncamp: repositry
- httpserver：镜像名
- v1.0：tag，常用来记录版本信息


## 3.3 登陆 Docker 官方镜像仓库
```
root@ubuntu20:~/module3/httpServer# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: xxxxxxxxxx035
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

## 3.4 上传到仓库
```
root@ubuntu20:~/module3/httpServer# docker push wupanfeng035/httpserver:v1.0
The push refers to repository [docker.io/wupanfeng035/httpserver]
b633ed876956: Pushed 
e2eb06d8af82: Pushed 
v1.0: digest: sha256:bd8642df2d30b900b1ded9e68b27794c28f60346650b52614826a04149370d66 size: 739
```


# 4. 本地运行 httpserver
```
root@ubuntu20:~/module3/httpServer# docker run -d httpserver:v1.0
4784531a3644a5ce784fb2402454ec44ff148c631d9d0a1b7fade70735aa7ef4
root@ubuntu20:~/module3/httpServer# docker ps 
CONTAINER ID   IMAGE             COMMAND         CREATED         STATUS         PORTS     NAMES
4784531a3644   httpserver:v1.0   "/httpserver"   3 seconds ago   Up 2 seconds   80/tcp    dazzling_easley
```

# 5. 通过 nsenter 进入容器查看 IP 配置
## 5.1 查看容器的 pid
```
root@ubuntu20:~/module3/httpServer# docker inspect 4784531a3644 | grep -i pid
            "Pid": 65484,
            "PidMode": "",
            "PidsLimit": null,
```
## 5.2 通过 nsenter 查看容器的 IP 配置
```
root@ubuntu20:~/module3/httpServer# nsenter -t 65484 -n ip a 
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
207: eth0@if208: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

# 6. 访问容器内的程序
```
root@ubuntu20:~/module3/httpServer# curl 172.17.0.2/hello
hello golang
```

```
root@ubuntu20:~/module3/httpServer# nsenter -t 65484 -n curl 127.0.0.1/healthz
200
```
