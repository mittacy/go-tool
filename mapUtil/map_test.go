package mapUtil

import "testing"

func TestGetKeysStringString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 map[string]string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: map[string]string{},
			Expect: []string{},
		},
		{
			Name:   "有键",
			Param1: map[string]string{"a": "b", "b": "a"},
			Expect: []string{"a", "b"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := GetKeysStringString(c.Param1)
			if len(res) != len(c.Expect) {
				t.Fatalf("input: %v, expected %v, but %v got",
					c.Param1, c.Expect, res)
			} else {
				for i := range res {
					if res[i] != c.Expect[i] {
						t.Fatalf("input: %v, expected %v, but %v got",
							c.Param1, c.Expect, res)
					}
				}
			}
		})
	}
}

func TestGetValuesStringString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 map[string]string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: map[string]string{},
			Expect: []string{},
		},
		{
			Name:   "值不重复",
			Param1: map[string]string{"a": "b", "b": "a"},
			Expect: []string{"b", "a"},
		},
		{
			Name:   "值重复",
			Param1: map[string]string{"a": "b", "b": "b", "c": ""},
			Expect: []string{"b", "b", ""},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := GetValuesStringString(c.Param1)
			t.Logf("input: %v, expected %v, %v got",
				c.Param1, c.Expect, res)
		})
	}
}
