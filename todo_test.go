package main

import (
	"fmt"
	"testing"
)

type TestReaderWriter struct {
	inner []byte
}

func (tw *TestReaderWriter) Write(p []byte) (n int, err error) {
	tw.inner = append(tw.inner, p...)
	return len(p), nil
}

func (tw *TestReaderWriter) Read(p []byte) (n int, err error) {
	n = copy(p, tw.inner)
	return n, nil
}

func testShow(t *testing.T, testWriter TestReaderWriter, taskList TaskList, expected string) {
	show(&testWriter, &taskList)
	actual := string(testWriter.inner)
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestShow_Empty(t *testing.T) {
	testWriter := TestReaderWriter{}
	var taskList TaskList
	expected := "No tasks yet\n"
	testShow(t, testWriter, taskList, expected)
}

func TestShow_SingleItem(t *testing.T) {
	testWriter := TestReaderWriter{}
	var taskList TaskList
	taskList = append(taskList, Task{
		description: "Some Task",
		completed:   true,
	})
	expected := "1. [X] Some Task\n\n"
	testShow(t, testWriter, taskList, expected)
}
func TestAdd_happy(t *testing.T) {
	var tasks TaskList
	output := TestReaderWriter{}
	expectedDescription := "Do jj"
	input := TestReaderWriter{
		inner: []byte(fmt.Sprintf("%s\n", expectedDescription)),
	}
	add(&output, &input, &tasks)
	if len(tasks) != 1 {
		t.Error("add didn't add any tasks")
	}
	addedTask := tasks[0]
	if addedTask.completed {
		t.Error("added task should not be completed")
	}
	if addedTask.description != expectedDescription {
		t.Errorf("added task description wrong! actual: %q, expected: %q", addedTask.description, expectedDescription)
	}
}

func TestAdd_empty_input(t *testing.T) {
	var tasks TaskList
	output := TestReaderWriter{}
	input := TestReaderWriter{
		inner: []byte(" \n"),
	}
	emptyErr := add(&output, &input, &tasks)
	if len(tasks) != 0 {
		t.Errorf("add can't add empty input to the task, taskList: %v", tasks)
	}
	if emptyErr.Error() != "you can't add an empty task" {
		t.Error("adding an empty task should return an error")
	}
}
