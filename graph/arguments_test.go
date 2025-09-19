package graph

import (
	"reflect"
	"testing"
)

// Helper function to create a pointer to an int
func ptrToInt(val int) *int {
	return &val
}

func TestGetIntOrDefault(t *testing.T) {
	tests := []struct {
		raw         any
		defaultVal  int
		expected    int
		description string
	}{
		{
			raw:         42,
			defaultVal:  0,
			expected:    42,
			description: "Valid int, should return the value",
		},
		{
			raw:         "string",
			defaultVal:  10,
			expected:    10,
			description: "Invalid type (string), should return default",
		},
		{
			raw:         nil,
			defaultVal:  5,
			expected:    5,
			description: "Nil value, should return default",
		},
		{
			raw:         0,
			defaultVal:  100,
			expected:    0,
			description: "Zero value, should return zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := getIntOrDefault(tt.raw, tt.defaultVal)
			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}

func TestGetIntPointerOrDefault(t *testing.T) {
	defaultVal := 10
	tests := []struct {
		raw         any
		defaultVal  *int
		expected    *int
		description string
	}{
		{
			raw:         42,
			defaultVal:  &defaultVal,
			expected:    ptrToInt(42),
			description: "Valid int, should return pointer to the value",
		},
		{
			raw:         ptrToInt(100),
			defaultVal:  &defaultVal,
			expected:    ptrToInt(100),
			description: "Pointer to int, should return the pointer",
		},
		{
			raw:         "string",
			defaultVal:  &defaultVal,
			expected:    &defaultVal,
			description: "Invalid type (string), should return default pointer",
		},
		{
			raw:         nil,
			defaultVal:  &defaultVal,
			expected:    &defaultVal,
			description: "Nil value, should return default pointer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := getIntPointerOrDefault(tt.raw, tt.defaultVal)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", *tt.expected, *result)
			}
		})
	}
}

func TestGetStringOrDefault(t *testing.T) {
	defaultVal := "default"
	tests := []struct {
		raw         any
		defaultVal  string
		expected    string
		description string
	}{
		{
			raw:         "Hello, World!",
			defaultVal:  defaultVal,
			expected:    "Hello, World!",
			description: "Valid string, should return the string",
		},
		{
			raw:         42, // Not a string
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Non-string type (int), should return default value",
		},
		{
			raw:         nil,
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Nil value, should return default value",
		},
		{
			raw:         3.14, // Not a string
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Non-string type (float), should return default value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := getStringOrDefault(tt.raw, tt.defaultVal)
			if result != tt.expected {
				t.Errorf("Expected '%s' but got '%s'", tt.expected, result)
			}
		})
	}
}

func TestGetBoolOrDefault(t *testing.T) {
	defaultVal := true
	tests := []struct {
		raw         any
		defaultVal  bool
		expected    bool
		description string
	}{
		{
			raw:         true,
			defaultVal:  defaultVal,
			expected:    true,
			description: "Valid bool (true), should return the boolean value",
		},
		{
			raw:         false,
			defaultVal:  defaultVal,
			expected:    false,
			description: "Valid bool (false), should return the boolean value",
		},
		{
			raw:         42, // Not a bool
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Non-bool type (int), should return default value",
		},
		{
			raw:         "string", // Not a bool
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Non-bool type (string), should return default value",
		},
		{
			raw:         nil, // Not a bool
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Nil value, should return default value",
		},
		{
			raw:         3.14, // Not a bool
			defaultVal:  defaultVal,
			expected:    defaultVal,
			description: "Non-bool type (float), should return default value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := getBoolOrDefault(tt.raw, tt.defaultVal)
			if result != tt.expected {
				t.Errorf("Expected '%v' but got '%v'", tt.expected, result)
			}
		})
	}
}
