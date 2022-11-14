/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tjololo/testtools-container/pkg/serve"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start testserver based on config",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configfile, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Printf("Failed to read config flag %v", err)
			os.Exit(1)
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			fmt.Printf("Failed to read port flag %v", err)
			os.Exit(1)
		}
		err = serve.StartTestsAndPrometheusEnpoint(configfile, port)
		if err != nil {
			fmt.Printf("Error occured while running tests %v", err)
			os.Exit(1)
		}
		fmt.Printf("Server shutdown completed\n")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("config", "c", "./tests-config.yaml", "Configfile containing tests config")
	serveCmd.Flags().IntP("port", "p", 2112, "Prometheus metrics port")
}
