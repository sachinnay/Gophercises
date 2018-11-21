package main

import (
	"errors"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment7/task/db"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	main()
	dashtest.ControlCoverage(m)
}
func TestMainError(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
	main()
}

func TestMainHomeDirError(t *testing.T) {
	tempDir := HomedirVar
	HomedirVar = func() (string, error) {
		return "", errors.New("Customised error")
	}
	main()
	defer func() {
		HomedirVar = tempDir
	}()

}
