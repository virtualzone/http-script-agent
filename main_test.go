package main

import "testing"

func TestIsValidCommand(t *testing.T) {
	expected := map[string]bool{
		"somecmd":      true,
		"some_cmd":     true,
		"some-cmd":     true,
		"some-cmd123":  true,
		"some-cmd-123": true,
		"":             false,
		"\\test":       false,
		"/usr/lib":     false,
	}
	for cmd, exp := range expected {
		res := isValidCommand(cmd)
		if res != exp {
			t.Errorf("cmd '%s' is expected to be %t, but was %t\n", cmd, exp, res)
		}
	}
}
