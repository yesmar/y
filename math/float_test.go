// y math © 2019, Ramsey Dow. All rights reserved.

package math

import (
	"log"
	"math/big"
	"testing"
)

func TestBigFloatToBigInt(t *testing.T) {
	bf := big.NewFloat(18446744073709551616) // 1.844674407370955e+19
	expected := new(big.Int)
	bf.Int(expected) // Store 18446744073709551616 into expected without overflowing.

	actual, err := BigFloatToBigInt(bf)
	if err != nil {
		log.Fatal(err)
	}
	if actual.Cmp(expected) != 0 {
		t.Errorf("Expected %v, actual %v\n", expected, actual)
	}
}

func TestNilBigFloatToBigInt(t *testing.T) {
	actual, err := BigFloatToBigInt(nil)
	if err == nil {
		t.Error("Expected nil, got", actual)
	}
}
