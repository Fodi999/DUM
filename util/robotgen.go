package util

import (
    "fmt"
    "os"
)

// CreateRobotFolder создает новую папку для бота с заданным именем.
func CreateRobotFolder(name string) error {
    path := "robot/" + name
    if _, err := os.Stat(path); !os.IsNotExist(err) {
        return fmt.Errorf("folder %s already exists", path)
    }
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return fmt.Errorf("error creating folder: %v", err)
    }
    return nil
}

// GenerateBotCode генерирует код бота Telegram.
func GenerateBotCode(name string) error {
    path := "robot/" + name + "/bot.go"
    botCode := `package main

import (
    "log"
    "os"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName)
        msg.ReplyToMessageID = update.Message.MessageID

        bot.Send(msg)
    }
}
`
    if err := os.WriteFile(path, []byte(botCode), 0644); err != nil {
        return fmt.Errorf("error writing bot code: %v", err)
    }
    return nil
}

