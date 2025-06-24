package cmd

import (
	"fmt"
	"sort"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v3/process"
)


var (
	headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	rowStyle    = lipgloss.NewStyle().Padding(0, 1)
)

type ProcStat struct {
	PID     int32
	Name    string
	CPU     float64
	Memory  float32
}


func NewPsCmd() *cobra.Command {
	var sortBy string

	cmd := &cobra.Command{
		Use:   "ps",
		Short: "Show top 5 processes sorted by CPU or memory",
		Run: func(cmd *cobra.Command, args []string) {
			if len(sortBy) > 0{
				if sortBy != "cpu" && sortBy != "mem" {
					fmt.Println("Use --sort=cpu or --sort=mem")
					return
				}
				showTopProcesses(sortBy)
			}else {
				
				
				procs, _ := process.Processes()
				for _, p := range procs {
					cpu, _ := p.CPUPercent()
					mem, _ := p.MemoryPercent()
					name, _ := p.Name()
					if cpu > 10 {
						fmt.Printf("PID: %d\tCPU: %.2f%%\tMEM: %.2f%%\tName: %s\n", p.Pid, cpu, mem, name)
					}
				}
			}
		},
	}

	cmd.Flags().StringVar(&sortBy, "sort", "", "Sort processes by 'cpu' or 'mem'")
	return cmd
}
func showTopProcesses(mode string) {
	procs, _ := process.Processes()
	stats := make([]ProcStat, 0, len(procs))

	for _, p := range procs {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()
		stats = append(stats, ProcStat{PID: p.Pid, Name: name, CPU: cpu, Memory: mem})
	}

	if mode == "cpu" {
		sort.Slice(stats, func(i, j int) bool {
			return stats[i].CPU > stats[j].CPU
		})
	} else {
		sort.Slice(stats, func(i, j int) bool {
			return stats[i].Memory > stats[j].Memory
		})
	}

		// Header Style for Table
	fmt.Println(headerStyle.Render(fmt.Sprintf(" %-9s%-8s%-10s%-20s", "PID", "CPU", "MEM", "COMMAND")))

	// Printing top 5 processes
	for i := 0; i < 5 && i < len(stats); i++ {
		p := stats[i]
		// Using Lipgloss Align to format columns
		fmt.Println(rowStyle.Render(fmt.Sprintf("%-8d %-8.2f%% %-8.2f%% %-20s", p.PID, p.CPU, p.Memory, p.Name)))
	}
}


