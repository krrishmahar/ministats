package cmd

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)


func NewCpuCmdAlt() *cobra.Command {
	return &cobra.Command{
		Use:   "cpu",
		Short: "Show total CPU usage",
		Run: func(cmd *cobra.Command, args []string) {
			perc, err := cpu.Percent(500*time.Millisecond, false)
			if err != nil {
				fmt.Println("Error:",err)
				return 
			}
			fmt.Println("Total CPU Usage: ", headerStyle.Render(fmt.Sprintf("%.2f%%", perc[0])))
		},
	}
}
