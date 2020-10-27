package route

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Way struct {
	Source string `json:"source"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"` 
}

// ListOfRoutes implements sort.Interface
type SliceOfRoutes []Route

func (slice SliceOfRoutes) Len() int { return len(slice)}
func (slice SliceOfRoutes) Less(i, j int) bool { 
	if slice[i].Duration == slice[j].Duration {
		return slice[i].Distance < slice[j].Distance
	}
	return slice[i].Duration < slice[j].Duration
}
func (slice SliceOfRoutes) Swap(i,j int) { slice[i], slice[j] = slice[j], slice[i] }

type RouterProject struct {
	Routes []struct {
		Legs []struct {
			Summary  string        `json:"summary"`
			Weight   int           `json:"weight"`
			Duration float64       `json:"duration"`
			Steps    []interface{} `json:"steps"`
			Distance float64       `json:"distance"`
		} `json:"legs"`
		WeightName string  `json:"weight_name"`
		Weight     int     `json:"weight"`
		Duration   float64 `json:"duration"`
		Distance   float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string    `json:"hint"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
	Code string `json:"code"`
}

// TODO this func needs to use http://project-osrm.org
func (r Route) CalculateDestinationValues(w http.ResponseWriter, src string) int {
	url := "http://router.project-osrm.org/route/v1/driving/" + src + ";" + r.Destination +"?overview=false"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	defer resp.Body.Close()

	var routerResponse RouterProject

	err = json.NewDecoder(resp.Body).Decode(&routerResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	r.Duration = routerResponse.Routes[0].Duration
	r.Distance = routerResponse.Routes[0].Distance

	// Convert response code to int
	respCode, err := strconv.Atoi(routerResponse.Code)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	return respCode
}

