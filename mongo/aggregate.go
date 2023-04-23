// 聚合aggregation相关逻辑

package mongo

import (
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

// GetGroupStage 获取$group的stage
// groupFiled: 要分组的字段名, subStages: 子stage, 如果需要则传入
// tips: 目前该库只提供少量stage支持, 也可以自定义bson传入
func (ch *Chain) GetGroupStage(groupFiled string, subStages ...bson.D) bson.D {
	valueBson := bson.D{{Key: "_id", Value: fmt.Sprintf("$%s", groupFiled)}}
	d := bson.D{{Key: AggregateOpeGroup.String()}}

	for _, stage := range subStages {
		valueBson = append(valueBson, stage...)
	}

	d[0].Value = valueBson
	return d
}

func (ch *Chain) GetAvgStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	filed = fmt.Sprintf("$%s", filed)
	d[0].Value = bson.D{{Key: AggregateOpeAvg.String(), Value: filed}}
	return d
}

func (ch *Chain) GetSumStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	filed = fmt.Sprintf("$%s", filed)
	d[0].Value = bson.D{{Key: AggregateOpeSum.String(), Value: filed}}
	return d
}
