package config

type Config struct {
	Server ServerConfig `json:"server"`
	Nats   NatsConfig   `json:"nats"`
}

type ServerConfig struct {
	GrpcPort uint `json:"grpcPort"`
}

type NatsConfig struct {
	Host      string `json:"host"`
	Port      uint   `json:"port"`
	Subject   string `json:"subject"`
	ServerURL string `json:"server-url"`
}
