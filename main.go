package main

import (
	"fmt"
	"homework-server/handler"
	"homework-server/router"
	"homework-server/service"
	"net/http"
	"os"
)

func main() {
	service := service.NewUserService()
	handler := handler.NewUserHandler(service)
	r := router.NewRouter(handler)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)
}
