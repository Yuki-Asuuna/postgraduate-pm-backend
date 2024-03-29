# postgraduate-pm-backend

**<u>注意：如果图片加载不出来，请挂vpn后刷新页面。</u>**

## 服务器配置

服务器ip&port：124.221.197.218:8000

系统信息：Centos 8.0 64位

系统配置：2vCPU 4GiB（I/O优化）Tencent Cloud

## 后端整体架构图

![后端整体架构图](./assets/architecture.png)


## 访问示例

没有申请域名，目前只能通过ip+端口号访问

提供2种请求方式：可以直接访问服务器，也可以通过反向代理服务器访问

### 方式一

**直接访问后端服务器**（<u>目前调试阶段允许这种访问，之后一律走反向代理，拒绝其余请求</u>）

Request：

```
http://124.221.197.218:8000/ping
```

Response：

```json
{
  "code": 0,
  "message": "OK",
  "result": "pong"
}
```


## Redis集群

由于安全问题，Redis不对公网开放 （docker部署）

服务器ip：124.221.197.218（与服务器ip相同）

Port：6379

用户名：/

密码：/

---

## Chevereto（图床服务）

服务器ip：124.221.197.218（与服务器ip相同）

Port：80

使用教程：https://chevereto-free.github.io/api/#api-key



---

## Elastic Search + Kibana(分布式日志管理服务)

目前仅接入了后端日志，预计之后时间充裕可接入前端以及客户端日志。

服务器ip&port：124.221.197.218:5601

---

## MinIO

一个文件管理系统，用于保存上传的文件

服务器ip&port：124.221.197.218:9001

---

## MySQL

服务器：124.221.197.218:3307

数据库版本：MySQL 5.7

从数据库（备用数据库）：/
