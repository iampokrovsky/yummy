package commands

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "ListAll subjects",
	}

	return cmd
}
