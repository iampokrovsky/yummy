package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) restoreCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore subject",
	}

	return cmd
}
