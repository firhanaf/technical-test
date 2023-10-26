package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     int
	DBName     string
	DBHost     string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DBUsername"); found {
		app.DBUsername = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPassword"); found {
		app.DBPassword = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBName"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPort"); found {
		conv, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Error on DBPort")
		}
		app.DBPort = conv
		isRead = false
	}
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		app.DBHost = viper.Get("DBHOST").(string)
		app.DBName = viper.Get("DBNAME").(string)
		app.DBPassword = viper.Get("DBPASS").(string)
		app.DBUsername = viper.Get("DBUSER").(string)
		app.DBPort, err = strconv.Atoi(viper.Get("DBPORT").(string))
		if err != nil {
			log.Println("Error Port env : ", err.Error())
		}
	}
	return &app
}
