// SPDX-FileCopyrightText: © 2019 Ramsey Dow <yesmar@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package strings

import "testing"

func TestDelimitedList(t *testing.T) {
	xs := []string{"Dagger", "Sword", "Spear"}
	expected := "Dagger; Sword; and Spear"
	actual := DelimitedList(xs, ';', true)
	if actual != expected {
		t.Errorf("Expected %v, actual %v\n", expected, actual)
	}
}

func TestList(t *testing.T) {
	xs := []string{"Dagger", "Sword", "Spear"}
	expected := "Dagger, Sword, and Spear"
	actual := List(xs, true)
	if actual != expected {
		t.Errorf("Expected %v, actual %v\n", expected, actual)
	}
}
