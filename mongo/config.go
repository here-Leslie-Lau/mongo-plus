package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	// 最大连接数
	MaxPoolSize uint64
	// mongodb用户名
	Username string
	// mongodb密码
	Password string
	// 要连接的数据库名
	Database string
	// mongodb连接的url
	Addr []string
}

func (cfg *config) getOption() *options.ClientOptions {
	opt := options.Client()
	opt.Hosts = cfg.Addr
	opt.Auth = &options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	opt.SetMaxPoolSize(cfg.MaxPoolSize)

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

func WithAddr(addrs ...string) Option {
	return func(cfg *config) {
		cfg.Addr = addrs
	}
}
