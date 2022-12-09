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
