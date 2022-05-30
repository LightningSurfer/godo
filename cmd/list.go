// Copyright © 2022 Lightning Surfer

package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all current TODOs",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current TODOs:")

		// Open file for reads. Create it if it does not exist.
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0755)
		defer file.Close()

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		reader := bufio.NewReader(file)

		// Print each todo
		for lineNumber := 1; ; lineNumber++ {
			line, _, err := reader.ReadLine()

			if err != nil {
				if err == io.EOF {
					// End loop at the end of the file
					if lineNumber == 1 {
						fmt.Println("\nNo TODOs to list 👻")
					}
					break
				}
				log.Fatalf("Error: %v", err)
			}

			todo := string(line)

			if strings.HasPrefix(todo, "x") {
				// Format completed todo
				fmt.Printf("%s %d. \033[0;32m%s\033[0m\n", checkedBox, lineNumber, strings.TrimSpace(todo[1:]))
			} else {
				// Format uncompleted todo
				fmt.Printf("%s %d. \033[0;35m%s\033[0m\n", uncheckedBox, lineNumber, strings.TrimSpace(todo))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
