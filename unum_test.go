package unum

import (
	"fmt"
	"testing"
)

func Test_IsReal(t *testing.T) {
	a := 0.0
	for range 1000 {
		a += 0.01
	}
	// a: 9.999999999999831
	b := 0.0
	for range 100 {
		b += 0.1
	}
	// b: 9.99999999999998
	c := 0.0
	for range 10 {
		c += 1.0
	}
	// c: 10
	d := 10.0
	// d: 10
	if a == b || a == c || b == c {
		t.Error("expected unequal, got equal")
	}
	if c != d {
		t.Error("expected equal, got unequal")
	}
	if !IsClose(a, b) {
		t.Error("expected real close, got not real close")
	}
	if !IsClose(a, c) {
		t.Error("expected real close, got not real close")
	}
	if !IsClose(a, d) {
		t.Error("expected real close, got not real close")
	}
	if !IsClose(b, c) {
		t.Error("expected real close, got not real close")
	}
	if !IsClose(b, d) {
		t.Error("expected real close, got not real close")
	}
	if !IsClose(c, d) {
		t.Error("expected real close, got not real close")
	}
	//if !IsZero(a - 10) {
	//	t.Errorf("expected real zero, got not real zero %f %f", a, a-10)
	//}
	//if !IsZero(b - 10) {
	//	t.Error("expected real zero, got not real zero")
	//}
	if !IsZero(c - 10) {
		t.Error("expected real zero, got not real zero")
	}
	if !IsZero(d - 10) {
		t.Error("expected real zero, got not real zero")
	}
	//if !IsZero(10 - a) {
	//	t.Error("expected real zero, got not real zero")
	//}
	//if !IsZero(10 - b) {
	//	t.Error("expected real zero, got not real zero")
	//}
	if !IsZero(10 - c) {
		t.Error("expected real zero, got not real zero")
	}
	if !IsZero(10 - d) {
		t.Error("expected real close, got not real close")
	}
}

func Test_StrToInt(t *testing.T) {
	data := []struct {
		S string
		I int
	}{
		{"159", 159},
		{"2196", 2196},
		{"0", 0},
		{"-14", -14},
		{"F4", 999},
	}
	for n, datum := range data {
		i := StrToInt(datum.S, 999)
		if i != datum.I {
			t.Errorf("#%d expected %q,\ngot %d", n, datum.I, i)
		}
	}
}

func Test_RoundToNearest(t *testing.T) {
	data := []struct {
		F float64
		E string
		N int
	}{
		{0, "0", 0},                // 0
		{11, "11", 0},              // 1
		{123, "123", 0},            // 2
		{1234, "1234", 0},          // 3
		{12345, "12345", 0},        // 4
		{123456, "123456", 0},      // 5
		{1234567, "1234567", 0},    // 6
		{0, "0", 10},               // 7
		{11, "10", 10},             // 8
		{123, "120", 10},           // 9
		{1234, "1230", 10},         // 10
		{12345, "12350", 10},       // 11
		{123456, "123460", 10},     // 12
		{1234567, "1234570", 10},   // 13
		{0, "0", 100},              // 14
		{11, "0", 100},             // 15
		{123, "100", 100},          // 16
		{1234, "1200", 100},        // 17
		{12345, "12300", 100},      // 18
		{123456, "123500", 100},    // 19
		{1234567, "1234600", 100},  // 20
		{0, "0", 1000},             // 21
		{11, "0", 1000},            // 22
		{123, "0", 1000},           // 23
		{1234, "1000", 1000},       // 24
		{12345, "12000", 1000},     // 25
		{123456, "123000", 1000},   // 26
		{1234567, "1235000", 1000}, // 27
		{0, "0", 5},
		{5, "5", 5},
		{10, "10", 5},
		{12, "10", 5},
		{16, "15", 5},
		{19, "20", 5},
		{20, "20", 5},
		{22, "20", 5},
		{26, "25", 5},
		{93, "95", 5},
		{-0, "0", 5},
		{-5, "-5", 5},
		{-10, "-10", 5},
		{-12, "-10", 5},
		{-16, "-15", 5},
		{-19, "-20", 5},
		{-20, "-20", 5},
		{-22, "-20", 5},
		{-26, "-25", 5},
	}
	for i, datum := range data {
		a := RoundToNearest(datum.F, datum.N)
		s := fmt.Sprintf("%.0f", a)
		if s != datum.E {
			t.Errorf("expected %q got %q @ %d", datum.E, s, i)
		}
	}
}

func TestCommas(t *testing.T) {
	ints := []int{
		-1, -200, -3450, -17392, 0, 1, 20, 344, 4834, 58302,
		2934849, 9879132421,
	}
	expected := []string{
		"-1", "-200", "-3,450", "-17,392", "0", "1", "20",
		"344", "4,834", "58,302", "2,934,849", "9,879,132,421",
	}
	for i := 0; i < len(ints); i++ {
		actual := Commas(ints[i])
		if actual != expected[i] {
			t.Errorf("expected %s got %s", expected[i], actual)
		}
	}
	uints := []uint{
		1, 200, 3450, 17392, 0, 1, 20, 344, 4834, 58302,
		2934849, 9879132421,
	}
	expected = []string{
		"1", "200", "3,450", "17,392", "0", "1", "20",
		"344", "4,834", "58,302", "2,934,849", "9,879,132,421",
	}
	for i := 0; i < len(uints); i++ {
		actual := Commas(uints[i])
		if actual != expected[i] {
			t.Errorf("expected %s got %s", expected[i], actual)
		}
	}
	runes := []rune{
		1, 200, 3450, 17392, 0, 1, 20, 344, 4834, 58302,
		2934849,
	}
	expected = []string{
		"1", "200", "3,450", "17,392", "0", "1", "20",
		"344", "4,834", "58,302", "2,934,849",
	}
	for i := 0; i < len(runes); i++ {
		actual := Commas(runes[i])
		if actual != expected[i] {
			t.Errorf("expected %s got %s", expected[i], actual)
		}
	}
}

func Test_Clamp(t *testing.T) {
	if v := Clamp(-5, 0, 11); v != 0 {
		t.Errorf("expected 0 got %d", v)
	}
	if v := Clamp(-11.5, -15, 0); !IsClose(v, -11.5) {
		t.Errorf("expected 0 got %f", v)
	}
	if v := Clamp(0, 101, 100); v != 100 {
		t.Errorf("expected 0 got %d", v)
	}
}

func Benchmark_commas(b *testing.B) {
	for i := range b.N {
		_ = Commas(i)
	}
}
