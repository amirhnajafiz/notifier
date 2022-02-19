package broker

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	ClientID string `koanf:"clientID"`
	Username string `koanf:"user"`
	Password string `koanf:"password"`
}
