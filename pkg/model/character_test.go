package model

import "testing"

func TestGetCurrentHealth(t *testing.T) {
	cases := []struct {
		desc      string
		character Character
		expected  int
	}{
		{"no combat log", Character{Health: 10}, 10},
		{"with combat log", Character{Health: 10, HPLog: []int{2}}, 2},
	}
	for _, tc := range cases {
		actual := tc.character.GetCurrentHealth()
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, actual, tc.expected)
		}
	}
}

func TestHeal(t *testing.T) {
	cases := []struct {
		desc      string
		character Character
		heal      int
		expected  int
	}{
		{"no combat log", Character{Health: 10}, 2, 10},
		{"with combat log lower then max", Character{Health: 10, HPLog: []int{2}}, 2, 4},
		{"with combat log higher then max", Character{Health: 10, HPLog: []int{10}}, 2, 10},
	}
	for _, tc := range cases {
		actual := tc.character.Heal(tc.heal)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, actual, tc.expected)
		}
	}
}

func TestDamage(t *testing.T) {
	cases := []struct {
		desc      string
		character Character
		damage    int
		expected  int
	}{
		{"no combat log", Character{Health: 10, TemporaryHealth: 0}, 2, 8},
		{"no combat log below 0", Character{Health: 10, TemporaryHealth: 0}, 2, 8},
		{"with combat log", Character{Health: 10, TemporaryHealth: 0, HPLog: []int{2}}, 2, 0},
		{"with combat log below 0", Character{Health: 10, TemporaryHealth: 0, HPLog: []int{2}}, 4, 0},
		{"with combat log no thp", Character{Health: 10, TemporaryHealth: 2, HPLog: []int{}}, 2, 10},
		{"with combat log with thp", Character{Health: 10, TemporaryHealth: 2, HPLog: []int{2}}, 2, 2},
		{"with combat log with thp below 0", Character{Health: 10, TemporaryHealth: 2, HPLog: []int{2}}, 6, 0},
	}
	for _, tc := range cases {
		actual := tc.character.Damage(tc.damage)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, actual, tc.expected)
		}
	}
}

func TestTemporaryHeal(t *testing.T) {
	cases := []struct {
		desc      string
		character Character
		tHeal     int
		expected  int
	}{
		{"no tHP", Character{Health: 10, TemporaryHealth: 0}, 2, 2},
		{"with tHP", Character{Health: 10, TemporaryHealth: 5}, 2, 7},
	}
	for _, tc := range cases {
		actual := tc.character.TemporaryHeal(tc.tHeal)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d", tc.desc, actual, tc.expected)
		}
	}
}
