package codewars_test

import (
	"codewars/codewars"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	// check testify
)

func TestRoman(t *testing.T) {
	type check func(t *testing.T, act int) error

	hasExpResult := func(exp int) check {
		return func(t *testing.T, act int) error {
			t.Helper()

			if diff := cmp.Diff(exp, act); diff != "" {
				return fmt.Errorf("unexpected error: wanted %d, got %d", exp, act)
			}
			return nil
		}
	}

	tc := []struct {
		name   string
		in     string
		checks []check
	}{
		{
			name:   "consecutive smaller roman digit values",
			in:     "MCXX",
			checks: []check{hasExpResult(1120)},
		}, {
			name:   "alternating smaller and larget roman digit values",
			in:     "MICIX",
			checks: []check{hasExpResult(1099)},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			res := codewars.FromRoman(c.in)
			for _, check := range c.checks {
				if err := check(t, res); err != nil {
					t.Fatalf("aaa!: %v", err)
				}
			}
		})
	}

}
