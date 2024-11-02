package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"strconv"
)

type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var tasksFile = "tasks.json"

func loadTasks() ([]Task, error) {
    file, err := ioutil.ReadFile(tasksFile)
    if err != nil {
        if os.IsNotExist(err) {
            return []Task{}, nil // Return an empty list if the file doesn't exist
        }
        return nil, err
    }

    var tasks []Task
    if err := json.Unmarshal(file, &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(tasksFile, data, 0644)
}

func addTask(title string) error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }

    newTask := Task{
        ID:        len(tasks) + 1,
        Title:     title,
        Completed: false,
    }
    tasks = append(tasks, newTask)
    return saveTasks(tasks)
}

func listTasks() error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }

    fmt.Println("Your Tasks:")
    for _, task := range tasks {
        status := "❌"
        if task.Completed {
            status = "✅"
        }
        fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
    }
    return nil
}

func completeTask(id int) error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }

    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Completed = true
            break
        }
    }
    return saveTasks(tasks)
}

func deleteTask(id int) error {
    tasks, err := loadTasks()
    if err != nil {
        return err
    }

    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
    return saveTasks(tasks)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go-task-manager <command> [arguments]")
        fmt.Println("Commands:")
        fmt.Println("  add <task>       Add a new task")
        fmt.Println("  list             List all tasks")
        fmt.Println("  complete <id>    Mark a task as completed")
        fmt.Println("  delete <id>      Delete a task")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go-task-manager add <task>")
            return
        }
        title := os.Args[2]
        if err := addTask(title); err != nil {
            fmt.Println("Error:", err)
        }
    case "list":
        if err := listTasks(); err != nil {
            fmt.Println("Error:", err)
        }
    case "complete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go-task-manager complete <id>")
            return
        }
        id := os.Args[2]
        if taskID, err := strconv.Atoi(id); err == nil {
            if err := completeTask(taskID); err != nil {
                fmt.Println("Error:", err)
            }
        } else {
            fmt.Println("Invalid ID")
        }
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go-task-manager delete <id>")
            return
        }
        id := os.Args[2]
        if taskID, err := strconv.Atoi(id); err == nil {
            if err := deleteTask(taskID); err != nil {
                fmt.Println("Error:", err)
            }
        } else {
            fmt.Println("Invalid ID")
        }
    default:
        fmt.Println("Unknown command:", command)
    }
}


