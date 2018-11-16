package main

import (
	"os"

	"github.com/CARVIN94/go-example/controller"
	"github.com/urfave/cli"

	"github.com/CARVIN94/go-restful"
	"github.com/CARVIN94/mgo"
)

func main() {
	app := cli.NewApp()
	app.Name = "Example"
	app.Usage = "A example service"
	app.Version = "0.0.1"
	app.Action = appAction
	// app.Commands = []cli.Command{}
	app.Run(os.Args)
}

func appAction(c *cli.Context) {
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

	quit := make(chan os.Signal)
	<-quit
}

func setRoute() (route restful.Route) {
	route.Get("/111", controller.Pipe, controller.Midd)
	route.Post("/111", restful.Urlencoded, controller.IndexHandler)
	// route.Get("/favicon.ico", ...)
	return route
}
