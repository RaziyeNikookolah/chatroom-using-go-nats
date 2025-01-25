package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
	Nats   NatsConfig   `json:"nats"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	GrpcPort uint   `json:"grpcPort"`
	Secret   string `json:"secret"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
type NatsConfig struct {
	Host    string `json:"host"`
	Port    uint   `json:"port"`
	Subject string `json:"subject"`
}
