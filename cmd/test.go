package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tjololo/testtools-container/pkg/nettools"
	"os"
	"time"
)

var (
	port    int
	timeout time.Duration
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
	},
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect [host]",
	Short: "Test if it's possible to connect to host",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		time, err := nettools.TestConnectPossible(args[0], "tcp", port, timeout)
		if err != nil {
			fmt.Printf("Connect to %s on port %d failed after %s. Due to %v\n", args[0], port, time, err)
			os.Exit(1)
		}
		fmt.Printf("Connect to %s on port %d ok after %s\n", args[0], port, time)
		os.Exit(0)
	},
}

var lookupCmd = &cobra.Command{
	Use:   "lookup [hostname]",
	Short: "Test dns lookup",
	Run: func(cmd *cobra.Command, args []string) {
		ips, time, err := nettools.TestDNSLookup(args[0])
		if err != nil {
			fmt.Printf("Lookup of %s failed after %s. Due to %v\n", args[0], time, err)
			os.Exit(1)
		}
		fmt.Printf("Lookup success after %s. IPs returned: \n", time)
		for _, ip := range ips {
			fmt.Printf("%s IN A %s\n", args[0], ip.String())
		}
		os.Exit(0)
	},
}

var apiCmd = &cobra.Command{
	Use:   "api [uri]",
	Short: "Test api request",
	Run: func(cmd *cobra.Command, args []string) {
		time, status, err := nettools.TestApiRequest(args[0], timeout)
		if err != nil {
			fmt.Printf("API request GET %s failed after %s. Due to %v\n", args[0], time, err)
			os.Exit(1)
		}
		fmt.Printf("API request succeeded after %s. Responsecode: %d\n", time, status)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(connectCmd)
	testCmd.AddCommand(lookupCmd)
	testCmd.AddCommand(apiCmd)

	connectCmd.Flags().IntVarP(&port, "port", "p", 443, "Port to check")
	connectCmd.Flags().DurationVarP(&timeout, "timeout", "t", time.Second*30, "Connect timeout")
	apiCmd.Flags().DurationVarP(&timeout, "timeout", "t", time.Second*30, "Connect timeout")
}
