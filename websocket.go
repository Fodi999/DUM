//websocket.go
package main

import (
	"bufio"
    "encoding/json"
    "log"
    "net"
    "net/http"
    
    
)

var wsClients = make(map[*websocketConn]bool)
var broadcast = make(chan Message)


type Message struct {
    Username string `json:"username"`
    Message  string `json:"message"`
}

type websocketConn struct {
    conn net.Conn
    rw   *bufio.ReadWriter
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Upgrade") != "websocket" || r.Header.Get("Connection") != "Upgrade" {
        http.Error(w, "Not a websocket handshake", http.StatusBadRequest)
        return
    }

    hijacker, ok := w.(http.Hijacker)
    if !ok {
        http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
        return
    }

    conn, rw, err := hijacker.Hijack()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    ws := &websocketConn{conn: conn, rw: rw}
    mu.Lock()
    wsClients[ws] = true
    mu.Unlock()

    go readMessages(ws)
    go writeMessages(ws)
}

func readMessages(ws *websocketConn) {
    for {
        message, err := ws.rw.ReadString('\n')
        if err != nil {
            log.Printf("error: %v", err)
            mu.Lock()
            delete(wsClients, ws)
            mu.Unlock()
            return
        }

        var msg Message
        err = json.Unmarshal([]byte(message), &msg)
        if err != nil {
            log.Printf("error: %v", err)
            continue
        }

        broadcast <- msg
    }
}

func writeMessages(ws *websocketConn) {
    for msg := range broadcast {
        msgBytes, err := json.Marshal(msg)
        if err != nil {
            log.Printf("error: %v", err)
            continue
        }

        ws.rw.WriteString(string(msgBytes) + "\n")
        ws.rw.Flush()
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        mu.Lock()
        for client := range wsClients {
            err := json.NewEncoder(client.rw).Encode(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.conn.Close()
                delete(wsClients, client)
            }
        }
        mu.Unlock()
    }
}

func startWebSocketServer() {
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()

    log.Println("WebSocket server started on :8080/ws")
}

