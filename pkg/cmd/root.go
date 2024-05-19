package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubernetes-playground",
		Short: "Kubernetes playground",
		Long:  "Kubernetes playground",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Starting Kubernetes playground")

			ticker := time.NewTicker(5 * time.Second)

			go func() {
				for {
					select {
					case <-ticker.C:
						log.Printf("Tick")
					}
				}
			}()

			// Wait for signal
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
			<-signalChan

			log.Printf("Shutting down Kubernetes playground")

			return nil
		},
	}

	return cmd
}

func Execute(ctx context.Context) {
	cobra.CheckErr(newCmdRoot().ExecuteContext(ctx))
}
