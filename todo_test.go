package main

import (
	"testing"
)

type TestWriter struct {
	inner []byte
}

func (tw *TestWriter) Write(p []byte) (n int, err error) {
	tw.inner = append(tw.inner, p...)
	return len(p), nil
}

func testShow(t *testing.T, testWriter TestWriter, taskList TaskList, expected string) {
	show(&testWriter, &taskList)
	actual := string(testWriter.inner)
	if expected != actual {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func TestShow_Empty(t *testing.T) {
	testWriter := TestWriter{}
	var taskList TaskList
	expected := "No tasks yet\n"
	testShow(t, testWriter, taskList, expected)
}

func TestShow_SingleItem(t *testing.T) {
	testWriter := TestWriter{}
	var taskList TaskList
	taskList = append(taskList, Task{
		description: "Some Task",
		completed:   true,
	})
	expected := "1. [X] Some Task\n\n"
	testShow(t, testWriter, taskList, expected)
}
