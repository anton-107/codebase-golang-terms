package main

import "strings"

// AnalyzerConfig holds configuration for the term analyzer
type AnalyzerConfig struct {
	MinTermLength int
	ExcludeTerms  map[string]bool
	MaxTerms      int
}

// NewDefaultConfig creates a default configuration
func NewDefaultConfig() *AnalyzerConfig {
	return &AnalyzerConfig{
		MinTermLength: 3,
		ExcludeTerms:  commonGoTerms,
		MaxTerms:      50,
	}
}

// Allow adding custom terms to exclude
func (c *AnalyzerConfig) AddExcludeTerms(terms ...string) {
	for _, term := range terms {
		c.ExcludeTerms[strings.ToLower(term)] = true
	}
}
