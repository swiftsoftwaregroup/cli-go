package main

import "testing"

func TestGetGreeting(t *testing.T) {
	testCases := []struct {
		language string
		expected string
	}{
		{"en", "Hello"},
		{"es", "Hola"},
		{"bg", "Здравей"},
		{"de", "Hello"}, // Default case
	}

	for _, tc := range testCases {
		t.Run(tc.language, func(t *testing.T) {
			result := getGreeting(tc.language)
			if result != tc.expected {
				t.Errorf("getGreeting(%q) = %q; want %q", tc.language, result, tc.expected)
			}
		})
	}
}
