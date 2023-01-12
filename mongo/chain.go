// 调用链逻辑封装

package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Chain struct {
	ctx  context.Context
	coll *mongo.Collection
	// 条件暂存区
	condStorage map[string]interface{}
	// FindOptions条件暂存区
	findOpt *options.FindOptions
	// FindOneOptions条件暂存区
	findOneOpt *options.FindOneOptions
}

// 对chain的字段进行初始化赋值
func (ch *Chain) init() {
	if ch.condStorage == nil {
		ch.condStorage = make(map[string]interface{})
	}
	if ch.findOneOpt == nil {
		ch.findOneOpt = options.FindOne()
	}
	if ch.findOpt == nil {
		ch.findOpt = options.Find()
	}
	if ch.ctx == nil {
		ch.ctx = context.Background()
	}
}

// FindOne 查询单个文档
func (ch *Chain) FindOne(des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.condStorage)
	return ch.coll.FindOne(ch.ctx, f).Decode(des)
}

// Find 查询多个文档
func (ch *Chain) Find(des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.condStorage)
	cur, err := ch.coll.Find(ch.ctx, f, ch.findOpt)
	if err != nil {
		return err
	}

	if err := cur.All(ch.ctx, des); err != nil {
		return err
	}

	return nil
}

// InsertOne 插入一条文档
func (ch *Chain) InsertOne(doc interface{}) error {
	_, err := ch.coll.InsertOne(ch.ctx, doc)
	return err
}

// InsertMany 插入多条文档, 需要interface{}数组
func (ch *Chain) InsertMany(docs []interface{}) error {
	_, err := ch.coll.InsertMany(ch.ctx, docs)
	return err
}

// UpdateOne 根据chain的条件更新指定的一条文档, updateMap为更新的内容
func (ch *Chain) UpdateOne(ctx context.Context, updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{bson.E{Key: "$set", Value: updateMap}}

	_, err := ch.coll.UpdateOne(ctx, f, updateContent)
	return err
}

// Update 根据chain的条件更新指定的文档, updateMap为更新的内容
func (ch *Chain) Update(ctx context.Context, updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{bson.E{Key: "$set", Value: updateMap}}

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

// Paginate 分页查询, f: 分页相关参数, 方法调用结束后会将总条数/总页数放入f内 des: 查询结果集
func (ch *Chain) Paginate(f *PageFilter, des interface{}) (err error) {
	// 计算符合条件的总条数
	f.TotalCount, err = ch.Count(ch.ctx)
	if err != nil {
		return errors.Wrapf(err, "Paginate Chain Count fail")
	}
	if f.PageSize > 0 && f.PageNum > 0 {
		f.TotalPage = f.TotalCount / f.PageSize
		ch.Skip((f.PageNum - 1) * f.PageSize).Limit(f.PageSize)
	}

	if err := ch.Find(des); err != nil {
		return errors.Wrapf(err, "Paginate Chain Find des fail")
	}

	return nil
}
