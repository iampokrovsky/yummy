package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new subject",
	}

	return cmd
}
