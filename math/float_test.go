// y math © 2019, Ramsey Dow. All rights reserved.
// SPDX-License-Identifier: BSD-2-Clause

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

func TestRoundup(t *testing.T) {
	testpairs := []struct {
		input    float64
		expected float64
	}{
		{input: 4.02, expected: 4.1},
		{input: 4.00, expected: 4.0},
		{input: 6.4467, expected: 6.5},
		{input: 8.265, expected: 8.3},
		{input: 4.8546, expected: 4.9},
		{input: 1.200003, expected: 1.2},
	}

	for _, pair := range testpairs {
		actual := Roundup(pair.input)
		if actual != pair.expected {
			t.Errorf("%v: expected %v, got %v", pair.input, pair.expected, actual)
		}
	}
}
