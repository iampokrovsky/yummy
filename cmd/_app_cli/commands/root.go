package commands

import (
	"github.com/spf13/cobra"
)

func (cli *CLI) rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yummy [commands]",
		Short: "Yammy CLI",
		Long:  "Yammy CLI is the command-line interface for interacting with Yammy food delivery service.",
	}

	return cmd
}
