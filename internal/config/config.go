package config

type Config struct {
	DB struct {
		URL string `yaml:"url" env:"DATABASE_URL"`
	} `yaml:"db"`
}
