package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var JWT_SECRRET = ""

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	jwtKey     string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUsername = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPassword = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		conv, _ := strconv.Atoi(val)
		app.DBPort = conv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		// viper.SetConfigName("server")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config: ", err.Error())
			return nil
		}
		app.jwtKey = viper.Get("JWT_KEY").(string)
		app.DBUsername = viper.Get("DBUSER").(string)
		app.DBPassword = viper.Get("DBPASS").(string)
		app.DBHost = viper.Get("DBHOST").(string)
		app.DBPort, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DBName = viper.Get("DBNAME").(string)
	}
	JWT_SECRRET = app.jwtKey
	return &app
}
