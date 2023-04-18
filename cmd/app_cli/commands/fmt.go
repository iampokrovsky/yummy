package commands

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"hash/fnv"
	"os"
	"regexp"
	"strconv"
	"strings"
	"yummy/cmd/app_cli/commands/utils"
)

func splitSentences(line string) []string {
	return regexp.MustCompile(`[\p{Lu}][^\p{Lu}]*`).FindAllString(line, -1)
}

func hasEndMarks(s string) bool {
	return regexp.MustCompile(`\p{P}$`).MatchString(s)
}

func formatLine(sb *strings.Builder, line string) string {
	defer sb.Reset()

	// Add tab
	sb.WriteString("\t")

	sens := splitSentences(line)

	for i, sen := range sens {
		sens[i] = strings.TrimSpace(sen)
		sb.WriteString(sens[i])

		if !hasEndMarks(sens[i]) {
			// Add period
			sb.WriteString(".")
		}

		sb.WriteString(" ")
	}

	sb.WriteString("\n")

	return sb.String()
}

func tempFilename(filename string) (string, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(filename))
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(h.Sum32())), nil
}

func formatText(cmd *cobra.Command, args []string) {
	filename := args[0]

	// Open the input file for reading
	inFile, err := os.Open(filename)
	if err != nil {
		cmd.PrintErrln(err)
	}
	defer inFile.Close()

	// Get temp file name
	tempname, err := tempFilename(filename)
	if err != nil {
		cmd.PrintErrln(err)
	}

	// Create a temporary file for writing the modified content
	outFile, err := os.CreateTemp("", tempname)
	if err != nil {
		cmd.PrintErrln(err)
	}
	defer os.Remove(outFile.Name()) // delete the temporary file when done
	defer outFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inFile)

	// Create a writer to write the output file line by line
	writer := bufio.NewWriter(outFile)

	var sb strings.Builder

	// Loop through each line in the input file
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// Format line
		newLine := formatLine(&sb, line)

		// Write the line to the output file
		_, err = writer.WriteString(newLine)
		if err != nil {
			cmd.PrintErrln(err)
		}
	}

	// Check for scanner errors
	if err = scanner.Err(); err != nil {
		cmd.PrintErrln(err)
	}

	// Flush the writer to write any remaining data to the output file
	if err = writer.Flush(); err != nil {
		cmd.PrintErrln(err)
	}

	// Close both files
	if err = inFile.Close(); err != nil {
		cmd.PrintErrln(err)
	}
	if err = outFile.Close(); err != nil {
		cmd.PrintErrln(err)
	}

	// Rename the temporary file to the original filename
	if err = os.Rename(outFile.Name(), filename); err != nil {
		cmd.PrintErrln(err)
	}
}

var ErrWrongTxtFilePath = errors.New("wrong path to .txt file")

func (cli *CLI) fmtCmd() *cobra.Command {
	// Create regex for checking path to txt file with a filename
	txtRegex, err := regexp.Compile(`^.+\.txt$`)
	if err != nil {
		fmt.Println(err)
	}

	cmd := &cobra.Command{
		Use:   "fmt [file]",
		Short: "Format text",
		Long:  "Insert a tab before each paragraph and puts a period at the end of sentences",
		Args:  utils.ValidArgs(1, txtRegex, ErrWrongTxtFilePath),
		Run:   formatText,
	}

	return cmd
}
