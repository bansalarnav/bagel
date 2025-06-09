package swarm

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var SwarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Start a swarm to work on multiple PRs at once",
	Long:  `<TODO>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from Swarm")

		return errors.New("WIP")
	},
}
