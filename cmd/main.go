package main

import "app/internal/controller"

func main() {
	c, err := controller.New()
	if err != nil {
		panic(err)
	}

	err = c.Serve()
	if err != nil {
		panic(err)
	}
}
