package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Task struct {
	description string
	completed   bool
}

func (t *Task) String() string {
	completedString := " "
	if t.completed {
		completedString = "X"
	}
	return fmt.Sprintf("[%s] %s", completedString, t.description)
}

type TaskList []Task

func (tl *TaskList) String() string {
	var b strings.Builder
	if len(*tl) == 0 {
		b.WriteString("No tasks yet")
	}
	for i, task := range *tl {
		b.WriteString(fmt.Sprintf("%d. ", i+1))
		b.WriteString(task.String())
		b.WriteString("\n")
	}
	return b.String()
}

func show(w io.Writer, tasks *TaskList) {
	fmt.Fprintln(w, tasks.String())
}

func add(w io.Writer, r io.Reader, tasks *TaskList) error {
	reader := bufio.NewReader(r)
	fmt.Fprintln(w, "Enter a task:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return fmt.Errorf("you can't add an empty task")
	}
	newTask := Task{
		description: input,
		completed:   false,
	}
	*tasks = append(*tasks, newTask)
	fmt.Fprintln(w, tasks.String())
	return nil
}

func mark(w io.Writer, r io.Reader, tasks *TaskList) {
	var i int
	fmt.Fprintln(w, "Select the task you want to mark as completed")
	fmt.Fprintln(w, tasks)
	fmt.Fscan(r, &i)
	var s int
	if i > len(*tasks) {
		fmt.Fprintln(w, "choose a number from 1 to", len(*tasks))
		fmt.Fscan(r, &s)
		s = s - 1
		fmt.Fprintln(w, tasks.String())
	} else {
		s = s - 1
		fmt.Fprintln(w, tasks.String())
	}
}

func execute(tasks *TaskList) {
	var i int
	fmt.Println("select an option by the number: \n1. show tasks (completed and uncompleted)\n2. add task\n3. mark tast complete\n4. quit")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("You want to see your tasks")
		show(os.Stdout, tasks)
	case 2:
		fmt.Println("You want to add a task")
		addError := add(os.Stdout, os.Stdin, tasks)
		if addError != nil {
			fmt.Fprintln(os.Stderr, addError)
		}
	case 3:
		fmt.Println("You want to mark a task")
		mark(os.Stdout, os.Stdin, tasks)
	case 4:
		fmt.Println("You want to quit")
	default:
		fmt.Println("Choose between 1-4")
	}
}

func main() {
	var tasks TaskList
	for {
		execute(&tasks)
	}
}
