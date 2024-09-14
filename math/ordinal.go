// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package math

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// Ordinal converts the generic integer type to a string ordinal,
// e.g., Ordinal(3) returns "3rd". Although any integeral type is
// accepted, internally it is treated as an int because the mod
// operator cannot be used with generically constrained types.
// Thus, whatever is passed in is converted to an int. Moreover,
// the absolute valus is taken, so Ordinal(-1) will return "1st".
func Ordinal[T constraints.Integer](n T) string {
	var x int
	switch {
	case n < 0:
		x = int(math.Abs(float64(n)))
	default:
		x = int(n)
	}

	if x >= 11 && x <= 13 {
		return fmt.Sprintf("%dth", x)
	}

	switch x % 10 {
	case 1:
		return fmt.Sprintf("%dst", x)
	case 2:
		return fmt.Sprintf("%dnd", x)
	case 3:
		return fmt.Sprintf("%drd", x)
	default:
		return fmt.Sprintf("%dth", x)
	}
}

// FromOrdinal will return the integral form of the caller-
// supplied ordinal string, e.g., FromOrdinal("11th") will
// return 11, with err set in case a parse error occurs.
func FromOrdinal(s string) (int, error) {
	var n int
	var junk string
	_, err := fmt.Sscanf(s, "%d%s", &n, &junk)
	if err != nil {
		return 0, err
	}
	return n, nil
}
