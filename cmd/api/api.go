package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"github.com/talon-one/talon-backend-assingment/internal/infrastructure/nats"
	"github.com/talon-one/talon-backend-assingment/internal/interface/middleware"
	"go.uber.org/zap"
	"net/http"
)

var apiCmd = &cobra.Command{
	Use:  "api",
	Long: "run api server",
	RunE: withApp(runRestApiCmd),
}

func runRestApiCmd(o *cmdOpts, cmd *cobra.Command, args []string) error {
	subject, err := cmd.Flags().GetString("subject")
	if err != nil {
		return err
	}
	queue, err := nats.New(o.Cfg.Nats.URI, subject)
	if err != nil {
		return err
	}
	auditLogger := middleware.NewAuditLogging(queue)

	// HTTP router
	r := mux.NewRouter()
	r.Use(auditLogger.Logging)

	setupRoutes(r, o.DBConn)
	// HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", o.Cfg.HTTPPort),
		Handler: r,
	}

	zap.L().Info("listening", zap.String("port", o.Cfg.HTTPPort))
	return srv.ListenAndServe()
}

// I'm not moving this because it was in default of project, but it must be in proper place like infrastructure or presentation
func setupRoutes(r *mux.Router, conn *pgxpool.Pool) {
	// health endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := conn.Exec(context.Background(), "select 1")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("ook!!!!!!")))
	}).Methods(http.MethodGet)
}

func init() {
	apiCmd.Flags().String("subject", "logs", "NATS subject to subscribe to (required)")
	rootCmd.AddCommand(apiCmd)
}
