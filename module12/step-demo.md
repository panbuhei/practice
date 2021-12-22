# 安装 Istio
```
### 通过 istio 下载脚本下载 istio
root@node1:~# curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.12.0 sh -
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:--  0:02:10 --:--:--     0
curl: (28) Failed to connect to istio.io port 443: Connection timed out
```

如果因为网络问题无法正常下载，可以手动下载后上传至主机，下载地址：https://github.com/istio/istio/releases

```
root@node1:~# tar -xvf istio-1.12.0-linux-amd64.tar.gz
root@node1:~# cd istio-1.12.0/
root@node1:~/istio-1.12.0# cp bin/istioctl /usr/local/bin
root@node1:~/istio-1.12.0# istioctl install --set profile=demo -y
```
![install-istio](https://user-images.githubusercontent.com/83450378/147064347-2da1c911-e268-491d-ac4a-a7f33a27d225.jpg)



## 查看 Istio 是否安装成功
```
root@node1:~/istio-1.12.0# istioctl version
client version: 1.12.0
control plane version: 1.12.0
data plane version: 1.12.0 (2 proxies)

root@node1:~/istio-1.12.0# kubectl get pod -n istio-system
NAME                                   READY   STATUS    RESTARTS   AGE
istio-egressgateway-7f4864f59c-snb86   1/1     Running   0          9m33s
istio-ingressgateway-55d9fb9f-fdrh4    1/1     Running   0          9m33s
istiod-555d47cb65-6mssk                1/1     Running   0          9m37s
```

# 创建 httpserver namespace，并开启 Istio sidecar 的自动注入
```
root@node1:~# kubectl create ns httpserver

root@node1:~# kubectl label ns httpserver istio-injection=enabled
```
**PS：** 当你在一个命名空间中设置了 istio-injection=enabled 标签，且 Injection webhook 被启用后，任何新的 Pod 都有将在创建时自动添加 Sidecar。


# 创建 httpserver 服务
## 首先，通过执行 deployment 创建 pod
```
root@node1:~# kubectl create -f deployment.yaml

root@node1:~# kubectl get pod -n httpserver
NAME                          READY   STATUS    RESTARTS   AGE
httpserver-75f86d7954-hx72l   2/2     Running   0          2m17s
httpserver-75f86d7954-p4z75   2/2     Running   0          2m17s
httpserver-75f86d7954-w68ns   2/2     Running   0          2m17s
```

## 接着，构建 service
```
root@node1:~# kubectl create -f service.yaml 

root@node1:~# kubectl get svc -n httpserver
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
httpserver   ClusterIP   10.233.34.138   <none>        80/TCP    10s
```

## 然后，验证是否可以正常提供服务
```
root@node1:~# curl 10.233.34.138/healthz
ok

root@node1:~# curl 10.233.34.138/hello
hello [stranger]
===================Details of the http request header:============
X-Forwarded-Proto=[http]
X-Request-Id=[35c3ab74-968b-9983-ae00-8f6834e53d1c]
X-B3-Traceid=[b50d3e09badad20363d85b3291076055]
X-B3-Spanid=[63d85b3291076055]
X-B3-Sampled=[1]
User-Agent=[curl/7.68.0]
Accept=[*/*]
```


# 通过 Istio 为 httpserver 配置 TLS
## 创建 TLS 证书
### 生成 crt 和 key
```
root@node1:~# openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=httpserver Inc./CN=*.panbuhei.com' -keyout panbuhei.com.key -out panbuhei.com.crt
```

### 用 crt 和 key 创建一个 tls secret
**注意：** 要创建至 istio-system namespace
```
root@node1:~# kubectl create -n istio-system secret tls httpserver-credential --key=panbuhei.com.key --cert=panbuhei.com.crt
```

## 执行 istio-httpserver.yaml 创建规则
```
root@node1:~# kubectl create -f istio-httpserver.yaml -n httpserver
```

## 将 istio-ingressgateway service 类型修改为 NodePort
```
root@node1:~# kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      10.233.17.117   <none>        80/TCP,443/TCP                                                               27m
istio-ingressgateway   LoadBalancer   10.233.42.41    <pending>     15021:31077/TCP,80:31694/TCP,443:30650/TCP,31400:30291/TCP,15443:32423/TCP   27m
istiod                 ClusterIP      10.233.7.1      <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        27m

root@node1:~# kubectl edit svc istio-ingressgateway -n istio-system
service/istio-ingressgateway edited

root@node1:~# kubectl get svc -n istio-system
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP   10.233.17.117   <none>        80/TCP,443/TCP                                                               28m
istio-ingressgateway   NodePort    10.233.42.41    <none>        15021:31077/TCP,80:31694/TCP,443:30650/TCP,31400:30291/TCP,15443:32423/TCP   28m
istiod                 ClusterIP   10.233.7.1      <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        28m
```

## 浏览器访问
![httpserver-tls](https://user-images.githubusercontent.com/83450378/147064733-93e380c6-a5af-44cc-a3ab-4537af3a0d51.jpg)


