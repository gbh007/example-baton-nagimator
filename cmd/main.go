package main

import (
	"app/internal/controller"
	"flag"
)

func main() {
	addr := flag.String("addr", ":8080", "web server address")
	debug := flag.Bool("d", false, "debug mode")
	dbType := flag.String("db", "sqlite", "db type sqlite, postgres, mysql")
	conn := flag.String("conn", "test.db", "db connection string")
	flag.Parse()

	c, err := controller.New(
		*addr,
		*debug,
		*dbType,
		*conn,
	)
	if err != nil {
		panic(err)
	}

	err = c.Serve()
	if err != nil {
		panic(err)
	}
}
