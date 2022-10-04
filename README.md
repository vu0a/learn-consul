## learn-consul

Learn then [consul](https://www.consul.io/) with [Go](https://go.dev).

## 常见的注册中心

- Netflix Eureka
- Alibaba Nacos
- HashiCore Consul
- Apache Zookeeper
- CoreOS Etcd
- CNCF CoreDNS

![常见的注册中心对比](assets/20221004_141743_image.png)

## Consul 角色

- 客户端(`client`)：无状态，将 HTTP 和 DNS 接口请求转发给局域网内的服务端集群。
- 服务端(`server`)：保存配置信息，高可用集群，每个数据中心的 server 数量推荐为 3 个或 5 个
- 代理(`agent`)：客户端和服务端统称为 `agent`

![](assets/20221004_143241_image.png)

## 安装 Consul

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor |  tee /usr/share/keyrings/hashicorp-archive-keyring.gpg && \
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list && \
apt update -y && apt install consul -y && \
consul version
```

## 开发模型

### ipv4

```bash
consul agent -dev -client=0.0.0.0
```

### ipv6

```bash
consul agent -dev -client=[::]
```

### 访问 web

- `http://ipv4:8500` 或
- `http://[ipv6]:8500`

## 将 consul 集成到 go

```bash
go get -u github.com/hashicorp/consul/api
```
