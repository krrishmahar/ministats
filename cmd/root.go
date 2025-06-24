package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ministats",
		Short: "A minimal system resource monitor CLI written in Go",
		Long:  "ministats is a simple CLI tool that shows system CPU, memory, and disk usage using gopsutil.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'ministats --help' to see available commands")
		},
	}

	// Register subcommands
	rootCmd.AddCommand(NewMemCmd())
	rootCmd.AddCommand(NewDiskCmd())
	rootCmd.AddCommand(NewCpuCmdAlt())
	rootCmd.AddCommand(NewPsCmd())
	rootCmd.AddCommand(NewSysInfoCmd())

	return rootCmd
}
