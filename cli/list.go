package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List subjects",
	}

	return cmd
}
