package main

import (
	"github.com/thernande/app-pedidos-fondos-backend/internal/model"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/appLogs"
)

func main() {
	logger := &appLogs.Logger{}
	logger.Init()
	logger.InfoLogPrint("Hello World")
	model.Migrate()
}
