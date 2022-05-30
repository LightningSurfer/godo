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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do 1+ TODOs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Doing TODO...")

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
			log.Fatalln("No todos to do")
		}

		// Checkmark done TODOs
		for _, arg := range args {
			toBeDoneLineNumber, err := strconv.Atoi(arg)

			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			if len(todos) < toBeDoneLineNumber {
				fmt.Printf("Invalid line number: %d", toBeDoneLineNumber)
				continue
			}

			if strings.HasPrefix(todos[toBeDoneLineNumber-1], "x ") {
				fmt.Printf("Could not do done TODO: %s\n", todos[toBeDoneLineNumber-1][2:])
			} else {
				fmt.Printf("Completed TODO: %s\n", todos[toBeDoneLineNumber-1])
				todos[toBeDoneLineNumber-1] = "x " + todos[toBeDoneLineNumber-1]
			}
		}

		// Reorder TODOs with uncompleted first & completed last
		var completedTodos, uncompletedTodos []string

		for _, todo := range todos {
			if strings.HasPrefix(todo, "x ") {
				completedTodos = append(completedTodos, todo)
			} else {
				uncompletedTodos = append(uncompletedTodos, todo)
			}
		}

		orderedTodos := append(uncompletedTodos, completedTodos...)

		// Either clear existing file for reuse or create a fresh one if it does not exist
		file, err = os.Create(filepath)

		if err != nil {
			file.Close()
			log.Fatalf("Error: %v", err)
		}

		// Rewrite file
		for _, orderedTodo := range orderedTodos {
			_, err = fmt.Fprintln(file, orderedTodo)

			if err != nil {
				file.Close()
				log.Fatalf("Error: %v", err)
			}
		}

		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
