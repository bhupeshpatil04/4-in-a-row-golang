package main

import (
    "log"
    "net/http"
    "fourinarow/internal/ws"
)

func main() {
    http.HandleFunc("/ws", ws.HandleWS)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
