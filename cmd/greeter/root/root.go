package root
 import
 (	"fmt"
	"github.com/spf13/cobra"
	"greeter/cmd/greeter/greet"
	 )
func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greeter",
		Short: "A simple greeting application",
		Long:  "This application greets the user with a friendly message.",	
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the greeter application!")
		},
	}

     cmd.AddCommand(greet.GreetCommand())

	return cmd	

}