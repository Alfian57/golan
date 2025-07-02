package config

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type ServerConfig struct {
	Url string `env:"APP_URL" envDefault:"localhost:8000"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "3306"),
			Username: GetEnv("DB_USERNAME", ""),
			Password: GetEnv("DB_PASSWORD", ""),
			Name:     GetEnv("DB_NAME", "golang"),
		},
		Server: ServerConfig{
			Url: GetEnv("APP_URL", "localhost:8000"),
		},
	}

	return cfg, nil
}
