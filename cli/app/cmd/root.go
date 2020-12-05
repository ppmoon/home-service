package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hscli",
		Short: "A cli for manager home service",
		Long: `Hscli is a terminal tool for manager your home service with systemd and podman.
For help type hscli -h or help
`,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
