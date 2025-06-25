package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)

func NewDiskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Show disk usage",
		Run: func(cmd *cobra.Command, args []string) {
		usage, err := disk.Usage("/")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cmd.Println(headerStyle.Render(fmt.Sprintf("Disk Used: %.2f%% ", usage.UsedPercent)), fmt.Sprintf(" | (%.2f GB / %.2f GB)", float64(usage.Used)/1e9, float64(usage.Total)/1e9))
		},
	}
}
