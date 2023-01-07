# mongo-plus

基于mongo go官方驱动的二次封装

## 快速开始

```shell
go get -u github.com/here-Leslie-Lau/mongo-plus/mongo
```

创建mongodb连接

```go
opts := []mongo.Option{
	// 要连接的数据库
	mongo.WithDatabase("test"),
	// 最大连接池数量
	mongo.WithMaxPoolSize(10),
	// 用户名
	mongo.WithUsername("your username"),
	// 密码
	mongo.WithPassword("your password"),
	// 连接url
	mongo.WithAddr("localhost:27017"),
}
conn, f := mongo.NewConn(opts...)
defer f()
```

获取collection对象

```go
type Demo struct{}

// Collection 实现mongo.Collection接口, 返回要操作的集合名
func (d *Demo) Collection() string {
	return "demo"
}

// 方法内获取collection对象
demo := &Demo{}
coll := conn.Collection(demo)
```

插入文档(insert)

```go
coll.InsertOne(ctx, document)
coll.InsertMany(ctx, documents)
```

查询文档

```go
// 查询name为leslie的单条文档
coll.Where("name", "leslie").FindOne(ctx, &document)
// 查询name为leslie的文档
coll.Where("name", "leslie").Find(ctx, &documents)
// 多条件查询
coll.Filter(map[string]interface{}{"name": "leslie", "age": 18}).FindOne(ctx, &document)
```


_其余文档补充中，更详细的用法参考test/chain_test.go_


## 初衷

*平时在工作或者自己写点小东东时，使用mongo官方提供的 [go driver](https://www.mongodb.com/docs/drivers/go/current/) ，总感觉哪里不方便。*

个人总结了下，有一下几点

- 当进行`mongodb`操作时，需要把官方驱动的各种Option对象准备好，再一口气传入。或许是`gorm`的调用链方式深得我心😄 ,所以也想封装成类似的方式。
- 官方驱动没有提供比较好的分页方式，_(例如:根据前端或客户端传入的页数/页码大小，获得相应的总页数/总条数)_ 每次都需要再次封装。
- 我认为一个库需要尽量屏蔽细节，使用者不应该多关注底层实现，开箱即用。_(比如开发者无需了解bson, $gt各种运算符, 分片等)_

**与官方mongodb驱动对比:**

- [ ] TODO

## 项目结构

.
├── LICENSE

├── makefile

├── mongo

│   ├── chain_cond.go   // 调用链条件拼接逻辑

│   ├── chain.go        // 核心结构体-chain定义

│   ├── collection.go   // collection接口定义

│   ├── config.go       // 连接mongodb配置

│   ├── conn.go         // 获取连接逻辑及一些初始化方法

│   ├── paginate.go     // 分页逻辑封装

│   └── type.go         // mongodb类型定义

├── README.md

└── test                  // 一些单元测试以及简单示例

## 核心代码

- [ ] TODO

## 版本管理

**v0.1(现在):**

1. 基本的curd
2. 分页封装
3. 达到开发者学习或基本使用`mongodb`的标准(基本功能支持)

**v0.2:**

1. gitbook文档支持
2. 完善的测试用例
3. 基准测试

**v0.3**

1. 事务支持
2. Aggregation支持
3. 操作集合、库级别的支持

**v2.0**

1. 去除官方驱动的依赖(待研究)

## 如何贡献

step one:

```shell
git checkout -b feature/要添加的功能描述 origin/master
# 开发自测完成后提交
git add .
git commit -m "功能描述"
```

step two:
发起**pull/request**

_tips: 或直接提issue_

## 捐赠
