package api

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"github.com/talon-one/talon-backend-assingment/internal/config"
	"github.com/talon-one/talon-backend-assingment/pkg"
	"github.com/talon-one/talon-backend-assingment/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type cmdOpts struct {
	Registry *pkg.ServiceRegistry
	Cfg      config.Config
	DBConn   *pgxpool.Pool
}

// WithAppHandler gets the app, service-provider and config as params to handle the command
type WithAppHandler func(o *cmdOpts, cmd *cobra.Command, args []string) error

func withApp(cmdF WithAppHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		logger.InitDefaultLogger(zapcore.InfoLevel)

		//zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

		sr := pkg.NewServiceRegistry()

		// environment variables
		cfg, err := config.Load()
		if err != nil {
			zap.L().Error("Failed to load config", zap.Error(err))
			return fmt.Errorf("failed to load config: %v", err)
		}

		conn, err := pgxpool.New(context.Background(), cfg.Database.ConnectionString())
		if err != nil {
			return fmt.Errorf("unable to connect to database: %v", err)
		}
		sr.Register(func() error {
			conn.Close()
			return nil
		})

		return cmdF(&cmdOpts{
			Registry: sr,
			Cfg:      cfg,
			DBConn:   conn,
		}, cmd, args)
	}
}
