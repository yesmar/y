// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package strings

import "strings"

// DelimitedList returns a list of the elements in xs
// as a string, delimited by delim.
func DelimitedList(xs []string, delim rune, conjunction bool) string {
	var sb strings.Builder
	for i, s := range xs {
		sb.WriteString(s)
		if i < len(xs)-1 {
			sb.WriteRune(delim)
			sb.WriteRune(' ')
		}
		if conjunction && i == len(xs)-2 {
			sb.WriteString("and ")
		}
	}
	return sb.String()
}

// List returns a comma-delimited list of the elements
// in xs as a string.
func List(xs []string, conjunction bool) string {
	return DelimitedList(xs, ',', conjunction)
}
