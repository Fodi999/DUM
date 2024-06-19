package library

// AddBorderStyles добавляет стили границ в карту cssClasses
func AddBorderStyles(cssClasses map[string]string) {
    borderStyles := map[string]string{
        "rounded-none":        "border-radius: 0px;",
        "rounded-sm":          "border-radius: 0.125rem;", /* 2px */
        "rounded":             "border-radius: 0.25rem;", /* 4px */
        "rounded-md":          "border-radius: 0.375rem;", /* 6px */
        "rounded-lg":          "border-radius: 0.5rem;", /* 8px */
        "rounded-xl":          "border-radius: 0.75rem;", /* 12px */
        "rounded-2xl":         "border-radius: 1rem;", /* 16px */
        "rounded-3xl":         "border-radius: 1.5rem;", /* 24px */
        "rounded-full":        "border-radius: 9999px;",
        "rounded-s-none":      "border-start-start-radius: 0px; border-end-start-radius: 0px;",
        "rounded-s-sm":        "border-start-start-radius: 0.125rem; border-end-start-radius: 0.125rem;", /* 2px */
        "rounded-s":           "border-start-start-radius: 0.25rem; border-end-start-radius: 0.25rem;", /* 4px */
        "rounded-s-md":        "border-start-start-radius: 0.375rem; border-end-start-radius: 0.375rem;", /* 6px */
        "rounded-s-lg":        "border-start-start-radius: 0.5rem; border-end-start-radius: 0.5rem;", /* 8px */
        "rounded-s-xl":        "border-start-start-radius: 0.75rem; border-end-start-radius: 0.75rem;", /* 12px */
        "rounded-s-2xl":       "border-start-start-radius: 1rem; border-end-start-radius: 1rem;", /* 16px */
        "rounded-s-3xl":       "border-start-start-radius: 1.5rem; border-end-start-radius: 1.5rem;", /* 24px */
        "rounded-s-full":      "border-start-start-radius: 9999px; border-end-start-radius: 9999px;",
        "rounded-e-none":      "border-start-end-radius: 0px; border-end-end-radius: 0px;",
        "rounded-e-sm":        "border-start-end-radius: 0.125rem; border-end-end-radius: 0.125rem;", /* 2px */
        "rounded-e":           "border-start-end-radius: 0.25rem; border-end-end-radius: 0.25rem;", /* 4px */
        "rounded-e-md":        "border-start-end-radius: 0.375rem; border-end-end-radius: 0.375rem;", /* 6px */
        "rounded-e-lg":        "border-start-end-radius: 0.5rem; border-end-end-radius: 0.5rem;", /* 8px */
        "rounded-e-xl":        "border-start-end-radius: 0.75rem; border-end-end-radius: 0.75rem;", /* 12px */
        "rounded-e-2xl":       "border-start-end-radius: 1rem; border-end-end-radius: 1rem;", /* 16px */
        "rounded-e-3xl":       "border-start-end-radius: 1.5rem; border-end-end-radius: 1.5rem;", /* 24px */
        "rounded-e-full":      "border-start-end-radius: 9999px; border-end-end-radius: 9999px;",
        "rounded-t-none":      "border-top-left-radius: 0px; border-top-right-radius: 0px;",
        "rounded-t-sm":        "border-top-left-radius: 0.125rem; border-top-right-radius: 0.125rem;", /* 2px */
        "rounded-t":           "border-top-left-radius: 0.25rem; border-top-right-radius: 0.25rem;", /* 4px */
        "rounded-t-md":        "border-top-left-radius: 0.375rem; border-top-right-radius: 0.375rem;", /* 6px */
        "rounded-t-lg":        "border-top-left-radius: 0.5rem; border-top-right-radius: 0.5rem;", /* 8px */
        "rounded-t-xl":        "border-top-left-radius: 0.75rem; border-top-right-radius: 0.75rem;", /* 12px */
        "rounded-t-2xl":       "border-top-left-radius: 1rem; border-top-right-radius: 1rem;", /* 16px */
        "rounded-t-3xl":       "border-top-left-radius: 1.5rem; border-top-right-radius: 1.5rem;", /* 24px */
        "rounded-t-full":      "border-top-left-radius: 9999px; border-top-right-radius: 9999px;",
        "rounded-r-none":      "border-top-right-radius: 0px; border-bottom-right-radius: 0px;",
        "rounded-r-sm":        "border-top-right-radius: 0.125rem; border-bottom-right-radius: 0.125rem;", /* 2px */
        "rounded-r":           "border-top-right-radius: 0.25rem; border-bottom-right-radius: 0.25rem;", /* 4px */
        "rounded-r-md":        "border-top-right-radius: 0.375rem; border-bottom-right-radius: 0.375rem;", /* 6px */
        "rounded-r-lg":        "border-top-right-radius: 0.5rem; border-bottom-right-radius: 0.5rem;", /* 8px */
        "rounded-r-xl":        "border-top-right-radius: 0.75rem; border-bottom-right-radius: 0.75rem;", /* 12px */
        "rounded-r-2xl":       "border-top-right-radius: 1rem; border-bottom-right-radius: 1rem;", /* 16px */
        "rounded-r-3xl":       "border-top-right-radius: 1.5rem; border-bottom-right-radius: 1.5rem;", /* 24px */
        "rounded-r-full":      "border-top-right-radius: 9999px; border-bottom-right-radius: 9999px;",
        "rounded-b-none":      "border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;",
        "rounded-b-sm":        "border-bottom-right-radius: 0.125rem; border-bottom-left-radius: 0.125rem;", /* 2px */
        "rounded-b":           "border-bottom-right-radius: 0.25rem; border-bottom-left-radius: 0.25rem;", /* 4px */
        "rounded-b-md":        "border-bottom-right-radius: 0.375rem; border-bottom-left-radius: 0.375rem;", /* 6px */
        "rounded-b-lg":        "border-bottom-right-radius: 0.5rem; border-bottom-left-radius: 0.5rem;", /* 8px */
        "rounded-b-xl":        "border-bottom-right-radius: 0.75rem; border-bottom-left-radius: 0.75rem;", /* 12px */
        "rounded-b-2xl":       "border-bottom-right-radius: 1rem; border-bottom-left-radius: 1rem;", /* 16px */
        "rounded-b-3xl":       "border-bottom-right-radius: 1.5rem; border-bottom-left-radius: 1.5rem;", /* 24px */
        "rounded-b-full":      "border-bottom-right-radius: 9999px; border-bottom-left-radius: 9999px;",
        "rounded-l-none":      "border-top-left-radius: 0px; border-bottom-left-radius: 0px;",
        "rounded-l-sm":        "border-top-left-radius: 0.125rem; border-bottom-left-radius: 0.125rem;", /* 2px */
        "rounded-l":           "border-top-left-radius: 0.25rem; border-bottom-left-radius: 0.25rem;", /* 4px */
        "rounded-l-md":        "border-top-left-radius: 0.375rem; border-bottom-left-radius: 0.375rem;", /* 6px */
        "rounded-l-lg":        "border-top-left-radius: 0.5rem; border-bottom-left-radius: 0.5rem;", /* 8px */
        "rounded-l-xl":        "border-top-left-radius: 0.75rem; border-bottom-left-radius: 0.75rem;", /* 12px */
        "rounded-l-2xl":       "border-top-left-radius: 1rem; border-bottom-left-radius: 1rem;", /* 16px */
        "rounded-l-3xl":       "border-top-left-radius: 1.5rem; border-bottom-left-radius: 1.5rem;", /* 24px */
        "rounded-l-full":      "border-top-left-radius: 9999px; border-bottom-left-radius: 9999px;",
        "rounded-ss-none":     "border-start-start-radius: 0px;",
        "rounded-ss-sm":       "border-start-start-radius: 0.125rem;", /* 2px */
        "rounded-ss":          "border-start-start-radius: 0.25rem;", /* 4px */
        "rounded-ss-md":       "border-start-start-radius: 0.375rem;", /* 6px */
        "rounded-ss-lg":       "border-start-start-radius: 0.5rem;", /* 8px */
        "rounded-ss-xl":       "border-start-start-radius: 0.75rem;", /* 12px */
        "rounded-ss-2xl":      "border-start-start-radius: 1rem;", /* 16px */
        "rounded-ss-3xl":      "border-start-start-radius: 1.5rem;", /* 24px */
        "rounded-ss-full":     "border-start-start-radius: 9999px;",
        "rounded-se-none":     "border-start-end-radius: 0px;",
        "rounded-se-sm":       "border-start-end-radius: 0.125rem;", /* 2px */
        "rounded-se":          "border-start-end-radius: 0.25rem;", /* 4px */
        "rounded-se-md":       "border-start-end-radius: 0.375rem;", /* 6px */
        "rounded-se-lg":       "border-start-end-radius: 0.5rem;", /* 8px */
        "rounded-se-xl":       "border-start-end-radius: 0.75rem;", /* 12px */
        "rounded-se-2xl":      "border-start-end-radius: 1rem;", /* 16px */
        "rounded-se-3xl":      "border-start-end-radius: 1.5rem;", /* 24px */
        "rounded-se-full":     "border-start-end-radius: 9999px;",
        "rounded-ee-none":     "border-end-end-radius: 0px;",
        "rounded-ee-sm":       "border-end-end-radius: 0.125rem;", /* 2px */
        "rounded-ee":          "border-end-end-radius: 0.25rem;", /* 4px */
        "rounded-ee-md":       "border-end-end-radius: 0.375rem;", /* 6px */
        "rounded-ee-lg":       "border-end-end-radius: 0.5rem;", /* 8px */
        "rounded-ee-xl":       "border-end-end-radius: 0.75rem;", /* 12px */
        "rounded-ee-2xl":      "border-end-end-radius: 1rem;", /* 16px */
        "rounded-ee-3xl":      "border-end-end-radius: 1.5rem;", /* 24px */
        "rounded-ee-full":     "border-end-end-radius: 9999px;",
        "rounded-es-none":     "border-end-start-radius: 0px;",
        "rounded-es-sm":       "border-end-start-radius: 0.125rem;", /* 2px */
        "rounded-es":          "border-end-start-radius: 0.25rem;", /* 4px */
        "rounded-es-md":       "border-end-start-radius: 0.375rem;", /* 6px */
        "rounded-es-lg":       "border-end-start-radius: 0.5rem;", /* 8px */
        "rounded-es-xl":       "border-end-start-radius: 0.75rem;", /* 12px */
        "rounded-es-2xl":      "border-end-start-radius: 1rem;", /* 16px */
        "rounded-es-3xl":      "border-end-start-radius: 1.5rem;", /* 24px */
        "rounded-es-full":     "border-end-start-radius: 9999px;",
        "rounded-tl-none":     "border-top-left-radius: 0px;",
        "rounded-tl-sm":       "border-top-left-radius: 0.125rem;", /* 2px */
        "rounded-tl":          "border-top-left-radius: 0.25rem;", /* 4px */
        "rounded-tl-md":       "border-top-left-radius: 0.375rem;", /* 6px */
        "rounded-tl-lg":       "border-top-left-radius: 0.5rem;", /* 8px */
        "rounded-tl-xl":       "border-top-left-radius: 0.75rem;", /* 12px */
        "rounded-tl-2xl":      "border-top-left-radius: 1rem;", /* 16px */
        "rounded-tl-3xl":      "border-top-left-radius: 1.5rem;", /* 24px */
        "rounded-tl-full":     "border-top-left-radius: 9999px;",
        "rounded-tr-none":     "border-top-right-radius: 0px;",
        "rounded-tr-sm":       "border-top-right-radius: 0.125rem;", /* 2px */
        "rounded-tr":          "border-top-right-radius: 0.25rem;", /* 4px */
        "rounded-tr-md":       "border-top-right-radius: 0.375rem;", /* 6px */
        "rounded-tr-lg":       "border-top-right-radius: 0.5rem;", /* 8px */
        "rounded-tr-xl":       "border-top-right-radius: 0.75rem;", /* 12px */
        "rounded-tr-2xl":      "border-top-right-radius: 1rem;", /* 16px */
        "rounded-tr-3xl":      "border-top-right-radius: 1.5rem;", /* 24px */
        "rounded-tr-full":     "border-top-right-radius: 9999px;",
        "rounded-br-none":     "border-bottom-right-radius: 0px;",
        "rounded-br-sm":       "border-bottom-right-radius: 0.125rem;", /* 2px */
        "rounded-br":          "border-bottom-right-radius: 0.25rem;", /* 4px */
        "rounded-br-md":       "border-bottom-right-radius: 0.375rem;", /* 6px */
        "rounded-br-lg":       "border-bottom-right-radius: 0.5rem;", /* 8px */
        "rounded-br-xl":       "border-bottom-right-radius: 0.75rem;", /* 12px */
        "rounded-br-2xl":      "border-bottom-right-radius: 1rem;", /* 16px */
        "rounded-br-3xl":      "border-bottom-right-radius: 1.5rem;", /* 24px */
        "rounded-br-full":     "border-bottom-right-radius: 9999px;",
        "rounded-bl-none":     "border-bottom-left-radius: 0px;",
        "rounded-bl-sm":       "border-bottom-left-radius: 0.125rem;", /* 2px */
        "rounded-bl":          "border-bottom-left-radius: 0.25rem;", /* 4px */
        "rounded-bl-md":       "border-bottom-left-radius: 0.375rem;", /* 6px */
        "rounded-bl-lg":       "border-bottom-left-radius: 0.5rem;", /* 8px */
        "rounded-bl-xl":       "border-bottom-left-radius: 0.75rem;", /* 12px */
        "rounded-bl-2xl":      "border-bottom-left-radius: 1rem;", /* 16px */
        "rounded-bl-3xl":      "border-bottom-left-radius: 1.5rem;", /* 24px */
        "rounded-bl-full":     "border-bottom-left-radius: 9999px;",
        "border-0":            "border-width: 0px;",
        "border-2":            "border-width: 2px;",
        "border-4":            "border-width: 4px;",
        "border-8":            "border-width: 8px;",
        "border":              "border-width: 1px;",
        "border-x-0":          "border-left-width: 0px; border-right-width: 0px;",
        "border-x-2":          "border-left-width: 2px; border-right-width: 2px;",
        "border-x-4":          "border-left-width: 4px; border-right-width: 4px;",
        "border-x-8":          "border-left-width: 8px; border-right-width: 8px;",
        "border-x":            "border-left-width: 1px; border-right-width: 1px;",
        "border-y-0":          "border-top-width: 0px; border-bottom-width: 0px;",
        "border-y-2":          "border-top-width: 2px; border-bottom-width: 2px;",
        "border-y-4":          "border-top-width: 4px; border-bottom-width: 4px;",
        "border-y-8":          "border-top-width: 8px; border-bottom-width: 8px;",
        "border-y":            "border-top-width: 1px; border-bottom-width: 1px;",
        "border-s-0":          "border-inline-start-width: 0px;",
        "border-s-2":          "border-inline-start-width: 2px;",
        "border-s-4":          "border-inline-start-width: 4px;",
        "border-s-8":          "border-inline-start-width: 8px;",
        "border-s":            "border-inline-start-width: 1px;",
        "border-e-0":          "border-inline-end-width: 0px;",
        "border-e-2":          "border-inline-end-width: 2px;",
        "border-e-4":          "border-inline-end-width: 4px;",
        "border-e-8":          "border-inline-end-width: 8px;",
        "border-e":            "border-inline-end-width: 1px;",
        "border-t-0":          "border-top-width: 0px;",
        "border-t-2":          "border-top-width: 2px;",
        "border-t-4":          "border-top-width: 4px;",
        "border-t-8":          "border-top-width: 8px;",
        "border-t":            "border-top-width: 1px;",
        "border-r-0":          "border-right-width: 0px;",
        "border-r-2":          "border-right-width: 2px;",
        "border-r-4":          "border-right-width: 4px;",
        "border-r-8":          "border-right-width: 8px;",
        "border-r":            "border-right-width: 1px;",
        "border-b-0":          "border-bottom-width: 0px;",
        "border-b-2":          "border-bottom-width: 2px;",
        "border-b-4":          "border-bottom-width: 4px;",
        "border-b-8":          "border-bottom-width: 8px;",
        "border-b":            "border-bottom-width: 1px;",
        "border-l-0":          "border-left-width: 0px;",
        "border-l-2":          "border-left-width: 2px;",
        "border-l-4":          "border-left-width: 4px;",
        "border-l-8":          "border-left-width: 8px;",
        "border-l":            "border-left-width: 1px;",
        "border-solid":        "border-style: solid;",
        "border-dashed":       "border-style: dashed;",
        "border-dotted":       "border-style: dotted;",
        "border-double":       "border-style: double;",
        "border-hidden":       "border-style: hidden;",
        "border-none":         "border-style: none;",
        "divide-x-0":          "border-right-width: 0px; border-left-width: 0px;",
        "divide-x-2":          "border-right-width: 0px; border-left-width: 2px;",
        "divide-x-4":          "border-right-width: 0px; border-left-width: 4px;",
        "divide-x-8":          "border-right-width: 0px; border-left-width: 8px;",
        "divide-x":            "border-right-width: 0px; border-left-width: 1px;",
        "divide-y-0":          "border-top-width: 0px; border-bottom-width: 0px;",
        "divide-y-2":          "border-top-width: 2px; border-bottom-width: 0px;",
        "divide-y-4":          "border-top-width: 4px; border-bottom-width: 0px;",
        "divide-y-8":          "border-top-width: 8px; border-bottom-width: 0px;",
        "divide-y":            "border-top-width: 1px; border-bottom-width: 0px;",
        "divide-y-reverse":    "--tw-divide-y-reverse: 1;",
        "divide-x-reverse":    "--tw-divide-x-reverse: 1;",
        "divide-solid":        "border-style: solid;",
        "divide-dashed":       "border-style: dashed;",
        "divide-dotted":       "border-style: dotted;",
        "divide-double":       "border-style: double;",
        "divide-none":         "border-style: none;",
        "outline-0":           "outline-width: 0px;",
        "outline-1":           "outline-width: 1px;",
        "outline-2":           "outline-width: 2px;",
        "outline-4":           "outline-width: 4px;",
        "outline-8":           "outline-width: 8px;",
        "outline-none":        "outline: 2px solid transparent; outline-offset: 2px;",
        "outline":             "outline-style: solid;",
        "outline-dashed":      "outline-style: dashed;",
        "outline-dotted":      "outline-style: dotted;",
        "outline-double":      "outline-style: double;",
        "outline-offset-0":    "outline-offset: 0px;",
        "outline-offset-1":    "outline-offset: 1px;",
        "outline-offset-2":    "outline-offset: 2px;",
        "outline-offset-4":    "outline-offset: 4px;",
        "outline-offset-8":    "outline-offset: 8px;",
        "ring-0":              "box-shadow: var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring-1":              "box-shadow: var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring-2":              "box-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring":                "box-shadow: var(--tw-ring-inset) 0 0 0 calc(3px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring-4":              "box-shadow: var(--tw-ring-inset) 0 0 0 calc(4px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring-8":              "box-shadow: var(--tw-ring-inset) 0 0 0 calc(8px + var(--tw-ring-offset-width)) var(--tw-ring-color);",
        "ring-inset":          "--tw-ring-inset: inset;",
        "ring-offset-0":       "--tw-ring-offset-width: 0px; box-shadow: 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color), var(--tw-ring-shadow);",
        "ring-offset-1":       "--tw-ring-offset-width: 1px; box-shadow: 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color), var(--tw-ring-shadow);",
        "ring-offset-2":       "--tw-ring-offset-width: 2px; box-shadow: 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color), var(--tw-ring-shadow);",
        "ring-offset-4":       "--tw-ring-offset-width: 4px; box-shadow: 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color), var(--tw-ring-shadow);",
        "ring-offset-8":       "--tw-ring-offset-width: 8px; box-shadow: 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color), var(--tw-ring-shadow);",
    }

    for key, value := range borderStyles {
        cssClasses[key] = value
    }
}
