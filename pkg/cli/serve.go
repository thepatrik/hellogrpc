package cli

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/thepatrik/hellogrpc/pkg/server"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run gRPC server",
	Long:  "Run gRPC server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetInt("port")
	if port == 0 {
		showHelp(cmd, "missing port")
	}

	server := server.New(server.WithLogger(logger))

	done := make(chan bool)
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		sig := <-gracefulStop
		logger.Printf("caught signal (%s). Gracefully shutting down...\n", sig)

		server.GracefulStop()

		close(done)
	}()

	if err := server.Serve(port); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	<-done
	logger.Println("Server stopped")

}

func init() {
	serveCmd.Flags().IntP("port", "p", 9090, "Port number")
	AddCommand(serveCmd)
}
