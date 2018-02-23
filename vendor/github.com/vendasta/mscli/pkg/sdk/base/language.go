package base

import (
	"fmt"
)

// Language programming language enum
type Language string

// Programming language enums
const (
	LanguagePython     Language = "python"
	LanguageTypescript Language = "typescript"
	LanguageJava       Language = "java"
)

func LanguageFromString(language string) (Language, error) {
	switch language {
	case "python":
		return LanguagePython, nil
	case "typescript":
		return LanguageTypescript, nil
	case "java":
		return LanguageJava, nil
	default:
		return LanguagePython, fmt.Errorf("could not detect language, must be one of ['python', 'typescript', 'java']")
	}
}
