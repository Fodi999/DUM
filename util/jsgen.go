
package util

import (
    "fmt"
    "os"
)

// CreateJSFile создает JS файл с заданным именем и содержимым.
func CreateJSFile(fileName, content string) error {
    err := os.MkdirAll("static/js", os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating js directory: %v", err)
    }

    filePath := "static/js/" + fileName
    _, err = os.Stat(filePath)
    if os.IsNotExist(err) {
        file, err := os.Create(filePath)
        if err != nil {
            return fmt.Errorf("error creating JS file: %v", err)
        }
        defer file.Close()

        _, err = file.WriteString(content)
        if err != nil {
            return fmt.Errorf("error writing to JS file: %v", err)
        }
    } else if err == nil {
        // Файл уже существует, не возвращаем ошибку, просто выходим из функции
        return nil
    } else {
        return fmt.Errorf("error checking if file exists: %v", err)
    }
    return nil
}

// CreateDefaultJSFile создает стандартный JS файл script.js, если он не существует.
func CreateDefaultJSFile() error {
    content := `document.addEventListener("DOMContentLoaded", function() {
    console.log("Hello, World!");

    const messagesDiv = document.getElementById("messages");
    const messageForm = document.getElementById("messageForm");
    const messageInput = document.getElementById("messageInput");

    let messageQueue = [];
    let ws = connectWebSocket();

    function connectWebSocket() {
        const socket = new WebSocket("ws://" + window.location.host + "/ws");

        socket.onopen = function() {
            console.log("WebSocket connection opened");
            while (messageQueue.length > 0) {
                const message = messageQueue.shift();
                sendMessage(socket, message);
            }
        };

        socket.onmessage = function(event) {
            console.log("WebSocket message received:", event.data);
            try {
                const message = JSON.parse(event.data);
                const messageElement = document.createElement("div");
                messageElement.classList.add("p-2", "my-2", "border-b", "border-gray-200");
                messageElement.textContent = "${message.username}: ${message.message}";
                messagesDiv.appendChild(messageElement);
                messagesDiv.scrollTop = messagesDiv.scrollHeight;
            } catch (e) {
                console.error("Error parsing WebSocket message:", e);
            }
        };

        socket.onclose = function() {
            console.log("WebSocket connection closed, attempting to reconnect...");
            setTimeout(() => {
                ws = connectWebSocket();
            }, 1000);
        };

        socket.onerror = function(error) {
            console.error("WebSocket error:", error);
        };

        return socket;
    }

    messageForm.addEventListener("submit", function(event) {
        event.preventDefault();
        const message = messageInput.value;
        if (message) {
            if (ws.readyState === WebSocket.OPEN) {
                sendMessage(ws, message);
            } else {
                messageQueue.push(message);
                console.error("WebSocket is not open. Current state: " + ws.readyState);
            }
            messageInput.value = '';
        }
    });

    function sendMessage(socket, message) {
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify({ username: "user", message: message }));
        } else {
            console.error("WebSocket is not open. Current state: " + socket.readyState);
            messageQueue.push(message);
        }
    }
});
`
    return CreateJSFile("script.js", content)
}




