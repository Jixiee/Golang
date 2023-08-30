//Write a program to create a To-Do List
package main
import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Text     string
	Complete bool
}

func main() {
	tasks := make([]Task, 0)

	for {
		fmt.Println("Todo List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Complete")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			taskText := scanner.Text()
			tasks = append(tasks, Task{Text: taskText, Complete: false})
			fmt.Println("Task added.")
			
		case 2:
			fmt.Print("Enter task number to mark as complete: ")
			var taskNumber int
			fmt.Scanln(&taskNumber)
			if taskNumber >= 1 && taskNumber <= len(tasks) {
				tasks[taskNumber-1].Complete = true
				fmt.Println("Task marked as complete.")
			} else {
				fmt.Println("Invalid task number.")
			}
			
		case 3:
			fmt.Println("Tasks:")
			for i, task := range tasks {
				status := " "
				if task.Complete {
					status = "✓"
				}
				fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
			}
			
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
