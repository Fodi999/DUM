// websocket.go
package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "sync"
    "bufio"
    "fmt"
    "dum/util" // Импортируем утилиты

    "github.com/gorilla/websocket"
)

var (
    wsClients      = make(map[*websocket.Conn]bool)
    wsClientsAbout = make(map[*websocket.Conn]bool)
    mu             sync.Mutex
    broadcast      = make(chan Message)
    bots           = make(map[string]string)
    upgrader       = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    circularBuffer = util.NewCircularBuffer(100) // Создаем круговой буфер на 100 элементов
)

type Message struct {
    Username string `json:"username"`
    Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    log.Println("New WebSocket connection request received")
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }
    defer ws.Close()

    mu.Lock()
    wsClients[ws] = true
    mu.Unlock()

    log.Println("WebSocket connection established")

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("Read error: %v", err)
            mu.Lock()
            delete(wsClients, ws)
            mu.Unlock()
            break
        }

        log.Printf("Received message: %s", msg.Message)
        broadcast <- msg

        // Отправка сообщения в Telegram бот
        sendMessageToTelegram(msg.Message)
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        circularBuffer.Add(fmt.Sprintf("%s: %s", msg.Username, msg.Message)) // Добавляем сообщение в круговой буфер
        mu.Lock()
        for client := range wsClients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Write error: %v", err)
                client.Close()
                delete(wsClients, client)
            } else {
                log.Printf("Message broadcasted to client: %s", msg.Message)
            }
        }
        mu.Unlock()
    }
}

func sendMessageToTelegram(message string) {
    botToken := bots["1"]
    chatID := "1142224362"

    if botToken == "" {
        log.Println("Error: Telegram bot token is missing")
        return
    }

    if chatID == "" {
        log.Println("Error: Telegram chat ID is missing")
        return
    }

    telegramAPI := "https://api.telegram.org/bot" + botToken + "/sendMessage"
    values := url.Values{}
    values.Set("chat_id", chatID)
    values.Set("text", message)

    log.Printf("Sending message to Telegram chat ID: %s", chatID)
    resp, err := http.PostForm(telegramAPI, values)
    if err != nil {
        log.Println("Error sending message to Telegram:", err)
        return
    }
    defer resp.Body.Close()

    var result struct {
        OK     bool   `json:"ok"`
        Error  string `json:"description"`
        Code   int    `json:"error_code"`
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error reading response body from Telegram:", err)
        return
    }

    log.Printf("Telegram API response: %s", body)

    if err := json.Unmarshal(body, &result); err != nil {
        log.Println("Error decoding response from Telegram:", err)
        return
    }

    if !result.OK {
        log.Printf("Error: Telegram API returned not OK. Code: %d, Description: %s", result.Code, result.Error)
    } else {
        log.Println("Message sent to Telegram")
    }
}

func handleConnectionsAbout(w http.ResponseWriter, r *http.Request) {
    log.Println("New WebSocket connection request received for About page")
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }
    defer ws.Close()

    mu.Lock()
    wsClientsAbout[ws] = true
    mu.Unlock()

    log.Println("WebSocket connection established for About page")

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("Read error: %v", err)
            mu.Lock()
            delete(wsClientsAbout, ws)
            mu.Unlock()
            break
        }

        log.Printf("Received message: %s", msg.Message)
        broadcast <- msg

        // Отправка сообщения в Telegram бот
        sendMessageToTelegramAbout(msg.Message)
    }
}

func handleMessagesAbout() {
    for {
        msg := <-broadcast
        circularBuffer.Add(fmt.Sprintf("%s: %s", msg.Username, msg.Message)) // Добавляем сообщение в круговой буфер
        mu.Lock()
        for client := range wsClientsAbout {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Write error: %v", err)
                client.Close()
                delete(wsClientsAbout, client)
            } else {
                log.Printf("Message broadcasted to client: %s", msg.Message)
            }
        }
        mu.Unlock()
    }
}

func sendMessageToTelegramAbout(message string) {
    botToken := bots["2"]
    chatID := "1142224362"

    if botToken == "" {
        log.Println("Error: Telegram bot token is missing")
        return
    }

    if chatID == "" {
        log.Println("Error: Telegram chat ID is missing")
        return
    }

    telegramAPI := "https://api.telegram.org/bot" + botToken + "/sendMessage"
    values := url.Values{}
    values.Set("chat_id", chatID)
    values.Set("text", message)

    log.Printf("Sending message to Telegram chat ID: %s", chatID)
    resp, err := http.PostForm(telegramAPI, values)
    if err != nil {
        log.Println("Error sending message to Telegram:", err)
        return
    }
    defer resp.Body.Close()

    var result struct {
        OK     bool   `json:"ok"`
        Error  string `json:"description"`
        Code   int    `json:"error_code"`
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error reading response body from Telegram:", err)
        return
    }

    log.Printf("Telegram API response: %s", body)

    if err := json.Unmarshal(body, &result); err != nil {
        log.Println("Error decoding response from Telegram:", err)
        return
    }

    if !result.OK {
        log.Printf("Error: Telegram API returned not OK. Code: %d, Description: %s", result.Code, result.Error)
    } else {
        log.Println("Message sent to Telegram")
    }
}

func startWebSocketServer() {
    http.HandleFunc("/ws", handleConnections)
    go handleMessages()

    log.Println("WebSocket server for main page started on :8080/ws")
}

func startWebSocketServerAbout() {
    http.HandleFunc("/ws_about", handleConnectionsAbout)
    go handleMessagesAbout()

    log.Println("WebSocket server for About page started on :8080/ws_about")
}

func loadEnvVariables() {
    file, err := os.Open(".env")
    if err != nil {
        log.Fatalf("Ошибка при открытии файла .env: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
            continue
        }
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            log.Printf("Некорректная строка в .env: %s", line)
            continue
        }
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        if err := os.Setenv(key, value); err != nil {
            log.Fatalf("Ошибка при установке переменной окружения: %v", err)
        }
        if strings.HasPrefix(key, "TELEGRAM_BOT_TOKEN_") {
            botName := strings.ToLower(strings.TrimPrefix(key, "TELEGRAM_BOT_TOKEN_"))
            bots[botName] = value
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Ошибка при чтении файла .env: %v", err)
    }

    fmt.Println("Bots loaded from .env file:")
    for botName, token := range bots {
        fmt.Printf(" - %s: %s\n", botName, token)
    }
}









































