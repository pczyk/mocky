package main

import "testing"

func TestMock(t *testing.T) {
	t.Run("latin, cyrillic and special characters", func(t *testing.T) {
		cases := map[string]string{
			"mock":                        "MoCk",
			"MOCK":                        "MoCk",
			"lorem ipsum dolor sit amet!": "LoReM iPsUm DoLoR sIt AmEt!",
			"хорошо":                      "ХоРоШо",
			"a!b@c#":                      "A!b@C#",
		}

		for input, want := range cases {
			got := mock(input)
			if got != want {
				t.Errorf("input %q, got %q want %q", input, got, want)
			}
		}
	})
}
