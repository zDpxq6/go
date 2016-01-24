package tempconv

import "testing"
//Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures
//in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
func TestCToF(t *testing.T) {
	const in, out = Celsius(100), Fahrenheit(212)
	if x := CToF(in); x != out {
		t.Errorf("CtoF(%v) = %v, want %v", in, x, out)
	}
}

func TestCToK(t *testing.T) {
	const in, out = Celsius(100), Kelvin(373.15)
	if x := CToK(in); x != out {
		t.Errorf("CtoK(%v) = %v, want %v", in, x, out)
	}
}

func TestFToC(t *testing.T) {
	const in, out = Fahrenheit(212), Celsius(100)
	if x := FToC(in); x != out {
		t.Errorf("CtoF(%v) = %v, want %v", in, x, out)
	}
}

func TestFToK(t *testing.T) {
	const in, out = Fahrenheit(212), Kelvin(373.15)
	if x := FToK(in); x != out {
		t.Errorf("FtoK(%v) = %v, want %v", in, x, out)
	}
}

func TestKToC(t *testing.T) {
	const in, out = Kelvin(373.15), Celsius(100)
	if x := KToC(in); x != out {
		t.Errorf("KtoC(%v) = %v, want %v", in, x, out)
	}
}

func TestKToF(t *testing.T) {
	const in, out = Celsius(100), Fahrenheit(212)
	if x := CToF(in); x != out {
		t.Errorf("CtoF(%v) = %v, want %v", in, x, out)
	}
}
