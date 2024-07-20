package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

type Client struct {
    conn net.Conn
    name string
}

var (
    clients  = make(map[net.Conn]Client)
    messages = make(chan string)
)

func handleConnection(conn net.Conn) {
    reader := bufio.NewReader(conn)

    // Get client name
    fmt.Fprintf(conn, "Enter your name: ")
    name,_:= reader.ReadString('\n')
    name = name[1 :len(name)-1] // Remove newline character 

    // Add client to the map
    clients[conn] = Client{conn: conn, name: name}

    // Notify everyone about the new client
    messages <- fmt.Sprintf("%s has joined the chat\n", name)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            delete(clients, conn)
            messages <- fmt.Sprintf("%s has left the chat\n", name)
            conn.Close()
            return
        }
        messages <- fmt.Sprintf("%s: %s", name, msg)
    }
}

func broadcastMessages() {
    for {
        msg := <-messages
        for _, client := range clients {
            _, err := fmt.Fprintf(client.conn, msg)
            if err != nil {
                delete(clients, client.conn)
                client.conn.Close()
            }
        }
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }
    defer listener.Close()

    go broadcastMessages()

    fmt.Println("Server started on port 8080")
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}



