// 调用链逻辑封装

package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Chain struct {
	coll *mongo.Collection
	// 条件暂存区
	condStorage map[string]interface{}
}

func (ch *Chain) init() {
	if ch.condStorage == nil {
		ch.condStorage = make(map[string]interface{})
	}
}

// FindOne 查询单个文档
func (ch *Chain) FindOne(ctx context.Context, des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.condStorage)
	return ch.coll.FindOne(ctx, f).Decode(des)
}

// Find 查询多个文档
func (ch *Chain) Find(ctx context.Context, des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.condStorage)
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

// UpdateOne 根据chain的条件更新指定的一条文档, updateMap为更新的内容
func (ch *Chain) UpdateOne(ctx context.Context, updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{{"$set", updateMap}}

	_, err := ch.coll.UpdateOne(ctx, f, updateContent)
	return err
}

// Update 根据chain的条件更新指定的文档, updateMap为更新的内容
func (ch *Chain) Update(ctx context.Context, updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{{"$set", updateMap}}

	_, err := ch.coll.UpdateMany(ctx, f, updateContent)
	return err
}

// DeleteOne 根据chain的条件删除一条文档
func (ch *Chain) DeleteOne(ctx context.Context) error {
	f := bson.M(ch.condStorage)
	_, err := ch.coll.DeleteOne(ctx, f)
	return err
}

// Delete 根据chain的条件删除指定的文档
func (ch *Chain) Delete(ctx context.Context) error {
	f := bson.M(ch.condStorage)
	_, err := ch.coll.DeleteMany(ctx, f)
	return err
}

// Count 根据chain内的条件查询满足条件的文档记录数
func (ch *Chain) Count(ctx context.Context) (int64, error) {
	f := bson.M(ch.condStorage)
	return ch.coll.CountDocuments(ctx, f)
}
