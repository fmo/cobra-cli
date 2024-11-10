package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Test usage",
	Short: "A test usage",
	Long: `A test usage with long
			description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command is being executed")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
