// $group相关逻辑

package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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

// GetAvgStage 获取$avg的stage
// calledFiled: 计算出平均值之后的字段命名, filed: 要计算平均值的字段
// 该stage为group的子stage, 故一般与GetGroupStage组合使用, 作为subStages传入

// such as: ch.GetGroupStage("name", ch.GetAvgStage("avg", "age")), 代表根据name字段分组, 计算每个分组age的平均值, 最终每个组的平均值命名为avg
func (ch *Chain) GetAvgStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	filed = fmt.Sprintf("$%s", filed)
	d[0].Value = bson.D{{Key: AggregateOpeAvg.String(), Value: filed}}
	return d
}

// GetSumStage 获取$sum的stage
// calledFiled: 计算出总和之后的字段命名, filed: 要计算总和的字段
// 该stage为group的子stage, 故一般与GetGroupStage组合使用, 作为subStages传入

// such as: ch.GetGroupStage("name", ch.GetSumStage("total_sum", "age")), 代表根据name字段分组, 计算每个组age的总和, 最终每个组的总和命名为total_sum
func (ch *Chain) GetSumStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	filed = fmt.Sprintf("$%s", filed)
	d[0].Value = bson.D{{Key: AggregateOpeSum.String(), Value: filed}}
	return d
}

func (ch *Chain) GetFirstStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	d[0].Value = bson.D{{Key: AggregateOpeFirst.String(), Value: fmt.Sprintf("$%s", filed)}}
	return d
}

func (ch *Chain) GetLastStage(calledFiled, filed string) bson.D {
	d := bson.D{{Key: calledFiled}}
	d[0].Value = bson.D{{Key: AggregateOpeLast.String(), Value: fmt.Sprintf("$%s", filed)}}
	return d
}
