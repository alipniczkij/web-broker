package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	broker "github.com/alipniczkij/web-broker"
	"github.com/alipniczkij/web-broker/pkg/handler"
	"github.com/alipniczkij/web-broker/pkg/repository"
	"github.com/alipniczkij/web-broker/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading env variables: %s", err.Error())
	// }

	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// 	Password: os.Getenv("RDB_PASSWORD"),
	// })

	// if err != nil {
	// 	logrus.Fatalf("error initialization db: %s", err.Error())
	// }
	queue := make([]string, 0)
	repos := repository.NewRepository(&queue) // db
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(broker.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s", err.Error())
		}
	}()
	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured while shutting down http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
