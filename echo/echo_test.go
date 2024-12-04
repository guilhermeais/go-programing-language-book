package echo_test

import (
	"bytes"
	"echo"
	"testing"
)

func TestEcho1(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"No arguments", []string{}, ""},
		{"One argument", []string{"go run"}, ""},
		{"Multiple arguments", []string{"go run", "arg2", "arg3"}, "arg2 arg3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := echo.Echo1(tt.args)
			if got := result; got != tt.expected {
				t.Errorf("Echo1() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkEcho1(b *testing.B) {
	args := []string{"go", "run", "arg2", "arg3"}
	for i := 0; i < b.N; i++ {
		echo.Echo1(args)
	}
}

func TestEcho2(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"No arguments", []string{}, ""},
		{"One argument", []string{"go run"}, ""},
		{"Multiple arguments", []string{"go run", "arg2", "arg3"}, "arg2 arg3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := echo.Echo2(tt.args)
			if got := result; got != tt.expected {
				t.Errorf("Echo2() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkEcho2(b *testing.B) {
	args := []string{"go", "run", "arg2", "arg3"}
	for i := 0; i < b.N; i++ {
		echo.Echo1(args)
	}
}

func TestEcho3(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"No arguments", []string{}, ""},
		{"One argument", []string{"go run"}, ""},
		{"Multiple arguments", []string{"go run", "arg2", "arg3"}, "arg2 arg3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			echo.Echo3(tt.args, writer)
			got := writer.String()
			if got != tt.expected {
				t.Errorf("Echo3() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkEcho3(b *testing.B) {
	args := []string{"go", "run", "arg2", "arg3"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer := &bytes.Buffer{}
		echo.Echo3(args, writer)
	}
}
