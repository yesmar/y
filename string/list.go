// y string © 2019, Ramsey Dow. All rights reserved.
// SPDX-License-Identifier: BSD-2-Clause

package string

import "strings"

// DelimitedList returns a list of the elements in xs
// as a string, delimited by delim.
func DelimitedList(delim rune, xs []string) string {
	var sb strings.Builder
	for i, s := range xs {
		sb.WriteString(s)
		if i < len(xs)-1 {
			sb.WriteRune(delim)
			sb.WriteRune(' ')
		}
		if i == len(xs)-2 {
			sb.WriteString("and ")
		}
	}
	return sb.String()
}

// List returns a comma-delimited list of the elements
// in xs as a string.
func List(xs []string) string {
	return DelimitedList(',', xs)
}
