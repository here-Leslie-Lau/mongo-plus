# 数据库级别使用文档

- [获取数据库实例](#获取数据库实例)
- [执行mongodb数据库管理命令](#执行mongodb数据库管理命令)
- [检查数据库服务状态](#检查数据库服务状态)
- [检查服务器是否可用](#检查服务器是否可用)
- [获取数据库的统计信息](#获取数据库的统计信息)
- [获取服务器的状态信息](#获取服务器的状态信息)

## 获取数据库实例

```go
func (c *Conn) GetDB() *mongo.Database
```

调用该方法可以获取到mongodb官方go驱动的`database`实例

## 执行mongodb数据库管理命令

```go
func (c *Conn) RunCommand(ctx context.Context, cmd bson.D, des interface{}) error
```

请求参数说明: ctx-标准库的context, cmd-命令集合,bson.D类型, des-命令执行结果,结构体指针

## 检查数据库服务状态

```go
func (c *Conn) IsMaster(ctx context.Context, des interface{}) error
```

请求参数说明: ctx-标准库的context, des-命令执行结果,结构体指针

用于检查服务器的状态，特别是用于确定服务器是否是主节点（Primary） 是否是一个副本集（Replica Set），以及其他关于服务器的基本信息 这个命令可以帮助你了解当前 MongoDB 部署的状态和拓扑

## 检查服务器是否可用

```go
func (c *Conn) Ping(ctx context.Context) error
```

请求参数说明: ctx-标准库的context

用于检查服务器是否可用

## 获取数据库的统计信息

```go
func (c *Conn) DbStats(ctx context.Context, des interface{}) error
```

请求参数说明: ctx-标准库的context, des-命令执行结果,结构体指针

用于获取数据库的统计信息 该命令返回的信息包括数据库的大小，对象的数量，索引的数量，以及其他统计信息

## 获取服务器的状态信息

```go
func (c *Conn) ServerStatus(ctx context.Context, des interface{}) error
```

请求参数说明: ctx-标准库的context, des-命令执行结果,结构体指针

用于获取服务器的状态信息 该命令返回的信息包括服务器的版本，操作系统，内存使用情况，连接数，以及其他统计信息
