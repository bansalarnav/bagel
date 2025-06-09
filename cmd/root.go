package cmd

import (
	"fmt"
	"os"

	"github.com/bansalarnav/bagel/cmd/swarm"
	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

var rootCmd = &cobra.Command{
	Use:   "bagel",
	Short: "Bagel is your AI pair programmer in your terminal",
	Long:  `<TODO>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Bagel")
	},
}

func init() {
	rootCmd.AddCommand(swarm.SwarmCmd)
	rootCmd.Flags().BoolP("init", "i", false, "Initialize Bagel in the current directory")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("An unexpected error occured", err)
		os.Exit(1)
	}
}
