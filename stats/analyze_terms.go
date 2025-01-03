package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"unicode"
)

// Term represents a term and its frequency in the codebase
type Term struct {
	Word  string
	Count int
}

// TermCounter keeps track of term frequencies
type TermCounter struct {
	terms map[string]int
}

// NewTermCounter creates a new TermCounter
func NewTermCounter() *TermCounter {
	return &TermCounter{
		terms: make(map[string]int),
	}
}

// Add increments the count for a term
func (tc *TermCounter) Add(term string) {
	// Convert to lowercase
	term = strings.ToLower(term)

	// Skip short terms and common Go terms
	if len(term) <= 2 || IsCommonGoTerm(term) {
		return
	}

	// Skip terms that are just numbers
	if isNumeric(term) {
		return
	}

	tc.terms[term]++
}

// Helper function to check if a string is purely numeric
func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// splitCamelCase splits camelCase or PascalCase into individual words
func splitCamelCase(s string) []string {
	var words []string
	var current strings.Builder

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) &&
			(unicode.IsLower(rune(s[i-1])) ||
				(i+1 < len(s) && unicode.IsLower(rune(s[i+1])))) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
		}
		current.WriteRune(r)
	}

	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}

// processIdentifier breaks down an identifier into its component terms
func (tc *TermCounter) processIdentifier(ident string) {
	// Split snake_case
	snakeParts := strings.Split(ident, "_")

	for _, part := range snakeParts {
		// Split camelCase
		camelParts := splitCamelCase(part)
		for _, word := range camelParts {
			tc.Add(word)
		}
	}
}

// visitor implements the ast.Visitor interface to traverse the AST
type visitor struct {
	counter *TermCounter
}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.Ident:
		// Process identifiers (variable names, function names, etc.)
		v.counter.processIdentifier(n.Name)
	case *ast.TypeSpec:
		// Process type names
		if n.Name != nil {
			v.counter.processIdentifier(n.Name.Name)
		}
	case *ast.FuncDecl:
		// Process function names
		if n.Name != nil {
			v.counter.processIdentifier(n.Name.Name)
		}
	}

	return v
}

// analyzeFile processes a single Go file
func analyzeFile(counter *TermCounter, filePath string) error {
	fset := token.NewFileSet()

	// Parse the Go source file
	file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", filePath, err)
	}

	// Walk the AST and collect terms
	ast.Walk(visitor{counter: counter}, file)
	return nil
}

// getTopTerms returns the most frequent terms
func getTopTerms(counter *TermCounter, limit int) []Term {
	// Convert map to slice for sorting
	terms := make([]Term, 0, len(counter.terms))
	for word, count := range counter.terms {
		terms = append(terms, Term{Word: word, Count: count})
	}

	// Sort by frequency (descending) and then alphabetically
	sort.Slice(terms, func(i, j int) bool {
		if terms[i].Count != terms[j].Count {
			return terms[i].Count > terms[j].Count
		}
		return terms[i].Word < terms[j].Word
	})

	// Return top N terms
	if limit > len(terms) {
		limit = len(terms)
	}
	return terms[:limit]
}
