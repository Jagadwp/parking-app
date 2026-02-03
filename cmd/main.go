package main

import (
	"bufio"
	"fmt"
	"os"

	"parking-app/internal/app"
	"parking-app/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: parking-app <input_file>")
		fmt.Println("Example: parking-app input.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	if err := executeFromFile(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func executeFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file '%s': %w", filename, err)
	}
	defer file.Close()

	executor := app.NewExecutor()
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		cmd, err := parser.ParseCommand(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on line %d: %v\n", lineNumber, err)
			continue
		}

		if cmd == nil {
			continue
		}

		if err := executor.ExecuteCommand(cmd); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command on line %d: %v\n", lineNumber, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}
