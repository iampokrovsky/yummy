package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) updateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update subject",
	}

	return cmd
}
