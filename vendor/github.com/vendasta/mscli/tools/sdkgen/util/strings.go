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
		if runes[i] != '_' && runes[i] != '-' {
			if i > 0 && (runes[i-1] == '_' || runes[i-1] == '-') {
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

	return strings.Replace(string(out), "-", "_", -1)
}

// ToKebabCase converts a string to kebab case (kebab-case)
func ToKebabCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		thisIsUpper := unicode.IsUpper(runes[i])
		nextIsLower := (i+1 < length && unicode.IsLower(runes[i+1]))
		prevIsLower := (i > 0 && unicode.IsLower(runes[i-1]))
		besideUnderscore := (i+1 < length && (runes[i+1] == '_' || runes[i+1] == '-')) || (i > 0 && (runes[i-1] == '_' || runes[i-1] == '-'))
		if i > 0 && thisIsUpper && (nextIsLower || prevIsLower) && !besideUnderscore {
			out = append(out, '-')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return strings.Replace(string(out), "_", "-", -1)
}
