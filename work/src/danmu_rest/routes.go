package danmu_rest

import (
	"net/http"
)

//Route ...structure of router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes ...a slice of Route
type Routes []Route

var controller = &Controller{MlController: MlController{}}

var routes = Routes{
	Route{
		"MostSimilar",
		"POST",
		"/most_similar",
		controller.MostSimilar,
	},
	Route{
		"RecentTopic",
		"POST",
		"/recent_topic",
		controller.RecentTopic,
	},
}
