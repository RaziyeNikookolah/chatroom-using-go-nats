package config

type Config struct {
	AUTH_DB DBConfig     `json:"auth_db"  yaml:"auth_db"`
	Logger  LoggerConfig `json:"logger"  yaml:"logger"`
	Server  ServerConfig `json:"server"  yaml:"server"`
}

type ServerConfig struct {
	// HttpPort            uint `json:"httpPort"  yaml:"httpPort"`
	// RateLimitMaxAttempt int  `json:"rate_limit_max_attempt"  yaml:"rate_limit_max_attempt"`
	// RatelimitTimePeriod int  `json:"ratelimit_time_period"  yaml:"ratelimit_time_period"`
}
type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoggerConfig struct {
	Level  string `json:"level"  yaml:"level"`
	Output string `json:"output"  yaml:"output"`
	Path   string `json:"path"  yaml:"path"`
}
