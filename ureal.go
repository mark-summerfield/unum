// Copyright © 2024 Mark Summerfield. All rights reserved.

package ureal

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
