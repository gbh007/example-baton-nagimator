package main

import (
	"app/internal/controller"
	"flag"
)

func main() {
	addr := flag.String("addr", ":8080", "web server address")
	debug := flag.Bool("d", false, "debug mode")
	flag.Parse()

	c, err := controller.New(*addr, *debug)
	if err != nil {
		panic(err)
	}

	err = c.Serve()
	if err != nil {
		panic(err)
	}
}
