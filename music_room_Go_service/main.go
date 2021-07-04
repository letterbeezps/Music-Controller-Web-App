package main

import (
	"musicRoom/models"
	"musicRoom/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	defer models.CloseDB()

	// fmt.Println(models.Generate_unique_code())

	s := &http.Server{
		Addr:    ":9898",
		Handler: router,
	}

	s.ListenAndServe()
}
