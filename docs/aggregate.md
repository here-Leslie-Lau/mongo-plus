# aggregate(聚合操作)使用文档

**调用Aggregate方法完成聚合操作**

```go
func (ch *Chain) Aggregate(des interface{}, stages ...bson.D) error
```

请求参数说明: des-方法执行后的结果集, stages-聚合操作内的stage(阶段)

该stage可以调用mongo-plus提供的方法直接获取, 下文会说明。如果不满足也可以自定义传入
