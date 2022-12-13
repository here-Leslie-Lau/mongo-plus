// 分页相关逻辑

package mongo

type PageFilter struct {
	// 页数
	PageNum int64
	// 每页大小
	PageSize int64
	// 总条数
	TotalCount int64
	// 总页数
	TotalPage int64
}
