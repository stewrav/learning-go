package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
//	"github.com/stewrav/tri/cmd/todo"
)

func init() {
	RootCmd.AddCommand(addRun)
}

var addRun = &cobra.Command{
	Use:   "add",
	Short: "add a TODO note",
	Long: `This is the long description
for the tri add command.`,
	Run: addRunFunc,
}

func addRunFunc(cmd *cobra.Command, args []string) {
            //fmt.Println("add called")
			for _, x := range args {
				fmt.Println(x)  // just show the args for now
			}
		}

