// Copyright © 2022 Lightning Surfer

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new TODO",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new TODO...")

		// Open file for reads + writes in append mode. Create it if it does not exist.
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		todo := strings.Join(args, " ")
		_, err = fmt.Fprintln(file, todo)

		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		file.Close()
		fmt.Printf("New TODO: %s\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
