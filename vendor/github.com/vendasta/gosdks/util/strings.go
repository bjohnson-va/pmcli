package util

import (
	"strings"
	"unicode"
)

// StringInSlice returns true if target is in the provided list.
func StringInSlice(target string, list []string) bool {
	for _, candidate := range list {
		if candidate == target {
			return true
		}
	}
	return false
}

// ToCamelCase converts a string to camel case (camelCase)
func ToCamelCase(in string) string {
	return toCamelCase(in, false)
}

// ToPascalCase converts a string to pascal case (PascalCase)
func ToPascalCase(in string) string {
	return toCamelCase(in, true)
}

func toCamelCase(in string, upper bool) string {
	runes := []rune(in)
	length := len(runes)

	if len(in) == 0 {
		return ""
	}

	var out []rune
	for i := 0; i < length; i++ {
		if runes[i] != '_' {
			if i > 0 && runes[i-1] == '_' {
				out = append(out, unicode.ToUpper(runes[i]))
			} else {
				out = append(out, unicode.ToLower(runes[i]))
			}
		}
	}

	if upper {
		out[0] = unicode.ToUpper(out[0])
	} else {
		out[0] = unicode.ToLower(out[0])
	}

	return string(out)
}

// ToSnakeCase converts a string to snake case (snake_case)
func ToSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// ReplaceAll returns a copy of the string s with all instances of all the old characters replaced by new.
// If olds contains the empty string, it matches at the beginning of the string and after each UTF-8 sequence,
// yielding up to k+1 replacements for a k-rune string. The characters are replace in the same order as 'olds'
// which means this is not transitive
func ReplaceAll(s string, olds []string, new string) string {
	r := string(s)
	for _, old := range olds {
		r = strings.Replace(r, old, new, -1)
	}
	return r
}
