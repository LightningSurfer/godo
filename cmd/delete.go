// Copyright © 2022 Lightning Surfer

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a TODO forever",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting TODO...")

		file, err := os.Open(filepath)

		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		// Copy each file line into an array
		var todos []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			todos = append(todos, scanner.Text())
		}

		file.Close()

		if len(todos) == 0 {
			log.Fatal("No todos to delete")
		}

		// Mark TODOs for deletion
		for _, arg := range args {
			toBeDeletedLineNumber, err := strconv.Atoi(arg)

			if err != nil {
				log.Fatal(err)
			}

			if len(todos) < toBeDeletedLineNumber {
				fmt.Printf("Invalid line number: %d", toBeDeletedLineNumber)
				break
			}

			if strings.HasPrefix(todos[toBeDeletedLineNumber-1], "x ") {
				todos[toBeDeletedLineNumber-1] = "- " + todos[toBeDeletedLineNumber-1][2:]
			} else {
				todos[toBeDeletedLineNumber-1] = "- " + todos[toBeDeletedLineNumber-1]
			}
			fmt.Printf("Deleted TODO: %s\n", todos[toBeDeletedLineNumber-1][2:])
		}

		// Remove TODOs
		var remainingTodos []string

		for _, todo := range todos {
			if strings.HasPrefix(todo, "- ") {
				continue
			} else {
				remainingTodos = append(remainingTodos, todo)
			}
		}

		file, err = os.Create(filepath)

		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		// Rewrite file
		for _, remainingTodo := range remainingTodos {
			_, err = file.WriteString(remainingTodo + "\n")

			if err != nil {
				file.Close()
				log.Fatalf("Error: %v", err)
			}
		}

		// Save file
		err = file.Sync()
		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
