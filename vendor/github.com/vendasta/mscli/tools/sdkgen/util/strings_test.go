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

func Test_ToKebabCase(t *testing.T) {
	var tests = []TestCase{
		{"", ""},
		{"a", "a"},
		{"kebab", "kebab"},
		{"A", "a"},
		{"ID", "id"},
		{"MOTD", "motd"},
		{"kebab", "kebab"},
		{"kebabTest", "kebab-test"},
		{"kebabID", "kebab-id"},
		{"kebabIDGoogle", "kebab-id-google"},
		{"LinuxMOTD", "linux-motd"},
		{"OMGWTFBBQ", "omgwtfbbq"},
		{"omg-wtf-bbq", "omg-wtf-bbq"},
		{"APIResponse", "api-response"},
		{"API_Response", "api-response"},
	}

	for _, test := range tests {
		if ToKebabCase(test.input) != test.output {
			t.Errorf(`ToKebabCase("%s"), wanted "%s", got "%s"`, test.input, test.output, ToKebabCase(test.input))
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
		{"api-response", "apiResponse"},
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
