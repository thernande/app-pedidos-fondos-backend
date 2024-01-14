package main

import (
	"os"

	"github.com/thernande/app-pedidos-fondos-backend/internal/model"
	"github.com/thernande/app-pedidos-fondos-backend/pkg/appLogs"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		model.Migrate()
		os.Exit(0)
	}
	logger := &appLogs.Logger{}
	logger.Init()
	logger.InfoLogPrint("Hello World")
}
