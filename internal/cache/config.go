package cache

type Config struct {
	Host     string `json:"host" koanf:"host"`
	Password string `json:"password" koanf:"password"`
}
