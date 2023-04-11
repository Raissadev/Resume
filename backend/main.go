package main

import (
	"api/src/config"
	"api/src/routes"
	. "api/src/utils"
	"fmt"
	"os"
	"runtime"
)

var env = Lenv.New()
var dataSource config.DataSource

func main() {
	fmt.Println("Hello World, 42")

	// dataSource.Open()

	routes := routes.New()

	routes.NewPublicAPI()
	routes.NewWeb()

	routes.Server.Listen(os.Getenv("SERVER_PORT"))

	config.Log.Printf(
		"Server v pid=%d started with processes: %d", os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

}
