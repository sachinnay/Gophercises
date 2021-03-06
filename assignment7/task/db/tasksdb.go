package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

//This file contains DB related oprations

var taskBucket = []byte("tasks")

//Db object for bolt db
var Db *bolt.DB

//Task  :: struct for task
type Task struct {
	Key   int
	Value string
}

// Init :: to initialize db connection and bucket
func Init(dbPath string) error {
	var err error
	Db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	fn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	}
	return Db.Update(fn)

}

//CreateTask :: creates the task
func CreateTask(task string) (int, error) {
	var id int
	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}

//AllTasks :: read all TODO tasks
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c :=
			b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//This function do the int to byte concersion
func itob(value int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(value))
	return b
}

//This function do the byte to int concersion
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

//DeleteTask :: delete the TODO task
func DeleteTask(id int) error {
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(id))
	})
}
