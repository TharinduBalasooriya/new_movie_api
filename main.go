package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TharinduBalasooriya/newmovieapi/routes"
)

func main() {
	r := routes.MovieRoutes()
	log.Println("Start listing ....")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))
}
