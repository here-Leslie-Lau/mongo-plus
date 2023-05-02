// 聚合aggregation相关逻辑

package mongo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Aggregate 聚合操作
// des: 聚合操作后的结果集(指针), stages: 对应的stage集合, 按顺序排列
// stages可以调用chain.GetXXXStage获取, 如果不满足则自定义bson.D传入
func (ch *Chain) Aggregate(des interface{}, stages ...bson.D) error {
	if len(stages) == 0 {
		return errors.New("invalid params")
	}

	cur, err := ch.coll.Aggregate(ch.ctx, mongo.Pipeline(stages))
	if err != nil {
		return errors.Wrap(err, "Chain Aggregate fail")
	}
	if err := cur.All(ch.ctx, des); err != nil {
		return errors.Wrap(err, "Chain Aggregate cursor All fail")
	}

	return nil
}

// GetMatchStage 获取$match的stage
// filed: 匹配的字段, val: 具体匹配值
func (ch *Chain) GetMatchStage(filed, val string) bson.D {
	return bson.D{
		{
			Key:   AggregateOpeMatch.String(),
			Value: bson.D{{Key: filed, Value: val}},
		},
	}
}

// GetSortStage 获取$sort的stage
// rules: 具体的排序规则集合, 可参考mongo.SortRule
func (ch *Chain) GetSortStage(rules ...SortRule) bson.D {
	d := bson.D{}
	for _, rule := range rules {
		d = append(d, bson.E{Key: rule.Field, Value: rule.Typ})
	}
	return bson.D{
		{
			Key:   AggregateOpeSort.String(),
			Value: d,
		},
	}
}

// GetLimitStage 获取$limit的stage
// num: 限制的文档数
func (ch *Chain) GetLimitStage(num int64) bson.D {
	return bson.D{
		{
			Key:   AggregateopeLimit.String(),
			Value: num,
		},
	}
}

// GetSkipStage 获取$skip的stage
// num: 要跳过的文档数
func (ch *Chain) GetSkipStage(num int64) bson.D {
	return bson.D{
		{
			Key:   AggregateOpeSkip.String(),
			Value: num,
		},
	}
}

// GetUnsetStage 获取$unset的stage
// fileds: 要忽略的字段
func (ch *Chain) GetUnsetStage(fileds ...string) bson.D {
	return bson.D{
		{
			Key:   AggregateOpeUnset.String(),
			Value: fileds,
		},
	}
}

func (ch *Chain) GetProjectStage(fileds ...string) bson.D {
	m := bson.M{}
	for _, filed := range fileds {
		m[filed] = true
	}
	return bson.D{
		{
			Key:   AggregateOpeProject.String(),
			Value: m,
		},
	}
}
