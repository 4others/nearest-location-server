package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nearest-location-server/controllers"
)

func main(){
	r := httprouter.New()	

	r.GET("/routes", controllers.GetSortedRoutes)

	log.Fatal(http.ListenAndServe(":8080", r))
}