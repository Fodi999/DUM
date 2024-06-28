package util

import (
    "encoding/json"
    "fmt"
    "os"
    "dum/util/library"
)

const cssFilePath = "static/css/style.css"
const stateFilePath = "static/css/state.json"
const jsFilePath = "static/js/checkStyles.js"

func GenerateCSS() {
    err := os.MkdirAll("static/css", os.ModePerm)
    if err != nil {
        fmt.Println("Error creating static/css directory:", err)
        return
    }

    err = os.MkdirAll("static/js", os.ModePerm)
    if err != nil {
        fmt.Println("Error creating static/js directory:", err)
        return
    }

    classSet := collectClassesFromLibrary()

    // Load previous state
    previousClassSet := loadPreviousState()

    // Check if there are any new classes
    isNewClassAdded := false
    for class := range classSet {
        if _, exists := previousClassSet[class]; !exists {
            isNewClassAdded = true
            break
        }
    }

    if !isNewClassAdded {
        fmt.Println("No new classes added, skipping CSS generation.")
        return
    }

    cssClasses := collectCSSClasses()

    err = writeCSSFile(cssClasses, classSet)
    if err != nil {
        fmt.Println("Error writing CSS file:", err)
        return
    }

    // Save the current state
    saveCurrentState(classSet)

    // Optional: Print a message indicating the CSS was generated successfully
    fmt.Println("CSS generated successfully.")

    // Create JavaScript file for checking styles
    createJavaScriptFile(classSet)
}

func loadPreviousState() map[string]struct{} {
    previousClassSet := make(map[string]struct{})
    if _, err := os.Stat(stateFilePath); os.IsNotExist(err) {
        return previousClassSet
    }

    data, err := os.ReadFile(stateFilePath)
    if err != nil {
        fmt.Println("Error reading previous state file:", err)
        return previousClassSet
    }

    err = json.Unmarshal(data, &previousClassSet)
    if err != nil {
        fmt.Println("Error unmarshalling previous state data:", err)
        return previousClassSet
    }

    return previousClassSet
}

func saveCurrentState(classSet map[string]struct{}) {
    data, err := json.Marshal(classSet)
    if err != nil {
        fmt.Println("Error marshalling current state data:", err)
        return
    }

    err = os.WriteFile(stateFilePath, data, os.ModePerm)
    if err != nil {
        fmt.Println("Error writing current state file:", err)
    }
}

func createJavaScriptFile(classSet map[string]struct{}) {
    file, err := os.Create(jsFilePath)
    if err != nil {
        fmt.Println("Error creating JavaScript file:", err)
        return
    }
    defer file.Close()

    requiredClasses := make([]string, 0, len(classSet))
    for class := range classSet {
        requiredClasses = append(requiredClasses, class)
    }

    classesJSON, err := json.Marshal(requiredClasses)
    if err != nil {
        fmt.Println("Error marshalling classes to JSON:", err)
        return
    }

    jsContent := `document.addEventListener("DOMContentLoaded", function() {
    const requiredClasses = ` + string(classesJSON) + `;
    const missingClasses = [];

    requiredClasses.forEach(function(className) {
        const element = document.createElement("div");
        element.className = className;
        document.body.appendChild(element);
        const computedStyle = window.getComputedStyle(element);

        // Проверяем применены ли стили
        const properties = ["width", "height", "color", "background-color", "border", "padding", "margin"];
        let isMissing = properties.some(prop => {
            const value = computedStyle.getPropertyValue(prop);
            return !value || value === "auto" || value === "";
        });

        if (isMissing) {
            missingClasses.push(className);
        }
        document.body.removeChild(element);
    });

    if (missingClasses.length > 0) {
        console.error("Following classes are missing or not styled correctly:", missingClasses);
    } else {
        console.log("All required classes are styled correctly.");
    }
});`

    _, err = file.WriteString(jsContent)
    if err != nil {
        fmt.Println("Error writing to JavaScript file:", err)
    }
}

func collectClassesFromLibrary() map[string]struct{} {
    classSet := make(map[string]struct{})

    addClasses := func(classes map[string]string) {
        for class := range classes {
            classSet[class] = struct{}{}
        }
    }

    cssClasses := map[string]string{}
    
    library.AddTextColors(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddTypography(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddBackgroundColors(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddBorderStyles(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddEffects(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddLayout(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddBackgrounds(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddGradientColorStops(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddFlexboxGrid(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddBorder(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddSizing(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddSpacing(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddBorderColorStyles(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddEffectsColors(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddFilters(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddTables(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddTransitionsAndAnimations(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddTransforms(cssClasses)
    addClasses(cssClasses)
    
    cssClasses = map[string]string{}
    library.AddInteractivity(cssClasses)
    addClasses(cssClasses)

    return classSet
}

func collectCSSClasses() map[string]string {
    cssClasses := map[string]string{}

    library.AddTextColors(cssClasses)
    library.AddTypography(cssClasses)
    library.AddBackgroundColors(cssClasses)
    library.AddBorderStyles(cssClasses)
    library.AddEffects(cssClasses)
    library.AddLayout(cssClasses)
    library.AddBackgrounds(cssClasses)
    library.AddGradientColorStops(cssClasses)
    library.AddFlexboxGrid(cssClasses)
    library.AddBorder(cssClasses)
    library.AddSizing(cssClasses)
    library.AddSpacing(cssClasses)
    library.AddBorderColorStyles(cssClasses)
    library.AddEffectsColors(cssClasses)
    library.AddFilters(cssClasses)
    library.AddTables(cssClasses)
    library.AddTransitionsAndAnimations(cssClasses)
    library.AddTransforms(cssClasses)
    library.AddInteractivity(cssClasses)

    return cssClasses
}

func writeCSSFile(cssClasses map[string]string, classSet map[string]struct{}) error {
    file, err := os.Create(cssFilePath)
    if err != nil {
        return fmt.Errorf("error creating CSS file: %w", err)
    }
    defer file.Close()

    for class := range classSet {
        if style, exists := cssClasses[class]; exists {
            _, err := file.WriteString(fmt.Sprintf(".%s {%s}\n", class, style))
            if err != nil {
                return fmt.Errorf("error writing to CSS file: %w", err)
            }
        }
    }

    globalStyle := ""
    _, err = file.WriteString(globalStyle)
    if err != nil {
        return fmt.Errorf("error writing to CSS file: %w", err)
    }

    return nil
}






