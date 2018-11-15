package db

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestItob(t *testing.T) {
	b := itob(int(2))
	var a = []byte{0, 0, 0, 0, 0, 0, 0, 2}
	if !reflect.DeepEqual(a, b) {
		t.Error("Expected is : [0 0 0 0 0 0 0 2], getting :: ", b)
	}
}
func TestItob_minusVal(t *testing.T) {
	b := itob(int(-2))
	var a = []byte{255, 255, 255, 255, 255, 255, 255, 254}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected is : %v, getting :: %v", a, b)
	}
}

func TestBtoi(t *testing.T) {
	b := btoi([]byte{0, 0, 0, 0, 0, 0, 0, 2})
	a := 2
	if !reflect.DeepEqual(a, b) {
		t.Error("Expected is : 2")
	}
}
func TestInit_chkDbIfFilePresent(t *testing.T) {

	_, err := os.Stat(createConnection())
	if os.IsNotExist(err) {
		t.Error("Db file is not present\n")
	}
}
func createConnection() string {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "testdb.db")
	Init(dbPath)
	return dbPath
}

func TestCreateTask(t *testing.T) {

	_, err := CreateTask("Task1")
	if err != nil {
		t.Errorf("Error while adding task: %s \n", err)

	}
}

func TestAllTasks(t *testing.T) {
	tasks, err := AllTasks()
	if err != nil || len(tasks) == 0 {
		t.Errorf("Error in retiriving : %s \n", err)

	}
}

func TestDeleteTasks(t *testing.T) {
	err := DeleteTask(11)
	if err != nil {
		t.Errorf("Error while deleting task :%d Error is : %s \n", 10, err)

	}
}

func TestCreateTask_DbError(t *testing.T) {
	Db.Close()
	_, err := CreateTask("Task1")

	if err.Error() != "database not open" {
		t.Errorf("Expecting \"database not open\" \n")

	}
}
func TestAllTask_DbError(t *testing.T) {

	Db.Close()
	_, err := AllTasks()
	if err != nil && err.Error() != "database not open" {
		t.Errorf("Expecting \"database not open\" \n")

	}
}
func TestInit_chkDb(t *testing.T) {
	createConnection()
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "testdb.db")
	err := Init(dbPath)
	if err != nil && err.Error() != "timeout" {
		t.Errorf("Expecting \"timeout\" \n")

	}

}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
