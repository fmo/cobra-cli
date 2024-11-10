package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "Test usage",
	Short: "A test usage",
	Long: `A test usage with long
			description`,
	Run: func(cmd *cobra.Command, args []string) {
		RunGameLoop()
	},
}

func main() {
	rootCmd.AddCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var action1Cmd = &cobra.Command{
	Use:   "action 1",
	Short: "action one details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Action 1 command run")
	},
}

func RunGameLoop() {
	fmt.Println("Welcome to the interactive cli")

	for {
		fmt.Println("\nAvailable actions: action1, action2, exit")
		fmt.Println("Choose an action: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "action1":
			action1Cmd.Run(nil, nil)
		default:
			fmt.Println("Invalid action, please try again.")
		}
	}
}
