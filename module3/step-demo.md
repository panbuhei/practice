1. 构建 httpserver 镜像
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

2. 本地运行 httpserver
```
root@ubuntu20:~/module3/httpServer# docker run -d httpserver:v1.0
4784531a3644a5ce784fb2402454ec44ff148c631d9d0a1b7fade70735aa7ef4
root@ubuntu20:~/module3/httpServer# docker ps 
CONTAINER ID   IMAGE             COMMAND         CREATED         STATUS         PORTS     NAMES
4784531a3644   httpserver:v1.0   "/httpserver"   3 seconds ago   Up 2 seconds   80/tcp    dazzling_easley
```

3. 通过 nsenter 进入容器查看 IP 配置
```
root@ubuntu20:~/module3/httpServer# docker inspect 4784531a3644 | grep -i pid
            "Pid": 65484,
            "PidMode": "",
            "PidsLimit": null,
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

4. 访问容器内的程序
```
root@ubuntu20:~/module3/httpServer# curl 172.17.0.2/hello
hello golang

root@ubuntu20:~/module3/httpServer# nsenter -t 65484 -n curl 127.0.0.1/healthz
200
```
