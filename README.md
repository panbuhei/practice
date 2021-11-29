# 极客时间云原生训练营作业
## 模块二作业：httpServer
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200

## 模块三作业：构建本地镜像
1. 编写 Dockerfile 将练习模块二作业编写的 httpserver 容器化(请思考有哪些最佳实践可以引入到 Dockerfile 中来)
2. 将镜像推送至 Docker 官方镜像仓库
3. 通过 Docker 命令本地启动 httpserver
4. 通过 nsenter 进入容器查看 IP 配置
### 演示步骤：
- [step-demo.md](./module3/step-demo.md)
- "学习委员-麟" 的模板：[dockerfile-template](./jpg/dockerfile-template.jpg)

## 模块八作业：将 httpserver 部署到 kubernetes 集群
**PS：** 由于 Go 能力有限，代码使用的是孟老师的 grace 分支
### 第一部分
1. 优雅启动
2. 优雅终止
3. 资源需求和 QoS 保证
4. 探活
5. 日常运维需求，日志等级
6. 配置和代码分离

### 第二部分
尝试用 Service, Ingress 将服务发布给集群外部的调用方
1. Service
2. Ingress

可以考虑的细节
- 如何确保整个应用的高可用
- 如何通过证书保证 httpServer 的通讯安全
