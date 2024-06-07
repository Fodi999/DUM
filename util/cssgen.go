package util

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func GenerateCSS() {
    err := os.MkdirAll("static/css", os.ModePerm)
    if err != nil {
        fmt.Println("Error creating static/css directory:", err)
        return
    }

    htmlFiles, err := filepath.Glob("templates/*.html")
    if err != nil {
        fmt.Println("Error reading html files", err)
        return
    }
    classSet := make(map[string]struct{})
    classRegex := regexp.MustCompile(`class="([^"]+)"`)
    for _, file := range htmlFiles {
        content, err := os.ReadFile(file)
        if err != nil {
            fmt.Println("Error reading file:", err)
            return
        }
        matches := classRegex.FindAllSubmatch(content, -1)
        for _, match := range matches {
            classes := strings.Split(string(match[1]), " ")
            for _, class := range classes {
                classSet[class] = struct{}{}
            }
        }
    }

    cssClasses := map[string]string{
        "bg-gray-100": "background-color: #f7fafc;",
        "bg-gray-200": "background-color: #edf2f7;",
        "bg-gray-300": "background-color: #e2e8f0;",
        "bg-gray-400": "background-color: #cbd5e0;",
        "bg-gray-500": "background-color: #a0aec0;",
        "bg-gray-600": "background-color: #718096;",
        "bg-gray-700": "background-color: #4a5568;",
        "bg-gray-800": "background-color: #2d3748;",
        "bg-gray-900": "background-color: #1a202c;",
        "bg-red-100": "background-color: #fff5f5;",
        "bg-red-200": "background-color: #fed7d7;",
        "bg-red-300": "background-color: #feb2b2;",
        "bg-red-400": "background-color: #fc8181;",
        "bg-red-500": "background-color: #f56565;",
        "bg-red-600": "background-color: #e53e3e;",
        "bg-red-700": "background-color: #c53030;",
        "bg-red-800": "background-color: #9b2c2c;",
        "bg-red-900": "background-color: #742a2a;",
        "bg-blue-100": "background-color: #ebf8ff;",
        "bg-blue-200": "background-color: #bee3f8;",
        "bg-blue-300": "background-color: #90cdf4;",
        "bg-blue-400": "background-color: #63b3ed;",
        "bg-blue-500": "background-color: #4299e1;",
        "bg-blue-600": "background-color: #3182ce;",
        "bg-blue-700": "background-color: #2b6cb0;",
        "bg-blue-800": "background-color: #2c5282;",
        "bg-blue-900": "background-color: #2a4365;",
        "bg-green-100": "background-color: #f0fff4;",
        "bg-green-200": "background-color: #c6f6d5;",
        "bg-green-300": "background-color: #9ae6b4;",
        "bg-green-400": "background-color: #68d391;",
        "bg-green-500": "background-color: #48bb78;",
        "bg-green-600": "background-color: #38a169;",
        "bg-green-700": "background-color: #2f855a;",
        "bg-green-800": "background-color: #276749;",
        "bg-green-900": "background-color: #22543d;",
        "bg-yellow-100": "background-color: #fffff0;",
        "bg-yellow-200": "background-color: #fefcbf;",
        "bg-yellow-300": "background-color: #faf089;",
        "bg-yellow-400": "background-color: #f6e05e;",
        "bg-yellow-500": "background-color: #ecc94b;",
        "bg-yellow-600": "background-color: #d69e2e;",
        "bg-yellow-700": "background-color: #b7791f;",
        "bg-yellow-800": "background-color: #975a16;",
        "bg-yellow-900": "background-color: #744210;",
        "bg-white": "background-color: #ffffff;",
        "bg-black": "background-color: #000000;",

        "btn-primary": "background-color: #007bff; color: #fff; border-radius: 5px; padding: 10px 20px; cursor: pointer; border: none; outline: none; transition: all 0.3s ease;",
        "btn-primary:active": "background-color: #003680;",
        "btn-primary:focus": "box-shadow: 0 0 0 0.2rem rgba(0,123,255,.5);",
        "btn-primary:hover": "background-color: #0056b3; transform: scale(1.1);",

        "text-gray-700": "color: #4a5568;",
        "text-red-500": "color: #f56565;",
        "text-center": "text-align: center;",
        "p-2": "padding: 8px;",
    }

    file, err := os.Create("static/css/style.css")
    if err != nil {
        fmt.Println("Error creating CSS file:", err)
        return
    }
    defer file.Close()

    for class := range classSet {
        if style, exists := cssClasses[class]; exists {
            _, err := file.WriteString(fmt.Sprintf(".%s {%s}\n", class, style))
            if err != nil {
                fmt.Println("Error writing to CSS file:", err)
                return
            }
        }
    }
    globalStyle := "body, html {margin: 0; padding: 0; font-family: 'Arial', sans-serif; height: 100%;}\n\n@keyframes slideIn {\n\tfrom { transform: translateX(-100%); opacity: 0; }\n\tto { transform: translateX(0); opacity: 1; }\n}\n"
    _, err = file.WriteString(globalStyle)
    if err != nil {
        fmt.Println("Error writing to CSS file:", err)
        return
    }
}







