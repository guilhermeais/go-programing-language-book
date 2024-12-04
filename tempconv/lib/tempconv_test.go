package tempconv_test

import (
	tempconv "tempconv/lib"
	"testing"
)

func TestCToF(t *testing.T) {
	var c tempconv.Celsius = 10
	var want tempconv.Fahrenheit = 50

	got := tempconv.CToF(c)

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestFToC(t *testing.T) {
	var f tempconv.Fahrenheit = 50
	var want tempconv.Celsius = 10

	got := tempconv.FToC(f)

	if got != want {
		t.Fatalf("got %v want %v", got, f)
	}
}

func TestCelsiusString(t *testing.T) {
	var c tempconv.Celsius = 10
	got := c.String()
	want := "10°C"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestFahrenheitString(t *testing.T) {
	var c tempconv.Fahrenheit = 10
	got := c.String()
	want := "10°F"
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
