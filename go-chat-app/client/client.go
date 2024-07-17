package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    // Enter your name
    fmt.Print("Enter your name: ")
    name, _ := reader.ReadString('\n')
    fmt.Fprintf(conn, name)

    go readMessages(conn)

    for {
        msg, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, msg)
    }
}

func readMessages(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Disconnected from server")
            os.Exit(1)
        }
        fmt.Print(msg)
    }
}
