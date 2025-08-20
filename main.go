package main

import (
	"fmt"
	"greeter/cmd/greeter/root"
)

func main() {

	rootCmd := root.RootCommand()

if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
}
}