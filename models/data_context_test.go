package models

import (
	"testing"
)

func TestInterface(t *testing.T) {

	targets := []DataContext{NewInMemoryDataContext()}

	for _, c := range targets {

		task00 := NewTask("test00", 8)
		task01 := NewTask("test01", 16)
		c.UpdateTask(&task00)
		c.UpdateTask(&task01)

		tasks := c.GetTasks()
		if len(tasks) != 2 {
			t.Fail()
			return
		}
		for _, task := range tasks {
			t.Log(task)
		}

		c.DeleteTask(task00.ID)
		tasks = c.GetTasks()
		if len(tasks) != 1 {
			t.Fail()
			return
		}
		for _, task := range tasks {
			t.Log(task)
		}

	}

}
