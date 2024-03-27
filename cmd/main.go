package main

import (
	Todo "github.com/buts00/ToDo"
	"github.com/buts00/ToDo/pkg/handler"
	"github.com/buts00/ToDo/pkg/repository"
	"github.com/buts00/ToDo/pkg/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"

	"os"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := gotenv.Load(); err != nil {
		logrus.Fatal("cannot load env variables: " + err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatal("error initializing config: " + err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
	})

	if err != nil {
		logrus.Fatal("cannot connect to database: " + err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(Todo.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error occurred while running http server" + err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
