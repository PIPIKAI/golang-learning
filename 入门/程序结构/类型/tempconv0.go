package main

import "fmt"

type Celsius float32    // 摄氏温度
type Fahrenheit float32 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func main() {
	fmt.Printf("%g\n", BoilingC)
	BoilingF := CToF(BoilingC)
	fmt.Printf("%g\n", BoilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", BoilingF-FreezingC)  编译错误 (mismatched types Fahrenheit and Celsius)
}
