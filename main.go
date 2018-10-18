package main

import (
	"os"

	"github.com/CARVIN94/go-example/controller"

	"github.com/CARVIN94/go-restful"
	"github.com/CARVIN94/mgo"
)

func init() {
	mgo.Connect(&mgo.Config{
		Hosts:    "localhost:27017",
		Database: "mqtt",
		// Timeout:  time.Second * 60,
	})
	restful.Start(&restful.Config{
		Route: setRoute(),
		Port:  3000,
		// ReadTimeout:  time.Second * 10,
		// WriteTimeout: time.Second * 4,
	})
}

func setRoute() (route restful.Route) {
	route.Get("/111", controller.Pipe, controller.Midd)
	route.Post("/111", controller.IndexHandler)
	// route.Get("/favicon.ico", ...)
	return route
}

func main() {
	quit := make(chan os.Signal)
	<-quit
}
