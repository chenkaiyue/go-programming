package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
const(
	AbsoluteZeroC Celsius =-273.15
	FreezingC Celsius = 0.00
	BoilingC Celsius = 100.00
)

func (c Celsius)String() string{
	return fmt.Sprintf("%g",c)
}

func (f Fahrenheit)String() string{
	return fmt.Sprintf("%g",f)
}
