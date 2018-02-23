package util

import (
	"testing"
)

type TestCase struct {
	input  string
	output string
}

func Test_ToSnakeCase(t *testing.T) {
	var tests = []TestCase{
		{"", ""},
		{"a", "a"},
		{"snake", "snake"},
		{"A", "a"},
		{"ID", "id"},
		{"MOTD", "motd"},
		{"Snake", "snake"},
		{"SnakeTest", "snake_test"},
		{"SnakeID", "snake_id"},
		{"SnakeIDGoogle", "snake_id_google"},
		{"LinuxMOTD", "linux_motd"},
		{"OMGWTFBBQ", "omgwtfbbq"},
		{"omg_wtf_bbq", "omg_wtf_bbq"},
		{"APIResponse", "api_response"},
	}

	for _, test := range tests {
		if ToSnakeCase(test.input) != test.output {
			t.Errorf(`ToSnake("%s"), wanted "%s", got "%s"`, test.input, test.output, ToSnakeCase(test.input))
		}
	}
}

func Test_ToCamelCase(t *testing.T) {
	var tests = []TestCase{
		{"", ""},
		{"a", "a"},
		{"snake", "snake"},
		{"A", "a"},
		{"ID", "id"},
		{"MOTD", "motd"},
		{"Snake", "snake"},
		{"snake_test", "snakeTest"},
		{"snake_id", "snakeId"},
		{"snake_id_google", "snakeIdGoogle"},
		{"linux_motd", "linuxMotd"},
		{"OMGWTFBBQ", "omgwtfbbq"},
		{"omg_wtf_bbq", "omgWtfBbq"},
		{"api_response", "apiResponse"},
	}

	for _, test := range tests {
		if ToCamelCase(test.input) != test.output {
			t.Errorf(`ToCamelCase("%s"), wanted "%s", got "%s"`, test.input, test.output, ToCamelCase(test.input))
		}
	}
}

func Test_ToPascalCase(t *testing.T) {
	var tests = []TestCase{
		{"", ""},
		{"a", "A"},
		{"snake", "Snake"},
		{"A", "A"},
		{"ID", "Id"},
		{"MOTD", "Motd"},
		{"Snake", "Snake"},
		{"snake_test", "SnakeTest"},
		{"snake_id", "SnakeId"},
		{"snake_id_google", "SnakeIdGoogle"},
		{"linux_motd", "LinuxMotd"},
		{"OMGWTFBBQ", "Omgwtfbbq"},
		{"omg_wtf_bbq", "OmgWtfBbq"},
		{"api_response", "ApiResponse"},
	}

	for _, test := range tests {
		if ToPascalCase(test.input) != test.output {
			t.Errorf(`ToPascalCase("%s"), wanted "%s", got "%s"`, test.input, test.output, ToPascalCase(test.input))
		}
	}
}

func Test_ReplaceAll(t *testing.T) {
	type ReplaceAllTestCase struct {
		input                  string
		charactersToReplace    []string
		characterToReplaceWith string
		output                 string
	}
	var tests = []ReplaceAllTestCase{
		{"", nil, "", ""},
		{"this is the string", nil, "a", "this is the string"},
		{"this is the string", []string{""}, "a", "atahaiasa aiasa atahaea asataraianaga"},
		{"this is the string", []string{"i"}, "", "ths s the strng"},
		{"this is the string", []string{"i", "s"}, "", "th  the trng"},
		{"this is the string", []string{"i", "s", "x"}, "", "th  the trng"},
		{"this is the string", []string{"i", "s", "t"}, "", "h  he rng"},
		{"this is the string", []string{"i", "s", "t"}, "y", "yhyy yy yhe yyryng"},
		{"this is the string", []string{"i", "i", "i"}, "y", "thys ys the stryng"},
		{"hi", []string{"h", "i"}, "ii", "iiiiii"},
		{"hi", []string{"i", "h"}, "ii", "iiii"},
	}

	for _, test := range tests {
		if ReplaceAll(test.input, test.charactersToReplace, test.characterToReplaceWith) != test.output {
			t.Errorf(`ReplaceAll("%s"), wanted "%s", got "%s"`, test.input, test.output, ReplaceAll(test.input, test.charactersToReplace, test.characterToReplaceWith))
		}
	}
}
