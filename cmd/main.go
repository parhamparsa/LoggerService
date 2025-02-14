package main

import (
	"github.com/talon-one/talon-backend-assingment/cmd/api"
	"go.uber.org/zap"
	"os"
)

func main() {
	if err := api.Execute(); err != nil {
		zap.Error(err)
		os.Exit(1)
	}
}
