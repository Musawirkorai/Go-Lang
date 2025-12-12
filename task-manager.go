package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a single to-do item
type Task struct {
	Title     string
	Completed bool
}

func main() {
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Simple Task Manager ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task Completed")
		fmt.Println("4. Exit")
		fmt.Print("Choose: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(&tasks, reader)
		case "2":
			listTasks(tasks)
		case "3":
			completeTask(&tasks, reader)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func addTask(tasks *[]Task, reader *bufio.Reader) {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	*tasks = append(*tasks, Task{Title: title})
	fmt.Println("Task added!")
}

func listTasks(tasks []Task) {
	fmt.Println("\nYour Tasks:")
	for i, t := range tasks {
		status := "❌"
		if t.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, t.Title)
	}
}

func completeTask(tasks *[]Task, reader *bufio.Reader) {
	fmt.Print("Enter task number completed: ")
	var index int
	fmt.Scan(&index)

	if index <= 0 || index > len(*tasks) {
		fmt.Println("Invalid task number")
		return
	}

	(*tasks)[index-1].Completed = true
	fmt.Println("Marked as completed!")
}
