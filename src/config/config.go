package config

import (
	"cmn-express/src/pkgs/utils"
	"log/slog"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	BuildEnv string `mapstructure:"BUILD_ENV"`

	// service info
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    int    `mapstructure:"SERVICE_PORT"`
	ServiceTimeout int    `mapstructure:"SERVICE_TIMEOUT"`

	// database info
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUserName string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`

	MailFrom   string `mapstructure:"MAIL_FROM"`
	MailServer string `mapstructure:"MAIL_SERVER"`
	MailPort   string `mapstructure:"MAIL_PORT"`
	MailPass   string `mapstructure:"MAIL_PASS"`
}

func MustLoadConfig(configPath string) ServiceConfig {
	var result ServiceConfig
	var dirPath = utils.GetDirectoryPath(configPath)
	var fileName = utils.GetFileName(configPath)

	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var err = viper.ReadInConfig()

	if err != nil {
		slog.Error("can not read config file",
			slog.String("path", configPath),
			slog.Any("error", err.Error()),
		)
		panic(err)
	}

	err = viper.Unmarshal(&result)
	if err != nil {
		slog.Error("can not unmarshal struct",
			slog.Any("result", result),
			slog.Any("error", err.Error()),
		)
		panic(err)
	}

	return result
}
