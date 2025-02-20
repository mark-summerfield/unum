// Copyright © 2024-25 Mark Summerfield. All rights reserved.

// ([TOC]) This package provides number-related functions and some generic
// numeric types.
//
// [TOC]: file:///home/mark/app/golib/doc/index.html
package unum

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed Version.dat
var Version string

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Number interface {
	~float64 | ~float32 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type SignedNumber interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

const decimalFactor = 1e-9 // use: 1e-1 (1 sf) to 1e-14 (14 sf)

// IsClose returns true if a and b are very close to each other.
// Should be adequate for test comparisons.
// See also [IsZero].
// tag::IsClose[]
func IsClose(a, b float64) bool {
	return math.Abs(a-b) <= max(decimalFactor*max(math.Abs(a),
		math.Abs(b)), math.SmallestNonzeroFloat64)
}

// end::IsClose[]

// IsZero returns true if x is close to 0.
// Should be adequate for test comparisons.
// See also [IsClose].
// tag::IsZero[]
func IsZero(x float64) bool {
	return math.Abs(x) <= max(decimalFactor*max(math.Abs(x), 0),
		math.SmallestNonzeroFloat64)
}

func RoundToNearest(f float64, nearestTo int) float64 {
	if nearestTo <= 1 {
		return f
	}
	ntf := float64(nearestTo)
	return math.Round(f/ntf) * ntf
}

func MustStrToInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func StrToInt(s string, default_ int) int {
	if i, err := strconv.Atoi(s); err != nil {
		return default_
	} else {
		return i
	}
}

func Commas[I Integer](i I) string {
	sign := ""
	value := fmt.Sprint(i) // Can't use Itoa() with Integer
	if value[0] == '-' {
		sign = "-"
		value = value[1:]
	}
	for i := len(value) - 3; i >= 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return sign + strings.TrimPrefix(value, ",")
}
