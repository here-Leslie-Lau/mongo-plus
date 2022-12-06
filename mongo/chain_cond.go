// chain条件拼接相关逻辑

package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Where 单个查询条件拼接
func (ch *Chain) Where(key string, val interface{}) *Chain {
	ch.init()
	ch.condStorage[key] = val
	return ch
}

// Filter 多个查询条件
func (ch *Chain) Filter(filter map[string]interface{}) *Chain {
	ch.init()
	for k, v := range filter {
		ch.condStorage[k] = v
	}
	return ch
}

type Comparison string

func (c Comparison) String() string {
	return string(c)
}

const (
	// ComparisonGt 大于比较符
	ComparisonGt Comparison = "$gt"
	// ComparisonGte 大于等于比较符
	ComparisonGte Comparison = "$gte"
	// ComparisonLt 小于比较符
	ComparisonLt Comparison = "$lt"
	// ComparisonLte 小于等于比较符号
	ComparisonLte Comparison = "$lte"
)

// Comparison 比较运算封装, field: 字段名, symbol: 比较符号, val: 比较值
// such as: age ComparisonGt 1, 筛选年龄大于1的
func (ch *Chain) Comparison(field string, symbol Comparison, val int64) *Chain {
	ch.init()
	cond := bson.D{{symbol.String(), val}}
	s, exist := ch.condStorage[field]
	if !exist {
		// 该字段首次加入条件
		ch.condStorage[field] = cond
	} else {
		// 在原有的条件上追加
		d, ok := s.(bson.D)
		if !ok {
			// 原有的条件直接为等号运算，无需追加了
			return ch
		}
		d = append(d, cond[0])
		ch.condStorage[field] = d
	}
	return ch
}

// Gt "大于"运算的条件拼接, field: 字段名, val: 比较值
// such as: ch.Gt("age", 18).Find(context.Background(), &des), 找出年龄大于18岁的
func (ch *Chain) Gt(field string, val int64) *Chain {
	return ch.Comparison(field, ComparisonGt, val)
}

// Gte "大于等于"运算的条件拼接, field: 字段名, val: 比较值
func (ch *Chain) Gte(field string, val int64) *Chain {
	return ch.Comparison(field, ComparisonGte, val)
}

// Lt "小于"运算的条件拼接, field: 字段名, val: 比较值
func (ch *Chain) Lt(field string, val int64) *Chain {
	return ch.Comparison(field, ComparisonLt, val)
}

// Lte "小于等于"运算的条件拼接, field: 字段名, val: 比较值
func (ch *Chain) Lte(field string, val int64) *Chain {
	return ch.Comparison(field, ComparisonLte, val)
}
