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
	return ch.coll.FindOne(ctx, f).Decode(des)
}

// Find 查询多个文档
func (ch *Chain) Find(ctx context.Context, des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.findStorage)
	cur, err := ch.coll.Find(ctx, f)
	if err != nil {
		return err
	}

	if err := cur.All(ctx, des); err != nil {
		return err
	}

	return nil
}

// InsertOne 插入一条文档
func (ch *Chain) InsertOne(ctx context.Context, doc interface{}) error {
	_, err := ch.coll.InsertOne(ctx, doc)
	return err
}

// InsertMany 插入多条文档, 需要interface{}数组
func (ch *Chain) InsertMany(ctx context.Context, docs []interface{}) error {
	_, err := ch.coll.InsertMany(ctx, docs)
	return err
}
