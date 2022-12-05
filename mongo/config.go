package mongo

type config struct {
	// 最大连接数
	MaxPoolSize uint64
	// mongodb用户名
	Username string
	// mongodb密码
	Password string
	// mongodb连接的url
	Addr []string
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
