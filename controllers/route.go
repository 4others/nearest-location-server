package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/nearest-location-server/route"
	"github.com/julienschmidt/httprouter"
)

func GetSortedRoutes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	
	source := query.Get("src")
	if source == "" {
		http.Error(w, "source was not entered", 404)
	}

	road := route.Way{}
	road.Source = source

	// Get array of all destinations.
	routesFromUser, ok := query["dst"]
	if !ok || len(routesFromUser) == 0 {
		http.Error(w, "at least one destination needs to be entered", 404)
	}

	var routesToSort = make(route.SliceOfRoutes,len(routesFromUser)) 

	// Fill in routesToSort with values from routesFromUser.
	for i, v := range routesFromUser {
		routesToSort[i].Destination = v
	}

	var respCode int

	// Add values of duration and distance to each destination.
	for _, rt := range routesToSort {
		respCode = rt.CalculateDestinationValues(w, source)
		if respCode != 200 {
			http.Error(w, "error while calculating distances", respCode)
		}
	}

	sort.Sort(routesToSort)

	// Add sorted routesToSort to road.
	road.Routes = routesToSort

	// Marshal road data to JSON.
	roadJSON, err := json.Marshal(road)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", roadJSON)
}

