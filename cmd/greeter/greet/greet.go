package greet

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GreetCommand() *cobra.Command{
	return &cobra.Command{
		Use:   "greet",
		Short: "Greet the user",
		Long:  "This command greets the user with a friendly message.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n",args[0])
		},
}
}