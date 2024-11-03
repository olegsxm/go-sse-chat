package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Security struct {
	JWTKey         string `mapstructure:"jwt_secret"`
	PasswordPepper string `mapstructure:"password_pepper"`
}

type PGConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"database"`
}

type Config struct {
	ServerPort string   `mapstructure:"server_port"`
	Pg         PGConfig `mapstructure:"pg"`
	Security   Security `mapstructure:"security"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Pg.User,
		c.Pg.Password,
		c.Pg.Host,
		c.Pg.Port,
		c.Pg.DbName,
	)
}

func New() (Config, error) {
	cfg := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
