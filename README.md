# gRPC

> 什么是gRPC
>
gRPC 是 RPC 的一种，它使用 Protocol Buffer (简称 Protobuf) 作为序列化格式，Protocol Buffer 是来自 google 的序列化框架，比
Json 更加轻便高效，同时基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP
连接上的多复用请求等特性。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。用 protoc 就能使用 proto 文件帮助我们生成上面的
option 层代码。
在 gRPC 中，客户端应用程序可以直接在另一台计算机上的服务器应用程序上调用方法，就好像它是本地对象一样，从而使您更轻松地创建分布式应用程序和服务。

## 通信模式

### Unary RPC（单向 RPC）

* 客户端发送一个请求，服务器返回一个响应。这是最简单的 gRPC 通信模式，类似于普通的函数调用
* 特点：单个请求和单个响应

```shell
# 通过protoc生成Go的gRPC代码
protoc --go_out=. --go-grpc_out=. message.proto
```

### Server Streaming RPC（服务端流式 RPC）

* 客户端发送一个请求，服务器返回一个响应流。客户端接收完服务器发送的所有数据后通信才会结束
* 特点：单个请求和多个响应

```shell
# 通过protoc生成Go的gRPC代码
protoc --go_out=. --go-grpc_out=. message.proto
```

### Client Streaming RPC（客户端流式 RPC）

* 客户端发送一个请求流，服务器在接收到客户端的所有数据后返回一个响应
* 特点：多个请求和单个响应

```shell
# 通过protoc生成Go的gRPC代码
protoc --go_out=. --go-grpc_out=. message.proto
```

### Bidirectional Streaming RPC（双向流式 RPC）

* 客户端和服务器都可以发送数据流，双方可以在任意时间读取和写入。客户端和服务器以流的方式进行持续的通信，彼此独立发送和接收数据
* 特点：多个请求和多个响应

```shell
# 通过protoc生成Go的gRPC代码
protoc --go_out=. --go-grpc_out=. message.proto
```