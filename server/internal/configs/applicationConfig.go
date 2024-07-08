package configs

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"address"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"dbname"`
}
