// 调用链逻辑封装

package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Chain struct {
	coll *mongo.Collection
	// 查询条件暂存区
	findStorage map[string]interface{}
}

func (ch *Chain) init() {
	if ch.findStorage == nil {
		ch.findStorage = make(map[string]interface{})
	}
}

// Where 单个查询条件拼接
func (ch *Chain) Where(key, val string) *Chain {
	ch.init()
	ch.findStorage[key] = val
	return ch
}

// Filter 多个查询条件
func (ch *Chain) Filter(filter map[string]interface{}) *Chain {
	ch.init()
	for k, v := range filter {
		ch.findStorage[k] = v
	}
	return ch
}

// FindOne 查询单个文档
func (ch *Chain) FindOne(ctx context.Context, des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.findStorage)
	if err := ch.coll.FindOne(ctx, f).Decode(des); err != nil {
		return err
	}
	return nil
}
