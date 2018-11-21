package cobra

import (
	"fmt"

	"github.com/sachinnay/Gophercises/assignment17/secret/vault"
	"github.com/spf13/cobra"
)

//This package is used by cobra lib for command line interface

//setCmd used for command line command for set
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println("Error occured ::", err)
		}
		fmt.Println("Value sets successfully")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
