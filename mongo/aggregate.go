// 聚合aggregation相关逻辑

package mongo

import (
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
