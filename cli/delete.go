package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete subject",
	}

	return cmd
}
