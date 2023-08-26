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

// IsMaster 用于检查服务器的状态，特别是用于确定服务器是否是主节点（Primary）
// 是否是一个副本集（Replica Set），以及其他关于服务器的基本信息
// 这个命令可以帮助你了解当前 MongoDB 部署的状态和拓扑
//
// des: 命令执行结果, 结构体指针
func (c *Conn) IsMaster(ctx context.Context, des interface{}) error {
	cmd := bson.D{{Key: "isMaster", Value: 1}}
	return c.RunCommand(ctx, cmd, des)
}

// Ping 用于检查服务器是否可用
func (c *Conn) Ping(ctx context.Context) error {
	res := make(map[string]interface{})
	cmd := bson.D{{Key: "ping", Value: 1}}
	return c.RunCommand(ctx, cmd, res)
}

// DbStats 用于获取数据库的统计信息
// 该命令返回的信息包括数据库的大小，对象的数量，索引的数量，以及其他统计信息
//
// des: 命令执行结果, 结构体指针
func (c *Conn) DbStats(ctx context.Context, des interface{}) error {
	cmd := bson.D{{Key: "dbStats", Value: 1}}
	return c.RunCommand(ctx, cmd, des)
}

// 需要admin数据库权限，该库权限过高，暂时取消
// func (c *Conn) ListDatabases(ctx context.Context, des interface{}) error {
//	cmd := bson.D{{Key: "listDatabases", Value: 1}}
//	return c.RunCommand(ctx, cmd, des)
// }

// ServerStatus 用于获取服务器的状态信息
// 该命令返回的信息包括服务器的版本，操作系统，内存使用情况，连接数，以及其他统计信息
//
// des: 命令执行结果, 结构体指针
func (c *Conn) ServerStatus(ctx context.Context, des interface{}) error {
	cmd := bson.D{{Key: "serverStatus", Value: 1}}
	return c.RunCommand(ctx, cmd, des)
}
