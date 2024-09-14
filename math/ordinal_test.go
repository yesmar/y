// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package math

import "testing"

func TestOrdinal(t *testing.T) {
	testpairs := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1st"},
		{input: 2, expected: "2nd"},
		{input: 3, expected: "3rd"},
		{input: 4, expected: "4th"},
		{input: 5, expected: "5th"},
		{input: 10, expected: "10th"},
		{input: 11, expected: "11th"},
		{input: 12, expected: "12th"},
		{input: 13, expected: "13th"},
		{input: 14, expected: "14th"},
		{input: 15, expected: "15th"},
		{input: 20, expected: "20th"},
		{input: 21, expected: "21st"},
		{input: 22, expected: "22nd"},
		{input: 23, expected: "23rd"},
		{input: 24, expected: "24th"},
		{input: 25, expected: "25th"},
		{input: 100, expected: "100th"},
		{input: 101, expected: "101st"},
		{input: 102, expected: "102nd"},
		{input: 103, expected: "103rd"},
		{input: 104, expected: "104th"},
		{input: 105, expected: "105th"},
	}

	for _, pair := range testpairs {
		actual := Ordinal(pair.input)
		if actual != pair.expected {
			t.Errorf("%v: expected %v, got %v", pair.input, pair.expected, actual)
		}
	}
}

func TestFromOrdinal(t *testing.T) {
	testpairs := []struct {
		input    string
		expected int
	}{
		{input: "1st", expected: 1},
		{input: "2nd", expected: 2},
		{input: "3rd", expected: 3},
		{input: "4th", expected: 4},
		{input: "5th", expected: 5},
		{input: "10th", expected: 10},
		{input: "11th", expected: 11},
		{input: "12th", expected: 12},
		{input: "13th", expected: 13},
		{input: "14th", expected: 14},
		{input: "15th", expected: 15},
		{input: "20th", expected: 20},
		{input: "21st", expected: 21},
		{input: "22nd", expected: 22},
		{input: "23rd", expected: 23},
		{input: "24th", expected: 24},
		{input: "25th", expected: 25},
		{input: "100th", expected: 100},
		{input: "101st", expected: 101},
		{input: "102nd", expected: 102},
		{input: "103rd", expected: 103},
		{input: "104th", expected: 104},
		{input: "105th", expected: 105},
	}

	for _, pair := range testpairs {
		actual, _ := FromOrdinal(pair.input)
		if actual != pair.expected {
			t.Errorf("%v: expected %v, got %v", pair.input, pair.expected, actual)
		}
	}
}
