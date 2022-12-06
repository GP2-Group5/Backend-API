package main

import (
	"fmt"

	"github.com/GP2-Group5/Backend/config"
	"github.com/GP2-Group5/Backend/factory"
	"github.com/GP2-Group5/Backend/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
