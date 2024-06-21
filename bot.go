package main

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "strings"
    "fmt"
    "time"
    "bufio"
)

type Update struct {
    UpdateID int `json:"update_id"`
    Message  struct {
        MessageID int `json:"message_id"`
        From      struct {
            ID        int    `json:"id"`
            FirstName string `json:"first_name"`
            LastName  string `json:"last_name"`
            Username  string `json:"username"`
        } `json:"from"`
        Chat struct {
            ID int64 `json:"id"`
        } `json:"chat"`
        Text string `json:"text"`
    } `json:"message"`
}

func loadEnvVariables() {
    // Открытие файла .env
    file, err := os.Open(".env")
    if err != nil {
        log.Fatalf("Ошибка при открытии файла .env: %v", err)
    }
    defer file.Close()

    // Чтение файла построчно
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Пропуск комментариев и пустых строк
        if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
            continue
        }
        // Разделение строки на ключ и значение
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            log.Printf("Некорректная строка в .env: %s", line)
            continue
        }
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        // Установка переменной окружения
        if err := os.Setenv(key, value); err != nil {
            log.Fatalf("Ошибка при установке переменной окружения: %v", err)
        }
        // Добавление ботов в карту
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
func startBot(botToken string) {
    defer wg.Done()
    telegramAPI := "https://api.telegram.org/bot" + botToken
    updatesChan := make(chan Update)

    // Горутина для получения обновлений
    go func() {
        offset := 0
        for {
            updates, err := getUpdates(telegramAPI, offset)
            if err != nil {
                log.Println("Error getting updates:", err)
                time.Sleep(1 * time.Second)
                continue
            }

            for _, update := range updates {
                updatesChan <- update
                offset = update.UpdateID + 1
            }
        }
    }()

    // Горутина для обработки обновлений
    go func() {
        for update := range updatesChan {
            if update.Message.Text != "" {
                log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
                broadcast <- Message{Username: update.Message.From.Username, Message: update.Message.Text}
                err := sendMessage(telegramAPI, update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName)
                if err != nil {
                    log.Println("Error sending message:", err)
                }
            }
        }
    }()

    // Ожидание завершения всех горутин
    wg.Wait()
}

func getUpdates(telegramAPI string, offset int) ([]Update, error) {
    resp, err := http.Get(telegramAPI + "/getUpdates?offset=" + strconv.Itoa(offset))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        OK     bool     `json:"ok"`
        Result []Update `json:"result"`
    }

    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.Result, nil
}

func sendMessage(telegramAPI string, chatID int64, text string) error {
    values := url.Values{}
    values.Set("chat_id", strconv.FormatInt(chatID, 10))
    values.Set("text", text)

    resp, err := http.PostForm(telegramAPI+"/sendMessage", values)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    var result struct {
        OK bool `json:"ok"`
    }

    return json.NewDecoder(resp.Body).Decode(&result)
}

