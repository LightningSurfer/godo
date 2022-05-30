// Copyright © 2022 Lightning Surfer

package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var filepath = func() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		return ".todo"
	}

	err = godotenv.Load()

	if err != nil {
		return homedir + "/.todo"
	}

	return homedir + "/" + os.Getenv("TODOS_LOCATION")
}()

const (
	uncheckedBox = "🔲"
	checkedBox   = "✅"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "CLI TODO app to go do things.",
	Long: `
  Welcome to GODO, a CLI TODO app to track things done and things left to go do.
  
  Functional, but not necessarily optimal. Use at your own risk.
  
  "Do. Or do not. There is no try." - Yoda
  `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
