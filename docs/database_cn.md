# 数据库级别使用文档

- [获取数据库实例](#获取数据库实例)
- [执行mongodb数据库管理命令](#执行mongodb数据库管理命令)
- [检查数据库服务状态](#检查数据库服务状态)

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
