package mongo

type Collection interface {
	// Collection 返回对应的集合名称
	Collection() string
}
