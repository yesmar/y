// y math © 2019, Ramsey Dow. All rights reserved.

package math

import (
	"errors"
	"math/big"
)

// BigFloatToBigInt converts big.Float values to big.Int values. This function is required because big.NewInt() simply casts to uint64, and is thus bounded by that type's maximum size, i.e. math.MaxUint64 or 18446744073709551615. To get around this, we use the Int() method on big.Float to directly store large values into a big.Int.
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
