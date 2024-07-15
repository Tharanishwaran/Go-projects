package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var tasks []string
    // var command string
    scanner := bufio.NewScanner(os.Stdin)


    for {
        fmt.Print("Enter a command (add/view/exit): ")
        
        scanner.Scan()
        command := strings.TrimSpace(scanner.Text())

        switch command {
        case "add":
            fmt.Print("Enter a task : ")
            var task string
            scanner.Scan()
            task = strings.TrimSpace(scanner.Text())
            tasks = append(tasks,task)
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
