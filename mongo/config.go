package mongo

import (
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	// 最大连接数
	MaxPoolSize uint64
	// 最小连接数
	MinPoolSize uint64
	// mongodb用户名
	Username string
	// mongodb密码
	Password string
	// 要连接的数据库名
	Database string
	// mongodb连接的url
	Addr []string
	// 官方事件监听器
	PoolMonitor *event.PoolMonitor
}

func (cfg *config) getOption() *options.ClientOptions {
	opt := options.Client()
	opt.Hosts = cfg.Addr
	opt.Auth = &options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	if cfg.MaxPoolSize > 0 && cfg.MaxPoolSize > cfg.MinPoolSize {
		opt.SetMaxPoolSize(cfg.MaxPoolSize)
	}
	if cfg.MinPoolSize > 0 {
		opt.SetMinPoolSize(cfg.MinPoolSize)
	}

	// regist event
	if cfg.PoolMonitor != nil {
		opt.SetPoolMonitor(cfg.PoolMonitor)
	}
	return opt
}

type Option func(*config)

func WithUsername(username string) Option {
	return func(cfg *config) {
		cfg.Username = username
	}
}

func WithPassword(passwd string) Option {
	return func(cfg *config) {
		cfg.Password = passwd
	}
}

func WithDatabase(dbName string) Option {
	return func(cfg *config) {
		cfg.Database = dbName
	}
}

func WithMaxPoolSize(maxPoolSize uint64) Option {
	return func(cfg *config) {
		cfg.MaxPoolSize = maxPoolSize
	}
}

func WithMinPoolSize(minPoolSize uint64) Option {
	return func(cfg *config) {
		cfg.MinPoolSize = minPoolSize
	}
}

func WithAddr(addrs ...string) Option {
	return func(cfg *config) {
		cfg.Addr = addrs
	}
}

func WithPoolMonitor(poolMonitor *event.PoolMonitor) Option {
	return func(cfg *config) {
		cfg.PoolMonitor = poolMonitor
	}
}
