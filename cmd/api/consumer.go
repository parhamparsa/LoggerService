package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/talon-one/talon-backend-assingment/internal/application/service/audit_log_service"
	"github.com/talon-one/talon-backend-assingment/internal/infrastructure/nats"
	"github.com/talon-one/talon-backend-assingment/internal/infrastructure/persistence/audit_log_repository_pg"
)

// consumeCmd represents the start command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "WaitForSignal the NATS consumer",
	RunE:  withApp(consumerCmd),
}

func init() {
	// Define flags
	consumeCmd.Flags().String("consumer", "logs_consumer", "NATS subject to subscribe to (required)")
	consumeCmd.Flags().String("subject", "logs", "NATS subject to subscribe to (required)")
	rootCmd.AddCommand(consumeCmd)
}

func consumerCmd(o *cmdOpts, cmd *cobra.Command, args []string) error {
	consumerName, err := cmd.Flags().GetString("consumer")
	if err != nil {
		return err
	}

	subject, err := cmd.Flags().GetString("subject")
	if err != nil {
		return err
	}

	queue, err := nats.NewConsumer(o.Cfg.Nats.URI, subject, consumerName)
	if err != nil {
		return err
	}

	service := audit_log_service.NewAuditLogService(audit_log_repository_pg.NewAuditLogRepositoryPostgres(o.DBConn))

	fmt.Println("running with subject:", subject)
	if err := queue.Consume(service); err != nil {
		return err
	}
	o.Registry.Register(func() error { return queue.Close() })
	return o.Registry.WaitForSignal()
}
