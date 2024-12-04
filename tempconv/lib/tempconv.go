package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

// CToF converte uma temperatura de Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // a conversão em si não altera a representação do valor, ele ainda é um float64, o que muda é a escala do valor (mudou pra Fahrenheit)
}

// FToC converte uma temperatura de Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) // a conversão em si não altera a representação do valor, ele ainda é um float64, o que muda é a escala do valor (mudou pra Fahrenheit)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
