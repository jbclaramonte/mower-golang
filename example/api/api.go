package main

import (
	apigen "github.com/jbclaramonte/mower-golang/example/api/gen"
	"github.com/jbclaramonte/mower-golang/example/api/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	apigen.RegisterHandlers(e, &service.ApiImpl{})
	e.Logger.Fatal(e.Start(":1323"))
}
