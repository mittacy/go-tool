package sliceUtil

import "testing"

func TestToMapStringStruct(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Expect map[string]struct{}
	}{
		{"空", []string{}, map[string]struct{}{}},
		{"不重复", []string{"a"}, map[string]struct{}{"a": {}}},
		{"重复", []string{"a", "b", "b", "c", "c", "d", "d"}, map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := ToMapStringStruct(c.Param1)
			t.Logf("expect: %+v, get: %+v\n", c.Expect, res)
		})
	}
}

func TestDiffString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Param2 [][]string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: []string{},
			Param2: [][]string{},
			Expect: []string{},
		},
		{
			Name:   "单个无差异",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"1", "2"}},
			Expect: []string{},
		},
		{
			Name:   "多个无差异",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"1"}, {"2"}},
			Expect: []string{},
		},
		{
			Name:   "单个单差异",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"2", "3"}},
			Expect: []string{"1"},
		},
		{
			Name:   "单个多差异",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"3", "4"}},
			Expect: []string{"1", "2"},
		},
		{
			Name:   "多个单差异",
			Param1: []string{"1", "2", "3"},
			Param2: [][]string{{"1"}, {"2"}},
			Expect: []string{"3"},
		},
		{
			Name:   "多个多差异",
			Param1: []string{"1", "2", "3", "4", "5", "6"},
			Param2: [][]string{{"3"}, {"2"}},
			Expect: []string{"1", "4", "5", "6"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := DiffString(c.Param1, c.Param2...)
			if len(res) != len(c.Expect) {
				t.Fatalf("input: %v、%v, expected %v, but %v got",
					c.Param1, c.Param2, c.Expect, res)
			} else {
				for i := range res {
					if res[i] != c.Expect[i] {
						t.Fatalf("input: %v、%v, expected %v, but %v got",
							c.Param1, c.Param2, c.Expect, res)
					}
				}
			}
		})
	}
}

func TestIntersectString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Param2 [][]string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: []string{},
			Param2: [][]string{},
			Expect: []string{},
		},
		{
			Name:   "单个无相同",
			Param1: []string{"1"},
			Param2: [][]string{{"2"}},
			Expect: []string{},
		},
		{
			Name:   "多个无相同",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"3"}, {"4"}},
			Expect: []string{},
		},
		{
			Name:   "单个相同",
			Param1: []string{"1", "2"},
			Param2: [][]string{{"2", "3"}},
			Expect: []string{"2"},
		},
		{
			Name:   "多个单相同",
			Param1: []string{"1", "2", "3"},
			Param2: [][]string{{"1"}, {"5"}},
			Expect: []string{"1"},
		},
		{
			Name:   "多个多相同",
			Param1: []string{"1", "2", "3", "4", "5", "6"},
			Param2: [][]string{{"3"}, {"2", "6"}},
			Expect: []string{"2", "3", "6"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := IntersectString(c.Param1, c.Param2...)
			if len(res) != len(c.Expect) {
				t.Fatalf("input: %v、%v, expected %v, but %v got",
					c.Param1, c.Param2, c.Expect, res)
			} else {
				for i := range res {
					if res[i] != c.Expect[i] {
						t.Fatalf("input: %v、%v, expected %v, but %v got",
							c.Param1, c.Param2, c.Expect, res)
					}
				}
			}
		})
	}
}

func TestTwoDiffString(t *testing.T) {
	cases := []struct {
		Name    string
		Param1  []string
		Param2  []string
		Expect1 []string
		Expect2 []string
	}{
		{
			Name:    "空",
			Param1:  []string{},
			Param2:  []string{},
			Expect1: []string{},
			Expect2: []string{},
		},
		{
			Name:    "顺序相等",
			Param1:  []string{"1", "2"},
			Param2:  []string{"1", "2"},
			Expect1: []string{},
			Expect2: []string{},
		},
		{
			Name:    "乱序相等",
			Param1:  []string{"1", "2", "3", "4"},
			Param2:  []string{"2", "4", "1", "3"},
			Expect1: []string{},
			Expect2: []string{},
		},
		{
			Name:    "a比b多",
			Param1:  []string{"1", "2", "3", "4"},
			Param2:  []string{"1", "3"},
			Expect1: []string{"2", "4"},
			Expect2: []string{},
		},
		{
			Name:    "b比a多",
			Param1:  []string{"1", "3"},
			Param2:  []string{"1", "2", "3", "4"},
			Expect1: []string{},
			Expect2: []string{"2", "4"},
		},
		{
			Name:    "都有差异",
			Param1:  []string{"1", "3", "5", "7", "9"},
			Param2:  []string{"1", "2", "3", "4"},
			Expect1: []string{"5", "7", "9"},
			Expect2: []string{"2", "4"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res1, res2 := TwoDiffString(c.Param1, c.Param2)
			if len(res1) != len(c.Expect1) || len(res2) != len(c.Expect2) {
				t.Fatalf("input: %v、%v, expected %v, %v, but %v, %v got",
					c.Param1, c.Param2, c.Expect1, c.Expect2, res1, res2)
			}

			for i := range res1 {
				if res1[i] != c.Expect1[i] {
					t.Fatalf("input: %v、%v, expected %v, %v, but %v, %v got",
						c.Param1, c.Param2, c.Expect1, c.Expect2, res1, res2)
				}
			}

			for i := range res2 {
				if res2[i] != c.Expect2[i] {
					t.Fatalf("input: %v、%v, expected %v, %v, but %v, %v got",
						c.Param1, c.Param2, c.Expect1, c.Expect2, res1, res2)
				}
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: []string{},
			Expect: []string{},
		},
		{
			Name:   "单个",
			Param1: []string{"1", "1", "2"},
			Expect: []string{"1", "2"},
		},
		{
			Name:   "多个",
			Param1: []string{"1", "4", "2", "3", "2", "1"},
			Expect: []string{"1", "4", "2", "3"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := UniqueString(c.Param1)
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

func TestReverseString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Expect []string
	}{
		{
			Name:   "空",
			Param1: []string{},
			Expect: []string{},
		},
		{
			Name:   "单个",
			Param1: []string{"1"},
			Expect: []string{"1"},
		},
		{
			Name:   "多个",
			Param1: []string{"1", "4", "2", "3", "2", "1"},
			Expect: []string{"1", "2", "3", "2", "4", "1"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := ReverseString(c.Param1)
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

func TestIndexString(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 []string
		Param2 string
		Expect int
	}{
		{
			Name:   "不存在",
			Param1: []string{},
			Param2: "12",
			Expect: -1,
		},
		{
			Name:   "第一个",
			Param1: []string{"1", "2", "3"},
			Param2: "1",
			Expect: 0,
		},
		{
			Name:   "最后一个",
			Param1: []string{"1", "2", "3"},
			Param2: "3",
			Expect: 2,
		},
		{
			Name:   "其他",
			Param1: []string{"1", "2", "3"},
			Param2: "2",
			Expect: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if res := IndexString(c.Param1, c.Param2); res != c.Expect {
				t.Fatalf("input: %v、%v, expected %v, but %v got",
					c.Param1, c.Param2, c.Expect, res)
			}
		})
	}
}
