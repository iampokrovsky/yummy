package utils

import "github.com/spf13/cobra"

func DefaultHelpFunc() func(*cobra.Command, []string) {
	cmd := &cobra.Command{}

	return cmd.HelpFunc()
}
