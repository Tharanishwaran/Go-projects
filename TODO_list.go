package main

import (
    "fmt"
)

func main() {
    var tasks []string
    var command string

    for {
        fmt.Print("Enter a command (add/view/exit): ")
        fmt.Scan(&command)

        switch command {
        case "add":
            var task string
            fmt.Print("Enter a task: ")
            fmt.Scan(&task)
            tasks = append(tasks, task)
        case "view":
            fmt.Println("To-Do List:")
            for i, task := range tasks {
                fmt.Printf("%d. %s\n", i+1, task)
            }
        case "exit":
            return
        default:
            fmt.Println("Unknown command")
        }
    }
}
