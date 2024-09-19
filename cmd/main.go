package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lucas4ndrade/CleanArchWorkshopAPI/api/http/handler"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/repository"
	"github.com/lucas4ndrade/CleanArchWorkshopAPI/coupon/usecase"
	mongo "github.com/lucas4ndrade/CleanArchWorkshopAPI/infra/mongodb"
)

func main() {
	mongoClient, err := mongo.New("mongodb://localhost:27017")
	if err != nil {
		fmt.Printf("Failed to connect to mongo!!! %v", err)
		return
	}
	defer mongoClient.Close()

	repo := repository.NewRepository(
		mongoClient.GetConnection(),
		"workshop-api",
		"coupon",
	)

	uc := usecase.NewUsecase(repo)

	router := handler.StartRoutes(uc)

	http.Handle("/", router)

	sv := &http.Server{
		Addr:         ":3000",
		Handler:      http.TimeoutHandler(http.DefaultServeMux, 10*time.Second, "Timeout"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Server started: %s\n", sv.Addr)
	if err := sv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error: %s", err.Error())
	}
}
