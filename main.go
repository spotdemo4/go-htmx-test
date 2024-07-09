package main

import (
	"go-htmx-test/db"
	"go-htmx-test/web/home"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	env := getEnv()
	db.Connect(env.Host, env.User, env.Password, env.DBName, env.Port, env.Timezone)

	homeHandler := home.HomeHandler{}

	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	e.Any("/", homeHandler.Any)
	e.Static("/assets", "assets")

	e.Logger.Fatal(e.Start(":1323"))
}

type envConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DBName   string `mapstructure:"DBNAME"`
	Port     int    `mapstructure:"PORT"`
	Timezone string `mapstructure:"TZ"`
}

func getEnv() envConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	config := envConfig{}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshaling env", err)
	}

	return config
}
