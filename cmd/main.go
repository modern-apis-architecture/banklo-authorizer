package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"os"
)

func main() {
	container, err := BuildAppContainer()

	var port = flag.Int("port", 6666, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(e, container)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))

}
