package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

func NewMemCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mem",
		Short: "Show memory usage",
		Run: func(cmd *cobra.Command, args []string) {
			vm, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(headerStyle.Render(fmt.Sprintf("Memory Used: %.2f%%", vm.UsedPercent)) +fmt.Sprintf(" | (%.2f GB / %.2f GB)", float64(vm.Used)/1e9, float64(vm.Total)/1e9))
		},
	}
}
