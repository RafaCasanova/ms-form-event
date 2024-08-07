package main

import (
	"challenger/adapter/input/controller"
	"challenger/adapter/input/controller/routes"
	"challenger/adapter/output/repository"
	"challenger/app/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	log.Println("Start application")
	repository := repository.NewContactRepository("Teste")
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
	env, _ := os.ReadFile(".env")
	envs := strings.Split(string(env), "\n")

	for _, env := range envs {
		envTrim := strings.ReplaceAll(env, " ", "")
		variables := strings.SplitN(envTrim, "=", 2)
		os.Setenv(strings.TrimSpace(variables[0]), strings.TrimSpace(variables[1]))
	}
}
