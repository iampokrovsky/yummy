package cli

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Return subject",
	}

	return cmd
}
