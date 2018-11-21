package cobra

import (
	"fmt"

	"github.com/sachinnay/Gophercises/assignment17/secret/vault"
	"github.com/spf13/cobra"
)

//This package is used by cobra lib for command line interface

//getCmd used for command line command for get

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from secret storage",
	Run: func(cmd *cobra.Command, args []string) {

		v := vault.File(encodingKey, secretPath())

		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("No value set")
			return
		}
		fmt.Printf("%s=%s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
