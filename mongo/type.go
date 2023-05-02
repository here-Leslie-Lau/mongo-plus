package mongo

// MongodbType mongodb数据类型定义
type MongodbType int

func (m MongodbType) int() int {
	return int(m)
}

const (
	Double MongodbType = iota + 1
	String
	Object
	Array
	BinaryData
	Undefined
	ObjectId
	Boolean
	Date
	Null
	RegularExpression
	JavaScript MongodbType = iota + 2
	Symbol
	Int32     MongodbType = 16
	Timestamp MongodbType = 17
	Int64     MongodbType = 18
)

// SortType 排序类型
type SortType int8

const (
	// SortTypeASC 排序类型: 正序
	SortTypeASC = 1
	// SortTypeDESC 排序类型: 倒序
	SortTypeDESC = -1
)

type SortRule struct {
	// 排序类型
	Typ SortType
	// 排序字段名
	Field string
}

// AggregateOpe 聚合操作表达式
type AggregateOpe string

func (a AggregateOpe) String() string {
	return string(a)
}

const (
	// AggregateOpeGroup 分组
	AggregateOpeGroup AggregateOpe = "$group"
	// AggregateOpeMatch 匹配
	AggregateOpeMatch AggregateOpe = "$match"
	// AggregateOpeSum 计算总和
	AggregateOpeSum AggregateOpe = "$sum"
	// AggregateOpeSort 排序
	AggregateOpeSort AggregateOpe = "$sort"
	// AggregateOpeAvg 计算平均值
	AggregateOpeAvg AggregateOpe = "$avg"
	// AggregateopeLimit 限制记录数
	AggregateopeLimit AggregateOpe = "$limit"
	// AggregateOpeSkip 跳过指定记录数
	AggregateOpeSkip AggregateOpe = "$skip"
	// AggregateOpeUnset 忽略相关字段
	AggregateOpeUnset AggregateOpe = "$unset"
	// AggregateOpeProject 只返回指定字段
	AggregateOpeProject AggregateOpe = "$project"
)
