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
