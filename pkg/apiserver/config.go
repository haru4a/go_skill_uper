package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DBType   string `toml:"db_type"`
	DBPath   string `toml:"db_path"`
}

// NewConfig ...
func NewConfig(bindAddr string, logLevel string, dbFlag string) *Config {
	return &Config{
		BindAddr: bindAddr,
		LogLevel: logLevel,
		DBType:   "sqlite3",
		DBPath:   dbFlag,
	}
}
