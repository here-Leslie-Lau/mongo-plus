// bench相关测试

package test

import (
	"context"
	"testing"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
)

func BenchmarkFindOne(b *testing.B) {
	conn, f := newConn()
	defer f()

	res := &result{}
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).WithCtx(ctx).Where("name", "leslie").FindOne(res)
	}
}

func BenchmarkFind(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).WithCtx(ctx).Find(&list)
	}
}

func BenchmarkInString(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).WithCtx(ctx).InString("name", []string{"leslie"}).Find(&list)
	}
}

func BenchmarkInInt64(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).WithCtx(ctx).InInt64("value", []int64{100}).Find(&list)
	}
}

func BenchmarkSort(b *testing.B) {
	conn, cancel := newConn()
	defer cancel()

	var list []*result
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).WithCtx(ctx).Sort(mongo.SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&list)
	}

}
