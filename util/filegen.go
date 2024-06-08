package util

import (
    "fmt"
    "os"
)

// CreateHTMLFile создает HTML файл с заданным именем и содержимым.
func CreateHTMLFile(fileName, content string) error {
    err := os.MkdirAll("templates", os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating templates directory: %v", err)
    }

    filePath := "templates/" + fileName
    _, err = os.Stat(filePath)
    if os.IsNotExist(err) {
        file, err := os.Create(filePath)
        if err != nil {
            return fmt.Errorf("error creating HTML file: %v", err)
        }
        defer file.Close()

        _, err = file.WriteString(content)
        if err != nil {
            return fmt.Errorf("error writing to HTML file: %v", err)
        }
    } else if err == nil {
        // Файл уже существует, не возвращаем ошибку, просто выходим из функции
        return nil
    } else {
        return fmt.Errorf("error checking if file exists: %v", err)
    }
    return nil
}

// CreateDefaultHTMLFile создает стандартный HTML файл hello.html, если он не существует.
func CreateDefaultHTMLFile() error {
    content := `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="Default description">
        <link href="/static/css/style.css" rel="stylesheet">
        <title>Hello</title>
        <meta name="theme-color" content="#007bff">
    </head>
    <body>
        <nav>
            | <a href="/">Home</a> 
            | <a href="/about">About</a> 
            | <a href="/contact">Contact</a> 
            | <a href="/user">User</a>
        </nav>
        <h1>Hello, World!</h1>
    </body>
</html>`
    return CreateHTMLFile("hello.html", content)
}






