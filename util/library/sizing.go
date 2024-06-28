//library/sizing.go
package library

// AddSizing добавляет все классы размеров в переданную карту CSS классов.
func AddSizing(cssClasses map[string]string) {
	sizing := map[string]string{
		// Width
		"w-0":         "width: 0px;",
		"w-px":        "width: 1px;",
		"w-0.5":       "width: 0.125rem;", // 2px
		"w-1":         "width: 0.25rem;",  // 4px
		"w-1.5":       "width: 0.375rem;", // 6px
		"w-2":         "width: 0.5rem;",   // 8px
		"w-2.5":       "width: 0.625rem;", // 10px
		"w-3":         "width: 0.75rem;",  // 12px
		"w-3.5":       "width: 0.875rem;", // 14px
		"w-4":         "width: 1rem;",     // 16px
		"w-5":         "width: 1.25rem;",  // 20px
		"w-6":         "width: 1.5rem;",   // 24px
		"w-7":         "width: 1.75rem;",  // 28px
		"w-8":         "width: 2rem;",     // 32px
		"w-9":         "width: 2.25rem;",  // 36px
		"w-10":        "width: 2.5rem;",   // 40px
		"w-11":        "width: 2.75rem;",  // 44px
		"w-12":        "width: 3rem;",     // 48px
		"w-14":        "width: 3.5rem;",   // 56px
		"w-16":        "width: 4rem;",     // 64px
		"w-20":        "width: 5rem;",     // 80px
		"w-24":        "width: 6rem;",     // 96px
		"w-28":        "width: 7rem;",     // 112px
		"w-32":        "width: 8rem;",     // 128px
		"w-36":        "width: 9rem;",     // 144px
		"w-40":        "width: 10rem;",    // 160px
		"w-44":        "width: 11rem;",    // 176px
		"w-48":        "width: 12rem;",    // 192px
		"w-52":        "width: 13rem;",    // 208px
		"w-56":        "width: 14rem;",    // 224px
		"w-60":        "width: 15rem;",    // 240px
		"w-64":        "width: 16rem;",    // 256px
		"w-72":        "width: 18rem;",    // 288px
		"w-80":        "width: 20rem;",    // 320px
		"w-96":        "width: 24rem;",    // 384px
		"w-auto":      "width: auto;",
		"w-1/2":       "width: 50%;",
		"w-1/3":       "width: 33.333333%;",
		"w-2/3":       "width: 66.666667%;",
		"w-1/4":       "width: 25%;",
		"w-2/4":       "width: 50%;",
		"w-3/4":       "width: 75%;",
		"w-1/5":       "width: 20%;",
		"w-2/5":       "width: 40%;",
		"w-3/5":       "width: 60%;",
		"w-4/5":       "width: 80%;",
		"w-1/6":       "width: 16.666667%;",
		"w-2/6":       "width: 33.333333%;",
		"w-3/6":       "width: 50%;",
		"w-4/6":       "width: 66.666667%;",
		"w-5/6":       "width: 83.333333%;",
		"w-1/12":      "width: 8.333333%;",
		"w-2/12":      "width: 16.666667%;",
		"w-3/12":      "width: 25%;",
		"w-4/12":      "width: 33.333333%;",
		"w-5/12":      "width: 41.666667%;",
		"w-6/12":      "width: 50%;",
		"w-7/12":      "width: 58.333333%;",
		"w-8/12":      "width: 66.666667%;",
		"w-9/12":      "width: 75%;",
		"w-10/12":     "width: 83.333333%;",
		"w-11/12":     "width: 91.666667%;",
		"w-full":      "width: 100%;",
		"w-screen":    "width: 100vw;",
		"w-svw":       "width: 100svw;",
		"w-lvw":       "width: 100lvw;",
		"w-dvw":       "width: 100dvw;",
		"w-min":       "width: min-content;",
		"w-max":       "width: max-content;",
		"w-fit":       "width: fit-content;",

		// Min-Width
		"min-w-0":       "min-width: 0px;",
		"min-w-1":       "min-width: 0.25rem;", // 4px
		"min-w-2":       "min-width: 0.5rem;",  // 8px
		"min-w-3":       "min-width: 0.75rem;", // 12px
		"min-w-4":       "min-width: 1rem;",    // 16px
		"min-w-5":       "min-width: 1.25rem;", // 20px
		"min-w-6":       "min-width: 1.5rem;",  // 24px
		"min-w-7":       "min-width: 1.75rem;", // 28px
		"min-w-8":       "min-width: 2rem;",    // 32px
		"min-w-9":       "min-width: 2.25rem;", // 36px
		"min-w-10":      "min-width: 2.5rem;",  // 40px
		"min-w-11":      "min-width: 2.75rem;", // 44px
		"min-w-12":      "min-width: 3rem;",    // 48px
		"min-w-14":      "min-width: 3.5rem;",  // 56px
		"min-w-16":      "min-width: 4rem;",    // 64px
		"min-w-20":      "min-width: 5rem;",    // 80px
		"min-w-24":      "min-width: 6rem;",    // 96px
		"min-w-28":      "min-width: 7rem;",    // 112px
		"min-w-32":      "min-width: 8rem;",    // 128px
		"min-w-36":      "min-width: 9rem;",    // 144px
		"min-w-40":      "min-width: 10rem;",   // 160px
		"min-w-44":      "min-width: 11rem;",   // 176px
		"min-w-48":      "min-width: 12rem;",   // 192px
		"min-w-52":      "min-width: 13rem;",   // 208px
		"min-w-56":      "min-width: 14rem;",   // 224px
		"min-w-60":      "min-width: 15rem;",   // 240px
		"min-w-64":      "min-width: 16rem;",   // 256px
		"min-w-72":      "min-width: 18rem;",   // 288px
		"min-w-80":      "min-width: 20rem;",   // 320px
		"min-w-96":      "min-width: 24rem;",   // 384px
		"min-w-px":      "min-width: 1px;",
		"min-w-0.5":     "min-width: 0.125rem;", // 2px
		"min-w-1.5":     "min-width: 0.375rem;", // 6px
		"min-w-2.5":     "min-width: 0.625rem;", // 10px
		"min-w-3.5":     "min-width: 0.875rem;", // 14px
		"min-w-full":    "min-width: 100%;",
		"min-w-min":     "min-width: min-content;",
		"min-w-max":     "min-width: max-content;",
		"min-w-fit":     "min-width: fit-content;",

		// Max-Width
		"max-w-0":          "max-width: 0px;",
		"max-w-px":         "max-width: 1px;",
		"max-w-0.5":        "max-width: 0.125rem;", // 2px
		"max-w-1":          "max-width: 0.25rem;",  // 4px
		"max-w-1.5":        "max-width: 0.375rem;", // 6px
		"max-w-2":          "max-width: 0.5rem;",   // 8px
		"max-w-2.5":        "max-width: 0.625rem;", // 10px
		"max-w-3":          "max-width: 0.75rem;",  // 12px
		"max-w-3.5":        "max-width: 0.875rem;", // 14px
		"max-w-4":          "max-width: 1rem;",     // 16px
		"max-w-5":          "max-width: 1.25rem;",  // 20px
		"max-w-6":          "max-width: 1.5rem;",   // 24px
		"max-w-7":          "max-width: 1.75rem;",  // 28px
		"max-w-8":          "max-width: 2rem;",     // 32px
		"max-w-9":          "max-width: 2.25rem;",  // 36px
		"max-w-10":         "max-width: 2.5rem;",   // 40px
		"max-w-11":         "max-width: 2.75rem;",  // 44px
		"max-w-12":         "max-width: 3rem;",     // 48px
		"max-w-14":         "max-width: 3.5rem;",   // 56px
		"max-w-16":         "max-width: 4rem;",     // 64px
		"max-w-20":         "max-width: 5rem;",     // 80px
		"max-w-24":         "max-width: 6rem;",     // 96px
		"max-w-28":         "max-width: 7rem;",     // 112px
		"max-w-32":         "max-width: 8rem;",     // 128px
		"max-w-36":         "max-width: 9rem;",     // 144px
		"max-w-40":         "max-width: 10rem;",    // 160px
		"max-w-44":         "max-width: 11rem;",    // 176px
		"max-w-48":         "max-width: 12rem;",    // 192px
		"max-w-52":         "max-width: 13rem;",    // 208px
		"max-w-56":         "max-width: 14rem;",    // 224px
		"max-w-60":         "max-width: 15rem;",    // 240px
		"max-w-64":         "max-width: 16rem;",    // 256px
		"max-w-72":         "max-width: 18rem;",    // 288px
		"max-w-80":         "max-width: 20rem;",    // 320px
		"max-w-96":         "max-width: 24rem;",    // 384px
		"max-w-none":       "max-width: none;",
		"max-w-xs":         "max-width: 20rem;",    // 320px
		"max-w-sm":         "max-width: 24rem;",    // 384px
		"max-w-md":         "max-width: 28rem;",    // 448px
		"max-w-lg":         "max-width: 32rem;",    // 512px
		"max-w-xl":         "max-width: 36rem;",    // 576px
		"max-w-2xl":        "max-width: 42rem;",    // 672px
		"max-w-3xl":        "max-width: 48rem;",    // 768px
		"max-w-4xl":        "max-width: 56rem;",    // 896px
		"max-w-5xl":        "max-width: 64rem;",    // 1024px
		"max-w-6xl":        "max-width: 72rem;",    // 1152px
		"max-w-7xl":        "max-width: 80rem;",    // 1280px
		"max-w-full":       "max-width: 100%;",
		"max-w-min":        "max-width: min-content;",
		"max-w-max":        "max-width: max-content;",
		"max-w-fit":        "max-width: fit-content;",
		"max-w-prose":      "max-width: 65ch;",
		"max-w-screen-sm":  "max-width: 640px;",
		"max-w-screen-md":  "max-width: 768px;",
		"max-w-screen-lg":  "max-width: 1024px;",
		"max-w-screen-xl":  "max-width: 1280px;",
		"max-w-screen-2xl": "max-width: 1536px;",

		// Height
		"h-0":         "height: 0px;",
		"h-px":        "height: 1px;",
		"h-0.5":       "height: 0.125rem;", // 2px
		"h-1":         "height: 0.25rem;",  // 4px
		"h-1.5":       "height: 0.375rem;", // 6px
		"h-2":         "height: 0.5rem;",   // 8px
		"h-2.5":       "height: 0.625rem;", // 10px
		"h-3":         "height: 0.75rem;",  // 12px
		"h-3.5":       "height: 0.875rem;", // 14px
		"h-4":         "height: 1rem;",     // 16px
		"h-5":         "height: 1.25rem;",  // 20px
		"h-6":         "height: 1.5rem;",   // 24px
		"h-7":         "height: 1.75rem;",  // 28px
		"h-8":         "height: 2rem;",     // 32px
		"h-9":         "height: 2.25rem;",  // 36px
		"h-10":        "height: 2.5rem;",   // 40px
		"h-11":        "height: 2.75rem;",  // 44px
		"h-12":        "height: 3rem;",     // 48px
		"h-14":        "height: 3.5rem;",   // 56px
		"h-16":        "height: 4rem;",     // 64px
		"h-20":        "height: 5rem;",     // 80px
		"h-24":        "height: 6rem;",     // 96px
		"h-28":        "height: 7rem;",     // 112px
		"h-32":        "height: 8rem;",     // 128px
		"h-36":        "height: 9rem;",     // 144px
		"h-40":        "height: 10rem;",    // 160px
		"h-44":        "height: 11rem;",    // 176px
		"h-48":        "height: 12rem;",    // 192px
		"h-52":        "height: 13rem;",    // 208px
		"h-56":        "height: 14rem;",    // 224px
		"h-60":        "height: 15rem;",    // 240px
		"h-64":        "height: 16rem;",    // 256px
		"h-72":        "height: 18rem;",    // 288px
		"h-80":        "height: 20rem;",    // 320px
		"h-96":        "height: 24rem;",    // 384px
		"h-auto":      "height: auto;",
		"h-1/2":       "height: 50%;",
		"h-1/3":       "height: 33.333333%;",
		"h-2/3":       "height: 66.666667%;",
		"h-1/4":       "height: 25%;",
		"h-2/4":       "height: 50%;",
		"h-3/4":       "height: 75%;",
		"h-1/5":       "height: 20%;",
		"h-2/5":       "height: 40%;",
		"h-3/5":       "height: 60%;",
		"h-4/5":       "height: 80%;",
		"h-1/6":       "height: 16.666667%;",
		"h-2/6":       "height: 33.333333%;",
		"h-3/6":       "height: 50%;",
		"h-4/6":       "height: 66.666667%;",
		"h-5/6":       "height: 83.333333%;",
		"h-full":      "height: 100%;",
		"h-screen":    "height: 100vh;",
		"h-svh":       "height: 100svh;",
		"h-lvh":       "height: 100lvh;",
		"h-dvh":       "height: 100dvh;",
		"h-min":       "height: min-content;",
		"h-max":       "height: max-content;",
		"h-fit":       "height: fit-content;",

		// Min-Height
		"min-h-0":       "min-height: 0px;",
		"min-h-1":       "min-height: 0.25rem;", // 4px
		"min-h-2":       "min-height: 0.5rem;",  // 8px
		"min-h-3":       "min-height: 0.75rem;", // 12px
		"min-h-4":       "min-height: 1rem;",    // 16px
		"min-h-5":       "min-height: 1.25rem;", // 20px
		"min-h-6":       "min-height: 1.5rem;",  // 24px
		"min-h-7":       "min-height: 1.75rem;", // 28px
		"min-h-8":       "min-height: 2rem;",    // 32px
		"min-h-9":       "min-height: 2.25rem;", // 36px
		"min-h-10":      "min-height: 2.5rem;",  // 40px
		"min-h-11":      "min-height: 2.75rem;", // 44px
		"min-h-12":      "min-height: 3rem;",    // 48px
		"min-h-14":      "min-height: 3.5rem;",  // 56px
		"min-h-16":      "min-height: 4rem;",    // 64px
		"min-h-20":      "min-height: 5rem;",    // 80px
		"min-h-24":      "min-height: 6rem;",    // 96px
		"min-h-28":      "min-height: 7rem;",    // 112px
		"min-h-32":      "min-height: 8rem;",    // 128px
		"min-h-36":      "min-height: 9rem;",    // 144px
		"min-h-40":      "min-height: 10rem;",   // 160px
		"min-h-44":      "min-height: 11rem;",   // 176px
		"min-h-48":      "min-height: 12rem;",   // 192px
		"min-h-52":      "min-height: 13rem;",   // 208px
		"min-h-56":      "min-height: 14rem;",   // 224px
		"min-h-60":      "min-height: 15rem;",   // 240px
		"min-h-64":      "min-height: 16rem;",   // 256px
		"min-h-72":      "min-height: 18rem;",   // 288px
		"min-h-80":      "min-height: 20rem;",   // 320px
		"min-h-96":      "min-height: 24rem;",   // 384px
		"min-h-px":      "min-height: 1px;",
		"min-h-0.5":     "min-height: 0.125rem;", // 2px
		"min-h-1.5":     "min-height: 0.375rem;", // 6px
		"min-h-2.5":     "min-height: 0.625rem;", // 10px
		"min-h-3.5":     "min-height: 0.875rem;", // 14px
		"min-h-full":    "min-height: 100%;",
		"min-h-screen":  "min-height: 100vh;",
		"min-h-svh":     "min-height: 100svh;",
		"min-h-lvh":     "min-height: 100lvh;",
		"min-h-dvh":     "min-height: 100dvh;",
		"min-h-min":     "min-height: min-content;",
		"min-h-max":     "min-height: max-content;",
		"min-h-fit":     "min-height: fit-content;",

		// Max-Height
		"max-h-0":          "max-height: 0px;",
		"max-h-px":         "max-height: 1px;",
		"max-h-0.5":        "max-height: 0.125rem;", // 2px
		"max-h-1":          "max-height: 0.25rem;",  // 4px
		"max-h-1.5":        "max-height: 0.375rem;", // 6px
		"max-h-2":          "max-height: 0.5rem;",   // 8px
		"max-h-2.5":        "max-height: 0.625rem;", // 10px
		"max-h-3":          "max-height: 0.75rem;",  // 12px
		"max-h-3.5":        "max-height: 0.875rem;", // 14px
		"max-h-4":          "max-height: 1rem;",     // 16px
		"max-h-5":          "max-height: 1.25rem;",  // 20px
		"max-h-6":          "max-height: 1.5rem;",   // 24px
		"max-h-7":          "max-height: 1.75rem;",  // 28px
		"max-h-8":          "max-height: 2rem;",     // 32px
		"max-h-9":          "max-height: 2.25rem;",  // 36px
		"max-h-10":         "max-height: 2.5rem;",   // 40px
		"max-h-11":         "max-height: 2.75rem;",  // 44px
		"max-h-12":         "max-height: 3rem;",     // 48px
		"max-h-14":         "max-height: 3.5rem;",   // 56px
		"max-h-16":         "max-height: 4rem;",     // 64px
		"max-h-20":         "max-height: 5rem;",     // 80px
		"max-h-24":         "max-height: 6rem;",     // 96px
		"max-h-28":         "max-height: 7rem;",     // 112px
		"max-h-32":         "max-height: 8rem;",     // 128px
		"max-h-36":         "max-height: 9rem;",     // 144px
		"max-h-40":         "max-height: 10rem;",    // 160px
		"max-h-44":         "max-height: 11rem;",    // 176px
		"max-h-48":         "max-height: 12rem;",    // 192px
		"max-h-52":         "max-height: 13rem;",    // 208px
		"max-h-56":         "max-height: 14rem;",    // 224px
		"max-h-60":         "max-height: 15rem;",    // 240px
		"max-h-64":         "max-height: 16rem;",    // 256px
		"max-h-72":         "max-height: 18rem;",    // 288px
		"max-h-80":         "max-height: 20rem;",    // 320px
		"max-h-96":         "max-height: 24rem;",    // 384px
		"max-h-none":       "max-height: none;",
		"max-h-full":       "max-height: 100%;",
		"max-h-screen":     "max-height: 100vh;",
		"max-h-svh":        "max-height: 100svh;",
		"max-h-lvh":        "max-height: 100lvh;",
		"max-h-dvh":        "max-height: 100dvh;",
		"max-h-min":        "max-height: min-content;",
		"max-h-max":        "max-height: max-content;",
		"max-h-fit":        "max-height: fit-content;",

		// Size
		"size-0":           "width: 0px; height: 0px;",
		"size-px":          "width: 1px; height: 1px;",
		"size-0.5":         "width: 0.125rem; height: 0.125rem;", // 2px
		"size-1":           "width: 0.25rem; height: 0.25rem;",   // 4px
		"size-1.5":         "width: 0.375rem; height: 0.375rem;", // 6px
		"size-2":           "width: 0.5rem; height: 0.5rem;",     // 8px
		"size-2.5":         "width: 0.625rem; height: 0.625rem;", // 10px
		"size-3":           "width: 0.75rem; height: 0.75rem;",   // 12px
		"size-3.5":         "width: 0.875rem; height: 0.875rem;", // 14px
		"size-4":           "width: 1rem; height: 1rem;",         // 16px
		"size-5":           "width: 1.25rem; height: 1.25rem;",   // 20px
		"size-6":           "width: 1.5rem; height: 1.5rem;",     // 24px
		"size-7":           "width: 1.75rem; height: 1.75rem;",   // 28px
		"size-8":           "width: 2rem; height: 2rem;",         // 32px
		"size-9":           "width: 2.25rem; height: 2.25rem;",   // 36px
		"size-10":          "width: 2.5rem; height: 2.5rem;",     // 40px
		"size-11":          "width: 2.75rem; height: 2.75rem;",   // 44px
		"size-12":          "width: 3rem; height: 3rem;",         // 48px
		"size-14":          "width: 3.5rem; height: 3.5rem;",     // 56px
		"size-16":          "width: 4rem; height: 4rem;",         // 64px
		"size-20":          "width: 5rem; height: 5rem;",         // 80px
		"size-24":          "width: 6rem; height: 6rem;",         // 96px
		"size-28":          "width: 7rem; height: 7rem;",         // 112px
		"size-32":          "width: 8rem; height: 8rem;",         // 128px
		"size-36":          "width: 9rem; height: 9rem;",         // 144px
		"size-40":          "width: 10rem; height: 10rem;",       // 160px
		"size-44":          "width: 11rem; height: 11rem;",       // 176px
		"size-48":          "width: 12rem; height: 12rem;",       // 192px
		"size-52":          "width: 13rem; height: 13rem;",       // 208px
		"size-56":          "width: 14rem; height: 14rem;",       // 224px
		"size-60":          "width: 15rem; height: 15rem;",       // 240px
		"size-64":          "width: 16rem; height: 16rem;",       // 256px
		"size-72":          "width: 18rem; height: 18rem;",       // 288px
		"size-80":          "width: 20rem; height: 20rem;",       // 320px
		"size-96":          "width: 24rem; height: 24rem;",       // 384px
		"size-auto":        "width: auto; height: auto;",
		"size-1/2":         "width: 50%; height: 50%;",
		"size-1/3":         "width: 33.333333%; height: 33.333333%;",
		"size-2/3":         "width: 66.666667%; height: 66.666667%;",
		"size-1/4":         "width: 25%; height: 25%;",
		"size-2/4":         "width: 50%; height: 50%;",
		"size-3/4":         "width: 75%; height: 75%;",
		"size-1/5":         "width: 20%; height: 20%;",
		"size-2/5":         "width: 40%; height: 40%;",
		"size-3/5":         "width: 60%; height: 60%;",
		"size-4/5":         "width: 80%; height: 80%;",
		"size-1/6":         "width: 16.666667%; height: 16.666667%;",
		"size-2/6":         "width: 33.333333%; height: 33.333333%;",
		"size-3/6":         "width: 50%; height: 50%;",
		"size-4/6":         "width: 66.666667%; height: 66.666667%;",
		"size-5/6":         "width: 83.333333%; height: 83.333333%;",
		"size-1/12":        "width: 8.333333%; height: 8.333333%;",
		"size-2/12":        "width: 16.666667%; height: 16.666667%;",
		"size-3/12":        "width: 25%; height: 25%;",
		"size-4/12":        "width: 33.333333%; height: 33.333333%;",
		"size-5/12":        "width: 41.666667%; height: 41.666667%;",
		"size-6/12":        "width: 50%; height: 50%;",
		"size-7/12":        "width: 58.333333%; height: 58.333333%;",
		"size-8/12":        "width: 66.666667%; height: 66.666667%;",
		"size-9/12":        "width: 75%; height: 75%;",
		"size-10/12":       "width: 83.333333%; height: 83.333333%;",
		"size-11/12":       "width: 91.666667%; height: 91.666667%;",
		"size-full":        "width: 100%; height: 100%;",
		"size-min":         "width: min-content; height: min-content;",
		"size-max":         "width: max-content; height: max-content;",
		"size-fit":         "width: fit-content; height: fit-content;",
	}

	for key, value := range sizing {
		cssClasses[key] = value
	}
}


