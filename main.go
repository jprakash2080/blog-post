package main

import (
	"log"
	"net/http"
//	database "blog/utils"
	"blog/routes"
)

func main() {
//	database.InitDB()
	//db := database.GetDB()
	//defer db.Close()

	router := routes.SetupRouter()

	port := ":8080"
	log.Printf("Server *** is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
