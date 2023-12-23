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
	PoolMonitor    *event.PoolMonitor
	CommandMonitor *event.CommandMonitor
	ServerMonitor  *event.ServerMonitor

	// 日志相关
	LogOpt *options.LoggerOptions
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
	if cfg.CommandMonitor != nil {
		opt.SetMonitor(cfg.CommandMonitor)
	}
	if cfg.ServerMonitor != nil {
		opt.SetServerMonitor(cfg.ServerMonitor)
	}

	// log
	if cfg.LogOpt != nil {
		opt.SetLoggerOptions(cfg.LogOpt)
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

// WithPoolMonitor set pool monitor into config
func WithPoolMonitor(poolMonitor *event.PoolMonitor) Option {
	return func(cfg *config) {
		cfg.PoolMonitor = poolMonitor
	}
}

// WithCommandMonitor set command monitor into config
func WithCommandMonitor(commandMonitor *event.CommandMonitor) Option {
	return func(cfg *config) {
		cfg.CommandMonitor = commandMonitor
	}
}

func WithServerMonitor(serverMonitor *event.ServerMonitor) Option {
	return func(cfg *config) {
		cfg.ServerMonitor = serverMonitor
	}
}

// WithMonitor set monitor into config
// monitor must be *event.PoolMonitor or *event.CommandMonitor
// if monitor isn't these type, panic
func WithMonitor(monitor interface{}) Option {
	switch v := monitor.(type) {
	case *event.PoolMonitor:
		return WithPoolMonitor(v)
	case *event.CommandMonitor:
		return WithCommandMonitor(v)
	case *event.ServerMonitor:
		return WithServerMonitor(v)
	default:
		panic("invalid monitor type")
	}
}

// WithLoggerOptions set logger options into config
func WithLoggerOptions(logOpt *options.LoggerOptions) Option {
	return func(cfg *config) {
		cfg.LogOpt = logOpt
	}
}
