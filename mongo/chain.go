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
	ctx      context.Context
	collName string
	coll     *mongo.Collection
	// 条件暂存区
	condStorage map[string]interface{}
	// FindOptions条件暂存区
	findOpt *options.FindOptions
	// FindOneOptions条件暂存区
	findOneOpt *options.FindOneOptions
	// UpdateOptions条件暂存区
	updateOpt *options.UpdateOptions
	// native mongo statement
	statement *Statement
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
	// init chain statement
	ch.statement = newStatement(ch.collName)
}

// FindOne 查询单个文档
func (ch *Chain) FindOne(des interface{}) error {
	// map => bson.M{}
	f := bson.M(ch.condStorage)
	return ch.coll.FindOne(ch.ctx, f, ch.findOneOpt).Decode(des)
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
func (ch *Chain) UpdateOne(updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{bson.E{Key: "$set", Value: updateMap}}

	_, err := ch.coll.UpdateOne(ch.ctx, f, updateContent)
	return err
}

// Update 根据chain的条件更新指定的文档, updateMap为更新的内容
func (ch *Chain) Update(updateMap map[string]interface{}) error {
	f := bson.M(ch.condStorage)
	updateContent := bson.D{bson.E{Key: "$set", Value: updateMap}}

	_, err := ch.coll.UpdateMany(ch.ctx, f, updateContent)
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
func (ch *Chain) Count() (int64, error) {
	f := bson.M(ch.condStorage)
	return ch.coll.CountDocuments(ch.ctx, f)
}

// Paginate 分页查询, f: 分页相关参数, 方法调用结束后会将总条数/总页数放入f内 des: 查询结果集
func (ch *Chain) Paginate(f *PageFilter, des interface{}) (err error) {
	// 计算符合条件的总条数
	f.TotalCount, err = ch.Count()
	if err != nil {
		return errors.Wrapf(err, "Paginate Chain Count fail")
	}
	if f.PageSize > 0 && f.PageNum > 0 {
		f.TotalPage = f.TotalCount / f.PageSize
		if f.TotalPage == 0 {
			f.TotalPage += 1
		}
		ch.Skip((f.PageNum - 1) * f.PageSize).Limit(f.PageSize)
	}

	if err := ch.Find(des); err != nil {
		return errors.Wrapf(err, "Paginate Chain Find des fail")
	}

	return nil
}

// Upsert 根据chain内的查询条件暂存区,更新或新增一条记录
// content: 要更新的内容
// defaultContent: 如果没找到记录, 插入文档的字段默认值
// such as: ch.Where("name", "leslie").UpsertOne(map[string]interface{}{"age": 18}, map[string]interface{}{"name", "leslie"}), 找到name为leslie的文档, 并将age更新为18, 如果没找到则新增一条文档, 默认值为name-leslie, age-18
func (ch *Chain) UpsertOne(content, defaultContent map[string]interface{}) error {
	if len(ch.condStorage) == 0 {
		// 避免操作整个集合, 直接返回错误
		return errors.New("chain condStorage is zero, Upsert fail")
	}

	// map => bson.M
	f := bson.M(ch.condStorage)
	c := bson.D{
		// 要更新的内容
		{
			Key:   "$set",
			Value: bson.M(content),
		},

		// 如果找不到记录, 要插入新记录的默认值
		{
			Key:   "$setOnInsert",
			Value: bson.M(defaultContent),
		},
	}

	if ch.updateOpt == nil {
		ch.updateOpt = options.Update()
	}

	_, err := ch.coll.UpdateOne(ch.ctx, f, c, ch.updateOpt.SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "Upsert fail")
	}

	return nil
}

// IncOne 给一个文档字段增加或减少指定的数字
// such as: ch.Where("name", "leslie").IncOne("age", 1) 代表给name为leslie的单个文档, age加一
func (ch *Chain) IncOne(field string, incNums int64) error {
	f := bson.M(ch.condStorage)

	updateContent := bson.D{bson.E{Key: "$inc", Value: map[string]interface{}{field: incNums}}}
	_, err := ch.coll.UpdateOne(ch.ctx, f, updateContent)
	return err
}

// Inc 给多个文档字段增加或减少指定的数字
// such as: ch.Where("name", "leslie").IncOne("age", 1) 代表给name为leslie的所有文档, age加一
func (ch *Chain) Inc(field string, incNums int64) error {
	if len(ch.condStorage) == 0 {
		// 避免操作整个集合, 直接返回错误
		return errors.New("chain condStorage is zero, Inc fail")
	}
	f := bson.M(ch.condStorage)

	updateContent := bson.D{bson.E{Key: "$inc", Value: map[string]interface{}{field: incNums}}}
	_, err := ch.coll.UpdateMany(ch.ctx, f, updateContent)
	return err
}
