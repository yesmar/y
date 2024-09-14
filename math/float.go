// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package math

import (
	"errors"
	"math"
	"math/big"
)

// BigFloatToBigInt converts big.Float values to big.Int values. This
// function is required because big.NewInt() simply casts to uint64, and
// is thus bounded by that type's maximum size, i.e. math.MaxUint64 or
// 18446744073709551615. To get around this, we use the Int() method on
// big.Float to directly store large values into a big.Int.
func BigFloatToBigInt(bf *big.Float) (*big.Int, error) {
	if bf == nil {
		return nil, errors.New("nil big.Float")
	}
	bi := new(big.Int)
	if bi == nil {
		return nil, errors.New("failed to allocate memory")
	}
	bf.Int(bi) // cf. https://stackoverflow.com/a/47546136
	return bi, nil
}

// Roundup returns the smallest number, specified to 1 decimal place,
// that is equal to or higher than its input. Roundup(4.02) returns
// 4.1, while Roundup(4.00) returns 4.0.
func Roundup(f float64) float64 {
	n := math.Round(f * 100000)
	if int64(n)%10000 == 0 {
		return n / 100000.0
	}
	return (math.Floor(n/10000) + 1) / 10.0
}
