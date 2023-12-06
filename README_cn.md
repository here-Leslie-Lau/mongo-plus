# mongo-plus  [![](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://img.shields.io/badge/build-passing-brightgreen.svg) [![](https://img.shields.io/badge/version-v1.0-orange.svg)](https://img.shields.io/badge/version-v1.0-orange.svg) [![](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)

中文 | [English](https://github.com/here-Leslie-Lau/mongo-plus)

![](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/mongo-plus.png)

基于mongo go官方驱动的二次封装

## 特性

- 调用链操作, 自由组合条件
- api友好
- 支持Context
- 开箱即用
- 分页查询支持
- [简易的聚合(aggregate)支持](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate_cn.md)
- [数据库管理命令操作](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/database_cn.md)
- 持续更新ing

## 快速开始

```shell
go get -u github.com/here-Leslie-Lau/mongo-plus
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

ctx支持

```go
coll = coll.WithCtx(ctx)
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

查询满足条件的文档数

```go
// 查询name为leslie的文档条数
cnt, err := coll.Where("name", "leslie").Count()
```

排序

```go
// 根据value字段升序查询
coll.Sort(mongo.SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&documents)
```

分页

```go
f := &mongo.PageFilter{
	// 当前查询第几页
	PageNum:  1,
	// 每页多少条
	PageSize: 2,
}

// 根据条件将匹配文档塞入documents内, 并将总条数与总页数放入f内
coll.Paginate(f, &documents)
```

逻辑操作

```go
// 找到age大于18的单条记录
coll.Gt("age", 18).FindOne(&document)
// 找到age小于18的单条记录
coll.Lt("age", 18).FindOne(&document)
// 找到age大于等于18的单条记录
coll.Gte("age", 18).FindOne(&document)
// 找到age不等于100的单条记录
coll.NotEq("age", 100).FindOne(&document)
// ...其他方法可以参考mongo/chain_cond.go
```

指定要查询的字段

```go
// 查询结果只对"name"字段赋值, 调用该方法后默认不对"_id"字段赋值
coll.Projection("name").Find(&documents)
```

更新或插入一条记录

```go
// 将age字段更新为18
content := map[string]interface{}{"age": 18}
// 如果筛选条件不存在, 要插入的默认值
default := map[string]interface{}{"name": "leslie"}

conn.Where("name", "leslie").UpsertOne(content, default)
// 期望结果, 如果name为leslie的文档存在, 则将age更新为18, 否则插入一条{"name": "leslie", "age": 18}的文档
```

Or查询(或运算查询)

```go
// 单条件
orMap := map[string]interface{}{"age": 18, "name": "leslie"}
// 查询name为leslie或者age为18的文档
conn.Or(orMap).Find(&documents)

// 多条件
orMap1 := map[string]interface{}{"name": "leslie", "age": 22}
orMap2 := map[string]interface{}{"name": "skyle", "age": 78}
// 查询name为leslie,age为22或者name为skyle,age为78的文档
conn.Ors(orMap1, orMap2).Find(&documents)
```

[Aggregate操作](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate_cn.md)

[数据库管理命令操作](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/database_cn.md)

_其余文档补充中，更详细的用法参考test/chain_test.go_

## Benchmark

machine: Macbook Pro M1 <br/>
memory: 8G

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
BenchmarkFindOne-8    	   11989	     96971 ns/op	    6906 B/op	      81 allocs/op
BenchmarkFind-8       	   13082	     91575 ns/op	    6449 B/op	      81 allocs/op
BenchmarkInString-8   	   12154	     98752 ns/op	    6899 B/op	      89 allocs/op
BenchmarkInInt64-8    	   12380	     97217 ns/op	    6517 B/op	      81 allocs/op
BenchmarkSort-8       	   10000	    101614 ns/op	    6811 B/op	      86 allocs/op
PASS
ok  	github.com/here-Leslie-Lau/mongo-plus/test	9.246s
```

## 初衷

*平时在工作或者自己写点小东东时，使用mongo官方提供的 [go driver](https://www.mongodb.com/docs/drivers/go/current/) ，总感觉哪里不方便。*

个人总结了下，有以下几点

- 当进行`mongodb`操作时，需要把官方驱动的各种Option对象准备好，再一口气传入。或许是`gorm`的调用链方式深得我心😄 ,所以也想封装成类似的方式。
- 官方驱动没有提供比较好的分页方式，_(例如:根据前端或客户端传入的页数/页码大小，获得相应的总页数/总条数)_ 每次都需要再次封装。
- 我认为一个库需要尽量屏蔽细节，使用者不应该多关注底层实现，开箱即用。_(比如开发者无需了解bson, $gt各种运算符, 分片等)_

## 项目结构

```shell
.
├── docs							// 项目文档
├── go.mod
├── go.sum
├── LICENSE
├── makefile						// 一些初始化工具
├── mongo							// 核心逻辑包
│   ├── aggregate.go				// mongodb聚合操作逻辑(aggregate)
│   ├── aggregate_group.go			// 聚合操作中group逻辑
│   ├── chain_cond.go				// 调用链条件拼接逻辑
│   ├── chain.go					// 核心结构体-chain定义, 与操作mongodb方法封装
│   ├── collection.go				// collection接口定义
│   ├── config.go					// 连接mongodb配置定义
│   ├── conn.go						// 获取连接逻辑及一些初始化方法
│   ├── database.go                 // 数据库管理相关方法
│   ├── paginate.go					// 分页逻辑封装
│   └── type.go						// mongodb类型定义
├── README.md						// 项目介绍文档
└── test
    ├── aggregate_test.go			// 聚合操作单元测试与用法示例
    ├── bench_test.go				// golang基准测试
    ├── chain_test.go				// 单元测试与用法示例
    └── conn_test.go				// 测试用例的初始化封装
```

## 如何贡献

Options 1: Fork仓库，提交后发起`pull request`

Options 2: 直接提交issue

## 版本规划

| 完成状况 | 计划内容 |
| --- | --- |
| DONE | mongodb基本操作(curd)、易用的聚合、分页操作 |
| DONE | README.md与聚合操作文档增加英文文档 |
| DONE | mongodb数据库管理命令支持 |
| TODO | 增加命令行工具支持(创建索引等), 使用Cobra |
| TODO | mongodb原生语句打印支持 |
| TODO | 事务支持 |

more and more...

## 捐赠

star一下即可~

## Star图表

[![Stargazers over time](https://starchart.cc/here-Leslie-Lau/mongo-plus.svg)](https://starchart.cc/here-Leslie-Lau/mongo-plus)

## 许可证

本项目基于MIT许可证发布，详情请参见[LICENSE](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/LICENSE)
