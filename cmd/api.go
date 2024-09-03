package cmd

import (
	"boilerplate/api"

	"github.com/spf13/cobra"
)

func apiCmd() *cobra.Command {
	var port int
	var command = &cobra.Command{
		Use:   "api",
		Short: "Run api server",
		Run: func(cmd *cobra.Command, args []string) {
			srv := api.NewServer()
			srv.Run(port)
		},
	}

	command.Flags().IntVar(&port, "port", 8080, "port server")
	return command
}
