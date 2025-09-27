package models

type Config struct {
	AppParams      AppParams      `json:"app_params" env:"APP_PARAMS"`
	PostgresParams PostgresParams `json:"postgres_params"`
	RedisParams    RedisParams    `json:"redis_params"`
	AuthParams     AuthParams     `json:"auth_params"`
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

type AuthParams struct {
	AccessTokenTllMinutes int `json:"access_token_tll_minutes"`
	RefreshTokenTllDays   int `json:"refresh_token_tll_days"`
}
