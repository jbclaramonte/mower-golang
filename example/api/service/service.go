package service

import (
	"log"

	"github.com/labstack/echo/v4"
)

type ApiImpl struct {
}

func (*ApiImpl) CreateGarden(ctx echo.Context) error {
	log.Println("inside CreateGarden")
	return nil
}
