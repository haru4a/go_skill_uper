package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DBType   string `toml:"db_type"`
	DBPath   string `toml:"db_path"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		DBType:   "sqlite3",
		DBPath:   "./test.db",
	}
}
