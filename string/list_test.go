// y string © 2019, Ramsey Dow. All rights reserved.
// SPDX-License-Identifier: BSD-2-Clause

package string

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
