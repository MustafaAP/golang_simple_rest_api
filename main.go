package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MustafaAP/ProjectK-backend-Go/router"
)

func main() {
	fmt.Println("ProjectK")

	r := router.Router()

	// fmt.Sprintf(":%s", os.Getenv("PORT")) use this when deploy
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
	fmt.Println("Listening at port 3000")

}
