package utils

import (
	"github.com/spf13/cobra"
	"regexp"
)

func ValidArgs(n int, r *regexp.Regexp, rErr error) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(n)(cmd, args); err != nil {
			return err
		}

		for _, arg := range args {
			if !r.MatchString(arg) {
				return rErr
			}
		}

		return nil
	}
}
