package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

func addOffsetToLines(str, offset string) string {
	var sb strings.Builder
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(offset)
		sb.WriteString(line)
	}
	return sb.String()
}

func printCommand(c *cobra.Command, offset string) {
	var sb strings.Builder
	sb.WriteString(offset)
	sb.WriteString(c.Use)
	if c.HasSubCommands() {
		if c.Runnable() {
			sb.WriteString(" [command]")
		} else {
			sb.WriteString(" [subject]")
		}
	}
	if c.HasLocalFlags() {
		sb.WriteString(" [fields]")
	}
	sb.WriteString("   ")
	sb.WriteString(c.Short)
	c.Println(sb.String())

	if c.HasLocalFlags() {
		flags := strings.TrimRight(c.LocalFlags().FlagUsages(), "\n")
		fmt.Println(addOffsetToLines(flags, "  "))
	}

	offset += "  "
	for _, cmd := range c.Commands() {
		printCommand(cmd, offset)
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func RootHelpFunc() func(*cobra.Command, []string) {
	return func(c *cobra.Command, a []string) {
		var sb strings.Builder
		sb.WriteString("\n")
		sb.WriteString(c.Long)
		sb.WriteString("\n\nUsage:\n")
		sb.WriteString(c.UseLine())
		sb.WriteString("\n\nCommands:\n")
		c.Print(sb.String())

		commands := c.Commands()

		sort.Slice(commands, func(i, j int) bool {
			return boolToInt(commands[i].HasSubCommands()) < boolToInt(commands[j].HasSubCommands())
		})

		for _, cmd := range commands {
			printCommand(cmd, "")
		}

		c.Print("\nUse flags -h or --help with any command to get more information about it.\n\n")
	}
}
