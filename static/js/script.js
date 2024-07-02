document.addEventListener("DOMContentLoaded", function() {
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
                messageElement.textContent = message.username + ": " + message.message;
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
