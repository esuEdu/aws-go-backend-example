package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv    string `mapstructure:"APP_ENV"`
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    string `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASSWORD"`
	DBName    string `mapstructure:"DB_NAME"`
	AWSRegion string `mapstructure:"AWS_REGION"`
	AWSBucket string `mapstructure:"AWS_BUCKET"`
}

func LoadEnv() *Env {

	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
