# go-micro-consul-grpc 踩坑记录

## 安装consul

这个就很简单了 到这里下载 然后文件解压以后的文件夹添加到本地的path里即可
[consul下载]( https://www.consul.io/downloads.html )

然后命令行执行
```
consul -dev -ui 
```
就运行起来了 然后到 localhost:8500 就可以看到注册中心页面了

## 安装 go-micro

```

go get github.com/micro/go-micro

```

## 安装 protobuf

```

go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u -v github.com/micro/protoc-gen-micro

```

## 安装 micro 工具包

```

go get github.com/micro/micro

```


## 写 proto

此处参考了 [li-peng]( https://www.cnblogs.com/li-peng/p/9598879.html )的博客 

这个例子中包含的数据类型很全面

## 生成模板

这里可以在proto所在目录下生成两个pb文件 这两个pb文件分别是grpc文件和micro文件
```
protoc --micro_out=. --go_out=. greeter.proto
```

## 书写serve端代码和client代码

这里需要注意的格式

server

1. 首先需要一个结构体来实现我们在proto中声明的方法 它会自动给我们生成一个接口 需要手动实现

        这里去看model 文件内容 

2. 然后再看server里面 注册consul 然后 把consul注册到micro 然后调用proto帮我们生成好的注册方法 固定填入当前的server和实现了服务方法的结构体即可完成服务注册

        这里去看server/main.go

3. 最后我们去看client 这是我们的客户端 客户端里面还是按照固定的流程 注册consul 注册micro 发现服务并且传入入参调用 


## 坑1 golang版本问题

go版本 一定要在1.13以上 否则各种包不存在的问题

## 坑2 etcd版本问题

如果报错 etcd/clientv3报错： etcd undefined: resolver.BuildOption

请跟我三步走

1. go mod edit -require=google.golang.org/grpc@v1.26.0
2. go get -u -x google.golang.org/grpc@v1.26.0
3. go build (或者goland里run)