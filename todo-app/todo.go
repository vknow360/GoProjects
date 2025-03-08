package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
	Category string `json:"category"`
	Deadline string `json:"deadline"`
}

func main() {
	LoadTasks()

	if len(os.Args) < 2 {
		printBanner()
		printHelp()
		os.Exit(0)
	}

	command := os.Args[1]
	switch command {
	case "help":
		printHelp()

	case "add":
		if len(os.Args) < 5 {
			color.Red("Usage: todo add <task> <priority:low|medium|high> <category> [deadline:YYYY-MM-DD]")
			os.Exit(1)
		}
		deadline := "None"
		if len(os.Args) > 5 {
			deadline = os.Args[5]
		}
		AddTask(os.Args[2], os.Args[3], os.Args[4], deadline)

	case "list":
		ListTasks()

	case "clear":
		RemoveAllTasks()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			color.Red("⚠️ Invalid task ID")
			return
		}
		UpdateTask(id)

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo remove <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			color.Red("⚠️ Invalid task ID")
			return
		}
		RemoveTask(id)
	}

}

var Tasks []Task = []Task{}

func printBanner() {
	color.Cyan("\n===================================")
	color.Cyan("   📌 Welcome to To-Do List CLI   ")
	color.Cyan("===================================\n")
}

func printHelp() {
	color.Blue("Available Commands:")
	fmt.Println("  -", color.GreenString("add <task> <priority> <category> [deadline]"), "  Add a new task")
	fmt.Println("  -", color.GreenString("list"), "                                         List all tasks")
	fmt.Println("  -", color.GreenString("done <id>"), "                                    Mark a task as done")
	fmt.Println("  -", color.GreenString("remove <id>"), "                                  Remove a task")
	fmt.Println("  -", color.GreenString("clear"), "                                        Remove all tasks")
}

func GetFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting user home directory: %v", err)
		os.Exit(1)
	}

	dirPath := homeDir + "/.todo-cli"
	filePath := dirPath + "/.tasks.json"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0700)
		os.WriteFile(filePath, []byte("[]"), 0600)
	}
	return filePath
}

func AddTask(title, priority, category, deadline string) {
	id := 1
	if len(Tasks) > 0 {
		id = Tasks[len(Tasks)-1].ID + 1
	}
	Tasks = append(Tasks, Task{
		ID:       id,
		Title:    title,
		Done:     false,
		Priority: priority,
		Category: category,
		Deadline: deadline,
	})
	SaveTasks()
	color.Green("✅ Task added: [%d] %s | Priority: %s | Category: %s | Deadline: %s", id, title, priority, category, deadline)
}

func SaveTasks() {
	filePath := GetFilePath()

	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		color.Red("❌ Error saving tasks: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		color.Red("❌ Error saving tasks: %v", err)
		os.Exit(1)
	}
}

func LoadTasks() []Task {
	filePath := GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		color.Red("❌ Error loading tasks: %v", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		color.Red("❌ Error loading tasks: %v", err)
		os.Exit(1)
	}
	return Tasks
}

func TasksSize() int {
	return len(Tasks)
}

func ListTasks() {
	if len(Tasks) == 0 {
		color.Yellow("No tasks found!")
		return
	}

	fmt.Println("\n📌 Your Tasks:")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("%-5s %-15s %-10s %-12s %-12s %-50s\n", "ID", "Status", "Priority", "Category", "Deadline", "Title")
	fmt.Println(strings.Repeat("-", 100))

	for _, task := range Tasks {
		status := "❌ Incomplete"
		if task.Done {
			status = "✅ Completed"
		}

		statusWidth := runewidth.StringWidth(status)
		paddedStatus := status + strings.Repeat(" ", 15-statusWidth)

		formattedTitle := runewidth.Truncate(task.Title, 50, "...")

		fmt.Printf("%-5d %s %-10s %-12s %-12s %-50s\n",
			task.ID, paddedStatus, task.Priority, task.Category, task.Deadline, formattedTitle)
	}
	fmt.Println(strings.Repeat("-", 100))
}

func UpdateTask(id int) {
	for i := range Tasks {
		if Tasks[i].ID == id {
			Tasks[i].Done = !Tasks[i].Done
			SaveTasks()
			status := "done"
			if !Tasks[i].Done {
				status = "undone"
			}
			color.Green("✅ Task [%d] marked as %s!", id, status)
			return
		}
	}
	color.Red("⚠️  Task ID [%d] not found!", id)
}

func RemoveTask(id int) {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			SaveTasks()
			color.Green("✅ Task [%d] removed!", id)
			return
		}
	}
	color.Red("⚠️  Task ID [%d] not found!", id)
}

func RemoveAllTasks() {
	Tasks = []Task{}
	SaveTasks()
	fmt.Println("✅ All tasks removed")
}
