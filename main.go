package main

import (
	"bufio"
	"fmt"
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

var tasks TaskList

func show() {
	fmt.Println(tasks.String())
}

func add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a task:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	newTask := Task{
		description: input,
		completed:   false,
	}
	tasks = append(tasks, newTask)
	fmt.Println(tasks.String())
}

func mark() {
	var i int
	fmt.Println("Select the task you want to mark as completed")
	fmt.Println(tasks)
	fmt.Scan(&i)
	var s int
	//var strikethroughText string
	if i > len(tasks) {
		fmt.Println("choose a number from 1 to", len(tasks))
		fmt.Scan(&s)
		s = s - 1
		// fmt.Println("\x1b[9m" + tasks[s] + "\x1b[0m")
		//tasks[s] = applyStrikethrough(tasks[s])
		fmt.Println(tasks.String())
	} else {
		s = s - 1
		//tasks[s] = applyStrikethrough(tasks[s])
		fmt.Println(tasks.String())
	}
}

// func applyStrikethrough(text string) string {
// 	result := ""
// 	for _, r := range text {
// 		result += string(r) + "\u0336"
// 	}
// 	return result
// }

func main() {
	var i int
	fmt.Println("select an option by the number: \n1. show tasks (completed and uncompleted)\n2. add task\n3. mark tast complete\n4. quit")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("You want to see your tasks")
		show()
	case 2:
		fmt.Println("You want to add a task")
		add()
	case 3:
		fmt.Println("You want to mark a task")
		mark()
	case 4:
		fmt.Println("You want to quit")
		break
	default:
		fmt.Println("Choose between 1-4")
	}
	main()
}
