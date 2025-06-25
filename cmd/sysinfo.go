package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
)

func NewSysInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sysinfo",
		Short: "Show basic system info",
		Run: func(cmd *cobra.Command, args []string) {

			info, _ := host.Info()
			users, _ := host.Users()
			cmd.Println("OS:", headerStyle.Render(info.Platform, info.PlatformVersion))
			cmd.Println("Uptime:", headerStyle.Render(fmt.Sprintf("%d seconds",info.Uptime)))
			cmd.Println("User:",headerStyle.Render(users[0].User))
		},
	}
}
