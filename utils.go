package main

import (
	"os"
)

// AppendToFile appends the given text to the specified file.
// If the file doesn't exist, it will be created.
func AppendToFile(text string) {
	// Open the file in append mode, create if not exists, with write permission
	f, err := os.OpenFile("nothing.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// return fmt.Errorf("could not open file: %w", err)
		return
	}
	defer f.Close()

	// Write the text to the file
	if _, err := f.WriteString(text); err != nil {
		// return fmt.Errorf("could not write to file: %w", err)
		return
	}

	return
}
