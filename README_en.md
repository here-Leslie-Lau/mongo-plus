# mongo-plus  [![](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://img.shields.io/badge/build-passing-brightgreen.svg) [![](https://img.shields.io/badge/version-v1.0-orange.svg)](https://img.shields.io/badge/version-v1.0-orange.svg) [![](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)

![](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/mongo-plus.png)

A secondary encapsulation based on the official MongoDB Go driver.

## Feature

- Call chain operations, freely combine conditions
- User-friendly API
- Support for Context
- Out-of-the-box functionality
- Pagination query support
- [Easy aggregation support](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate.md)
- Continuously being updated

## Quick Start

```shell
go get -u github.com/here-Leslie-Lau/mongo-plus
```

Create MongoDB Connection

```go
opts := []mongo.Option{
	// Database to Connect
	mongo.WithDatabase("test"),
	// Maximum Connection Pool Size
	mongo.WithMaxPoolSize(10),
	// Username
	mongo.WithUsername("your username"),
	// Password
	mongo.WithPassword("your password"),
	// Connection URL
	mongo.WithAddr("localhost:27017"),
}
conn, f := mongo.NewConn(opts...)
defer f()
```

Get Collection Object

```go
type Demo struct{}

// Collection Implement the mongo.Collection interface, returning the name of the collection to operate on.
func (d *Demo) Collection() string {
	return "demo"
}

// Get the collection object within the method.
demo := &Demo{}
coll := conn.Collection(demo)
```

Context Support

```go
coll = coll.WithCtx(ctx)
```

Insert Document

```go
coll.InsertOne(document)
coll.InsertMany(documents)
```

Query Documents

```go
// Query a Single Document with the name "leslie"
coll.Where("name", "leslie").FindOne(&document)
// Query Documents with the name "leslie"
coll.Where("name", "leslie").Find(&documents)
// Multi-Condition Query
coll.Filter(map[string]interface{}{"name": "leslie", "age": 18}).FindOne(&document)
```

Query the Number of Documents that Meet the Criteria

```go
// Count of Documents with the name "leslie"
cnt, err := coll.Where("name", "leslie").Count()
```

Sorting

```go
// Query Ascending by the "value" Field
coll.Sort(mongo.SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&documents)
```

Pagination

```go
f := &mongo.PageFilter{
	// Current Page of the Query
	PageNum:  1,
	// Number of Items per Page
	PageSize: 2,
}

// Place matching documents into the "documents" based on the criteria and place the total count and total pages into the "f".
coll.Paginate(f, &documents)
```

Logical Operations

```go
// Find a Single Record with age Greater Than 18
coll.Gt("age", 18).FindOne(&document)
// Find a Single Record with age Less Than 18
coll.Lt("age", 18).FindOne(&document)
// Find a Single Record with age Greater Than or Equal to 18
coll.Gte("age", 18).FindOne(&document)
// Find a Single Record where age is Not Equal to 100
coll.NotEq("age", 100).FindOne(&document)
// ...other methods can be referenced in mongo/chain_cond.go
```

Specify the Fields to Query

```go
// Query Results Only Assign to the "name" Field, After Calling This Method, the "_id" Field Is Not Assigned by Default
coll.Projection("name").Find(&documents)
```

Update or Insert a Record

```go
// Update the age Field to 18
content := map[string]interface{}{"age": 18}
// Default Values to Insert if the Filter Criteria Do Not Exist
default := map[string]interface{}{"name": "leslie"}

conn.Where("name", "leslie").UpsertOne(content, default)
// Desired Outcome: If a document with name "leslie" exists, update the age to 18; otherwise, insert a document {"name": "leslie", "age": 18}.
```

OR Query (Logical OR Operation Query)

```go
// Single Condition
orMap := map[string]interface{}{"age": 18, "name": "leslie"}
// Query Documents where name is "leslie" or age is 18
conn.Or(orMap).Find(&documents)

// Multiple Conditions
orMap1 := map[string]interface{}{"name": "leslie", "age": 22}
orMap2 := map[string]interface{}{"name": "skyle", "age": 78}
// Query Documents where name is "leslie" and age is 22, or name is "skyle" and age is 78.
conn.Ors(orMap1, orMap2).Find(&documents)
```

[Aggregate Operation](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate.md)

_Additional documentation can be found in the provided test/chain_test.go file for more detailed usage examples._

## Benchmark

machine: Macbook Pro M1 <br/>
memory: 8G

```shell
make benchmark
```

Output Results (Other Methods Are Supplemented):

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
