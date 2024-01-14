package enviroment

import (
	"log"

	"github.com/spf13/viper"
)

type Enviroment struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         uint   `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 uint   `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  uint   `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour uint   `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GinMode                string `mapstructure:"GIN_MODE"`
	DBMode                 string `mapstructure:"DB_MODE"`
}

func New() *Enviroment {
	return &Enviroment{}
}

func (e *Enviroment) Load() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(e)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if e.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
}
