package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vivek080/Go-PushNotifications/handler"
)

func main() {

	godotenv.Load(".env")
	http.HandleFunc("/pushNotification", handler.ServeHTTP)
	fmt.Println("Push Notification Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
