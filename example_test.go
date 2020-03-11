package inf

import (
	"fmt"
	"log"
)

func ExampleDec_SetString() {
	d := new(Dec)
	d.SetString("012345.67890") // decimal; leading 0 ignored; trailing 0 kept
	fmt.Println(d)
	// Output: 12345.67890
}

func ExampleDec_Scan() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	d := new(Dec)
	_, err := fmt.Sscan("184467440.73709551617", d)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(d)
	}
	// Output: 184467440.73709551617
}

func ExampleDec_QuoRound_scale2RoundDown() {
	// 10 / 3 is an nite decimal; it has no exact Dec representation
	x, y := NewDec(10, 0), NewDec(3, 0)
	// use 2 digits beyond the decimal point, round towards 0
	z := new(Dec).QuoRound(x, y, 2, RoundDown)
	fmt.Println(z)
	// Output: 3.33
}

func ExampleDec_QuoRound_scale2RoundCeil() {
	// -42 / 400 is an finite decimal with 3 digits beyond the decimal point
	x, y := NewDec(-42, 0), NewDec(400, 0)
	// use 2 digits beyond decimal point, round towards positive nity
	z := new(Dec).QuoRound(x, y, 2, RoundCeil)
	fmt.Println(z)
	// Output: -0.10
}

func ExampleDec_QuoExact_ok() {
	// 1 / 25 is a finite decimal; it has exact Dec representation
	x, y := NewDec(1, 0), NewDec(25, 0)
	z := new(Dec).QuoExact(x, y)
	fmt.Println(z)
	// Output: 0.04
}

func ExampleDec_QuoExact_fail() {
	// 1 / 3 is an nite decimal; it has no exact Dec representation
	x, y := NewDec(1, 0), NewDec(3, 0)
	z := new(Dec).QuoExact(x, y)
	fmt.Println(z)
	// Output: <nil>
}
