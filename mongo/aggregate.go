// 聚合aggregation相关逻辑

package mongo

import (
	"github.com/pkg/errors"
)

// PipelineStage 聚合管道的stage结构体
type PipelineStage struct {
	// Ope stage操作符
	Ope AggregateOpe
	// Filed 要操作的字段名
	Filed string
	// Value 要操作的具体值
	Value string

	// Child 子stage, 如果没有则不用传
	Child *PipelineStage
}

func (ch *Chain) Aggregate(des interface{}, stages ...PipelineStage) error {
	if len(stages) == 0 {
		return errors.New("invalid params")
	}

	return nil
}
