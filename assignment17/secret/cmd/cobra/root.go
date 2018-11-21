package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

//This package is used by cobra lib for command line interface

//RootCmd used for command line command
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is APIkey manager",
}
var encodingKey string

//To initialise arguments/flags before application start
func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "The key is for encoding and decoding secrets")
}

//to get secret file path
func secretPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
