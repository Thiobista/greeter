package main

import (
	"github.com/spf13/cobra"
	"fmt"

)

func main() {

	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "A simple greeting application",
		Long:  "This application greets the user with a friendly message.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the greeter application!")
		},
}

greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet the user",
		Long:  "This command greets the user with a friendly message.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n",args[0])
		},
}
rootCmd.AddCommand(greetCmd)


if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
}
}