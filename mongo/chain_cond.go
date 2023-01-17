// chain条件拼接相关逻辑

package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (ch *Chain) WithCtx(ctx context.Context) *Chain {
	ch.ctx = ctx
	return ch
}

// Where 单个查询条件拼接
func (ch *Chain) Where(key string, val interface{}) *Chain {
	ch.condStorage[key] = val
	return ch
}

// Filter 多个查询条件
func (ch *Chain) Filter(filter map[string]interface{}) *Chain {
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
	// ComparisonIn 范围查询符号(匹配项)
	ComparisonIn Comparison = "$in"
	// ComparisonNotIn 范围查询符号(排除匹配项)
	ComparisonNotIn Comparison = "$nin"
	// ComparisonEq 等于比较符
	ComparisonEq Comparison = "$eq"
	// ComparisonNotEq 不等于比较符
	ComparisonNotEq Comparison = "$ne"
)

// Comparison 比较运算封装, field: 字段名, symbol: 比较符号, val: 比较值
// such as: Comparison(age, ComparisonGt, 1), 筛选年龄大于1的
func (ch *Chain) Comparison(field string, symbol Comparison, val interface{}) *Chain {
	cond := bson.D{bson.E{Key: symbol.String(), Value: val}}
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
// such as: ch.Gt("age", 18).Find(&des), 找出年龄大于18岁的
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

// In 匹配数组中指定的任何值, field: 字段名, arrays: 数组
// such as: ch.In("age", []interface{}{18, 19}).Find(&des), 找年龄为18和19岁的
func (ch *Chain) In(field string, arrays []interface{}) *Chain {
	return ch.Comparison(field, ComparisonIn, arrays)
}

// InInt64 匹配数组中指定的任何值, 数组类型为int64(语法糖), field: 字段名, arrays: 数组
func (ch *Chain) InInt64(field string, arrays []int64) *Chain {
	return ch.Comparison(field, ComparisonIn, arrays)
}

// InString 匹配数组中指定的任何值, 数组类型为string(语法糖), field: 字段名, arrays: 数组
func (ch *Chain) InString(field string, arrays []string) *Chain {
	return ch.Comparison(field, ComparisonIn, arrays)
}

// NotIn 不匹配数组中指定的任何值, field: 字段名, arrays: 数组
func (ch *Chain) NotIn(field string, arrays []interface{}) *Chain {
	return ch.Comparison(field, ComparisonNotIn, arrays)
}

// NotInInt64 不匹配数组中指定的任何值, 数组类型为int64(语法糖), field: 字段名, arrays: 数组
func (ch *Chain) NotInInt64(field string, arrays []int64) *Chain {
	return ch.Comparison(field, ComparisonNotIn, arrays)
}

// NotInString 不匹配数组中指定的任何值, 数组类型为string(语法糖), field: 字段名, arrays: 数组
func (ch *Chain) NotInString(field string, arrays []string) *Chain {
	return ch.Comparison(field, ComparisonNotIn, arrays)
}

// Eq "等于"运算的条件拼接, field: 字段名, val: 比较值
func (ch *Chain) Eq(field string, val interface{}) *Chain {
	return ch.Comparison(field, ComparisonEq, val)
}

// NotEq "不等于"运算的条件拼接, field: 字段名, val: 比较值
func (ch *Chain) NotEq(field string, val interface{}) *Chain {
	return ch.Comparison(field, ComparisonNotEq, val)
}

type Element string

func (e Element) String() string {
	return string(e)
}

func (e Element) comparison() Comparison {
	return Comparison(e)
}

const (
	// ElementExists 匹配具有指定字段的文档
	ElementExists Element = "$exists"
	// ElementType 如果字段属于指定类型，则选择文档
	ElementType Element = "$type"
)

// Exists 匹配具有指定字段的文档， field: 字段名, exist: 布尔值
// such as: ch.Exists("name", true).Find(&des), 找出存在name字段的数据
func (ch *Chain) Exists(field string, exist bool) *Chain {
	return ch.Comparison(field, ElementExists.comparison(), exist)
}

// Type 如果字段属于指定类型，则选择文档
func (ch *Chain) Type(field string, typ MongodbType) *Chain {
	return ch.Comparison(field, ElementType.comparison(), typ.int())
}

// Limit 指定查询返回的文档数
func (ch *Chain) Limit(limit int64) *Chain {
	ch.findOpt.SetLimit(limit)
	return ch
}

// Sort 根据条件进行排序
// such as: ch.Sort(SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&list), 根据value升序查询
func (ch *Chain) Sort(rules ...SortRule) *Chain {
	sRules := bson.D{}
	for _, rule := range rules {
		sRules = append(sRules, bson.E{Key: rule.Field, Value: rule.Typ})
	}
	// 双写
	ch.findOpt.SetSort(sRules)
	ch.findOneOpt.SetSort(sRules)
	return ch
}

// Skip 跳过指定条数查询
// such as: ch.Skip(1) 跳过第一条查询
func (ch *Chain) Skip(skip int64) *Chain {
	ch.findOpt.SetSkip(skip)
	ch.findOneOpt.SetSkip(skip)
	return ch
}

// Projection 指定要返回的字段
// such as: ch.Projection("name", "age") 只返回name和age字段
func (ch *Chain) Projection(fileds ...string) *Chain {
	// 默认不显示_id字段
	cond := bson.D{bson.E{Key: "_id", Value: 0}}

	for _, filed := range fileds {
		// 显示需要展示的字段
		cond = append(cond, bson.E{
			Key:   filed,
			Value: 1,
		})
	}

	// 双写
	ch.findOpt.SetProjection(cond)
	ch.findOneOpt.SetProjection(cond)
	return ch
}
