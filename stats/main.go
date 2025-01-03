package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <directory_path>")
		os.Exit(1)
	}

	// Create configuration with default settings
	config := NewDefaultConfig()

	// Optionally add more terms to exclude
	config.AddExcludeTerms("customterm1", "customterm2")

	rootDir := os.Args[1]
	counter := NewTermCounter()

	// Walk through the directory
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			if err := analyzeFile(counter, path); err != nil {
				fmt.Printf("Warning: %v\n", err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Get and print top terms
	topTerms := getTopTerms(counter, config.MaxTerms)

	fmt.Printf("\nTop %d domain-specific terms in the codebase:\n", config.MaxTerms)
	fmt.Println("Term\t\tCount")
	fmt.Println("--------------------")
	for _, term := range topTerms {
		fmt.Printf("%-20s %d\n", term.Word, term.Count)
	}

	// Print some statistics
	fmt.Printf("\nAnalysis Statistics:\n")
	fmt.Printf("Total unique terms (excluding common Go terms): %d\n", len(counter.terms))
}
