package main

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"softline-test-task/internal/config"
	"softline-test-task/internal/hasher"
	"softline-test-task/internal/repo"
	"softline-test-task/internal/service"
	"softline-test-task/internal/token"
	"softline-test-task/internal/transport/rest"
	"softline-test-task/internal/validator"
	"syscall"
)

func main() {

	conf := config.GetConfig()

	// Открываем соединение с бд
	db, err := repo.OpenConnection(conf.Database)
	if err != nil {
		panic(err)
	}

	newRepository := repo.NewRepository(db)
	newHasher := hasher.NewHasher()
	newAuthToken := token.NewAuthToken(conf.Token.SecretWord, conf.Token.Expired)
	newValidator := validator.NewValidator()

	newService := service.NewService(newRepository, newHasher, newAuthToken, newValidator)
	newController := rest.NewController(newService)

	server := rest.CreateServer(newController, conf.Server)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server started on:", conf.Server.Port)

	stop := make(chan os.Signal)
	defer close(stop)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	if err := db.Close(); err != nil {
		log.Println(err.Error())
	}
	log.Println("Server stop")
}
