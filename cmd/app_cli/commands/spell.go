package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
	"yummy/cmd/app_cli/commands/utils"
)

var ErrNonEnglishWord = errors.New("the word contains non-English characters")

func (cli *CLI) spellCmd() *cobra.Command {
	// Create regex for checking word contains English characters
	engRegex, err := regexp.Compile("^[a-zA-Z]+$")
	if err != nil {
		fmt.Println(err)
	}

	cmd := &cobra.Command{
		Use:   "spell [word]",
		Short: "Print a word's letters separated by a space",
		Args:  utils.ValidArgs(1, engRegex, ErrNonEnglishWord),
		Run: func(cmd *cobra.Command, args []string) {
			word := args[0]
			letters := strings.Split(word, "")
			spaced := strings.Join(letters, " ")
			cmd.Println(spaced)
		},
	}

	return cmd
}
