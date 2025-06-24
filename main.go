package main

import (
	cmd 	 "github.com/krrishmahar/ministats/cmd" 
	lipgloss "github.com/charmbracelet/lipgloss"
	doc 	 "github.com/spf13/cobra/doc"
	"fmt"
	"os"
)

func main() {

	header := &doc.GenManHeader{
		Title:   "MINISTATS",
		Section: "1",
	}

	if err := os.MkdirAll("man", 0755); err != nil {
		fmt.Println("Failed to create man dir:", err)
		return
	}

	if err := doc.GenManTree(cmd.NewRootCmd(), header, "./man"); err != nil {
		fmt.Println("Error generating man page:", err)
	}

	rootCmd := cmd.NewRootCmd()
	
	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Padding(1, 2).
			Bold(true).
			Render(fmt.Sprintf("Error: %s", err))
		fmt.Println(errorStyle)
	}
}
