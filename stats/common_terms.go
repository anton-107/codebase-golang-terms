package main

import "strings"

// Add this after the Term and TermCounter types
var commonGoTerms = map[string]bool{
	// Language keywords
	"err":    true,
	"error":  true,
	"string": true,
	"int":    true,
	"bool":   true,
	"func":   true,
	"return": true,
	"if":     true,
	"else":   true,
	"for":    true,
	"range":  true,
	"break":  true,
	"nil":    true,

	// Common variable names
	"ctx":    true,
	"cmd":    true,
	"pkg":    true,
	"ptr":    true,
	"src":    true,
	"dst":    true,
	"buf":    true,
	"resp":   true,
	"req":    true,
	"val":    true,
	"vars":   true,
	"params": true,
	"args":   true,

	// Common package names
	"fmt":     true,
	"os":      true,
	"net":     true,
	"http":    true,
	"json":    true,
	"io":      true,
	"ioutil":  true,
	"sync":    true,
	"context": true,
	"bytes":   true,
	"strings": true,
	"assert":  true,
	"path":    true,
	"test":    true,
	"file":    true,

	// Common type-related terms
	"interface": true,
	"struct":    true,
	"type":      true,
	"map":       true,
	"slice":     true,
	"chan":      true,

	// Common method / variable names
	"init":   true,
	"new":    true,
	"make":   true,
	"len":    true,
	"cap":    true,
	"append": true,
	"close":  true,
	"delete": true,
	"name":   true,
	"config": true,
	"get":    true,
	"value":  true,
	"run":    true,
	"equal":  true,
	"dir":    true,

	// Common CLI terms
	"cobra":   true,
	"command": true,
}

// Add this helper function to check if a term is common
func IsCommonGoTerm(term string) bool {
	return commonGoTerms[strings.ToLower(term)]
}
