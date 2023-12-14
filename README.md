# mongo-plus  [![](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://img.shields.io/badge/build-passing-brightgreen.svg) [![](https://img.shields.io/badge/version-v1.0-orange.svg)](https://img.shields.io/badge/version-v1.0-orange.svg) [![](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)](https://img.shields.io/badge/golang-%3E%3D%201.18-red.svg)

English | [中文](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/README_cn.md)

![](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/mongo-plus.png)

A secondary encapsulation based on the official MongoDB Go driver.

## Feature

- Call chain operations, freely combine conditions
- User-friendly API
- Support for Context
- Out-of-the-box functionality
- Pagination query support
- [Easy aggregation support](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/aggregate.md)
- [Database Management Command Operations](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/database.md)
- monitor support
- Continuously being updated

## Quick Start

```shell
go get -u github.com/here-Leslie-Lau/mongo-plus
```

**Create MongoDB Connection**

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

**Add monitor support**

```go
monitor := &event.CommandMonitor{
	Started: func(_ context.Context, evt *event.CommandStartedEvent) {
        fmt.Println("start")
	},
	Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
        fmt.Println("success")
	},
}
opts = append(opts, mongo.WithMonitor(monitor))
conn, f := mongo.NewConn(opts...)
defer f()
```

_The current monitor support includes `*event.CommandMonitor, *event.PoolMonitor, *event.ServerMonitor`. For more use cases, refer to test/monitor_test.go._

**Get Collection Object**

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

**Context Support**

```go
coll = coll.WithCtx(ctx)
```

**Insert Document**

```go
coll.InsertOne(document)
coll.InsertMany(documents)
```

**Query Documents**

```go
// Query a Single Document with the name "leslie"
coll.Where("name", "leslie").FindOne(&document)
// Query Documents with the name "leslie"
coll.Where("name", "leslie").Find(&documents)
// Multi-Condition Query
coll.Filter(map[string]interface{}{"name": "leslie", "age": 18}).FindOne(&document)
```

**Query the Number of Documents that Meet the Criteria**

```go
// Count of Documents with the name "leslie"
cnt, err := coll.Where("name", "leslie").Count()
```

**Sorting**

```go
// Query Ascending by the "value" Field
coll.Sort(mongo.SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&documents)
```

**Pagination**

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

**Logical Operations**

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

**Specify the Fields to Query**

```go
// Query Results Only Assign to the "name" Field, After Calling This Method, the "_id" Field Is Not Assigned by Default
coll.Projection("name").Find(&documents)
```

**Update or Insert a Record**

```go
// Update the age Field to 18
content := map[string]interface{}{"age": 18}
// Default Values to Insert if the Filter Criteria Do Not Exist
default := map[string]interface{}{"name": "leslie"}

conn.Where("name", "leslie").UpsertOne(content, default)
// Desired Outcome: If a document with name "leslie" exists, update the age to 18; otherwise, insert a document {"name": "leslie", "age": 18}.
```

**OR Query (Logical OR Operation Query)**

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

[Database Management Command Operations](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/docs/database.md)

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

## Original Intention

*When working or writing small pieces on my own, I often feel that using the official [Go driver](https://www.mongodb.com/docs/drivers/go/current/) provided by MongoDB is somewhat inconvenient.*

I've summarized personally, and there are several points as follows:

- When performing mongodb operations, you need to prepare various Option objects from the official driver and pass them all at once. Maybe it's the call chain approach of gorm that resonates with me 😄, so I also want to encapsulate it in a similar way.
- The official driver doesn't provide a very convenient way for pagination, _(for example: getting the total page count/total number of items based on the page number/page size passed from the frontend or client),_ so each time it needs to be encapsulated again.
- I believe a library should try to shield the details as much as possible, and users shouldn't have to focus on the underlying implementation, making it ready to use. _(For example, developers shouldn't need to understand BSON, various operators like $gt, sharding, etc.)_

## Project Structure

```shell
.
├── docs							// project documentation
├── go.mod
├── go.sum
├── LICENSE
├── makefile						// some initialization tools
├── mongo							// core logic package
│   ├── aggregate.go				// mongodb aggregation operation logic
│   ├── aggregate_group.go			// group logic in aggregation operation
│   ├── chain_cond.go				// call chain condition concatenation logic
│   ├── chain.go					// core struct definition for chain and encapsulation of mongoDB operation methods
│   ├── collection.go				// collection interface definition
│   ├── config.go					// mongoDB connection configuration definition
│   ├── conn.go						// connection retrieval logic and initialization methods
│   ├── database.go                 // database management methods
│   ├── paginate.go					// pagination logic encapsulation
│   └── type.go						// mongoDB type definitions
├── README.md
└── test
    ├── aggregate_test.go
    ├── bench_test.go
    ├── chain_test.go
    └── conn_test.go
```

## How to Contribute

Options 1: Fork the repository, make your changes, and then initiate a "pull request" to submit your changes.

Options 2: Submit an Issue Directly

## Version Planning

| Status | Content |
| --- | --- |
| DONE | Basic MongoDB Operations (CRUD), User-Friendly Aggregation, Pagination Operations |
| DONE | Add english documentation |
| DONE | Support for MongoDB database management commands |
| TODO | Adding command-line tool support (creating indexes, etc.), use Cobra |
| TODO | Printing support for MongoDB native statements |
| TODO | Transaction Support |

more and more...

## Donation

Just give it a star~

## Star Charts

[![Stargazers over time](https://starchart.cc/here-Leslie-Lau/mongo-plus.svg)](https://starchart.cc/here-Leslie-Lau/mongo-plus)

## License

Mongo-plus is licensed under the MIT License. See the [LICENSE](https://github.com/here-Leslie-Lau/mongo-plus/blob/master/LICENSE)
