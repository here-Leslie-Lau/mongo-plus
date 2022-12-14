# mongo-plus

基于mongo go官方驱动的二次封装

[![](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://img.shields.io/badge/build-passing-brightgreen.svg)&ensp;&ensp;&ensp;&ensp;[![](https://img.shields.io/badge/version-v0.1-orange.svg)](https://img.shields.io/badge/version-v0.1-orange.svg)&ensp;&ensp;&ensp;&ensp;[![](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)

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
coll.InsertOne(document)
coll.InsertMany(documents)
```

查询文档

```go
// 查询name为leslie的单条文档
coll.Where("name", "leslie").FindOne(&document)
// 查询name为leslie的文档
coll.Where("name", "leslie").Find(&documents)
// 多条件查询
coll.Filter(map[string]interface{}{"name": "leslie", "age": 18}).FindOne(&document)
```


_其余文档补充中，更详细的用法参考test/chain_test.go_


## Benchmark

```shell
make benchmark
```

输出结果(其他方法补充中):

```shell
$ make benchmark
cd test && go test -bench=. -benchmem -benchtime=1s -run=none
goos: darwin
goarch: arm64
pkg: github.com/here-Leslie-Lau/mongo-plus/test
BenchmarkFindOne-8    	   11860	     97556 ns/op	    6617 B/op	      79 allocs/op
BenchmarkFind-8       	   12999	     92032 ns/op	    6417 B/op	      80 allocs/op
BenchmarkInString-8   	   12180	     98361 ns/op	    6867 B/op	      88 allocs/op
BenchmarkInInt64-8    	   12309	     99718 ns/op	    6484 B/op	      80 allocs/op
PASS
ok  	github.com/here-Leslie-Lau/mongo-plus/test	8.226s
```

## 初衷

*平时在工作或者自己写点小东东时，使用mongo官方提供的 [go driver](https://www.mongodb.com/docs/drivers/go/current/) ，总感觉哪里不方便。*

个人总结了下，有一下几点

- 当进行`mongodb`操作时，需要把官方驱动的各种Option对象准备好，再一口气传入。或许是`gorm`的调用链方式深得我心😄 ,所以也想封装成类似的方式。
- 官方驱动没有提供比较好的分页方式，_(例如:根据前端或客户端传入的页数/页码大小，获得相应的总页数/总条数)_ 每次都需要再次封装。
- 我认为一个库需要尽量屏蔽细节，使用者不应该多关注底层实现，开箱即用。_(比如开发者无需了解bson, $gt各种运算符, 分片等)_

**与官方mongodb驱动对比:**

- [ ] TODO

## 项目结构

```shell
.
├── LICENSE
├── README.md			// 项目介绍文档
├── go.mod
├── go.sum
├── makefile			// 一些初始化工具
├── mongo
│   ├── chain.go		// 核心结构体-chain定义, 与操作mongodb方法封装
│   ├── chain_cond.go	// 调用链条件拼接逻辑
│   ├── collection.go	// collection接口定义
│   ├── config.go		// 连接mongodb配置定义
│   ├── conn.go			// 获取连接逻辑及一些初始化方法
│   ├── paginate.go		// 分页逻辑封装
│   └── type.go			// mongodb类型定义
└── test
    ├── bench_test.go	// golang基准测试
    ├── chain_test.go	// 单元测试与用法事例
    └── conn_test.go	// 测试用例的初始化封装
```

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
