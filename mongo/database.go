package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetDB 获取go driver的database对象
func (c *Conn) GetDB() *mongo.Database {
	return c.db
}

// RunCommand 执行mongodb数据库管理命令
// cmd: 命令集合, bson.D类型
// des: 命令执行结果, 结构体指针
func (c *Conn) RunCommand(ctx context.Context, cmd bson.D, des interface{}) error {
	return c.GetDB().RunCommand(ctx, cmd).Decode(des)
}
