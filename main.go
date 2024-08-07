package main

import (
	"challenger/adapter/input/controller"
	"challenger/adapter/input/controller/routes"
	kafka_liste "challenger/adapter/output/listen/kafka"
	"challenger/adapter/output/repository"
	mongodb "challenger/app/config/database"
	"challenger/app/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	log.Println("Start application")
	log.Println("Connect to database")
	banco, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Error creating DB connection")
	}
	log.Println("Listening kafka server")
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatal("something went wrong with kafka connection")
	}
	conn.Close()
	go kafka_liste.InitConsumer()
	repository := repository.NewContactRepository(banco)
	service := service.NewContactServoce(repository)
	controller := controller.NewController(service)
	routes.ContactRouter(controller)
	log.Println("Application running at port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func init() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		log.Println("** Interruption received! **")
		log.Println("the program will be terminated...")
		time.Sleep(5 * time.Second)
		log.Println("Shutting down")
		os.Exit(1)
	}()
	loadEnv()
}

func loadEnv() {

	log.Println(" Loading environment variables ")
	env, _ := os.ReadFile(".Myenv")
	envs := strings.Split(string(env), "\n")

	for _, env := range envs {
		envTrim := strings.ReplaceAll(env, " ", "")
		variables := strings.SplitN(envTrim, "=", 2)
		os.Setenv(strings.TrimSpace(variables[0]), strings.TrimSpace(variables[1]))
	}
}
