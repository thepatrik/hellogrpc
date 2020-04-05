package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hellogrpc",
	Short: "hellogrpc is a cli for testing gRPC in golang",
	Long:  "hellogrpc is a command line interface for testing gRPC in golang",
}

// Execute executes cli
func Execute() error {
	return rootCmd.Execute()
}

// AddCommand adds one or more commands to root command.
func AddCommand(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}

func showHelp(cmd *cobra.Command, msg string) {
	fmt.Println(msg)
	_ = cmd.Help()
	os.Exit(0)
}
