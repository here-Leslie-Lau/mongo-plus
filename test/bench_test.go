// bench相关测试

package test

import (
	"context"
	"testing"
)

func BenchmarkFindOne(b *testing.B) {
	conn, f := newConn()
	defer f()

	res := &result{}
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).Where("name", "leslie").FindOne(ctx, res)
	}
}

func BenchmarkFind(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).Find(ctx, &list)
	}
}

func BenchmarkInString(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).InString("name", []string{"leslie"}).Find(ctx, &list)
	}
}

func BenchmarkInInt64(b *testing.B) {
	conn, f := newConn()
	defer f()

	var list []*result
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_ = conn.Collection(&demo{collName: "demo"}).InInt64("value", []int64{100}).Find(ctx, &list)
	}
}
