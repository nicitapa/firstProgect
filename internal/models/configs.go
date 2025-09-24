package models

type Config struct {
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
	RedisParams    RedisParams    `json:"redis_params"`
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`
	PortRun    string `json:"port_run"`
	ServerUrl  string `json:"server_url"`
	ServerName string `json:"server_name"`
}

type PostgresParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
}
type RedisParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database int    `json:"database"`
}
