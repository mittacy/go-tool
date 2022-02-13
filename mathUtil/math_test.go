package mathUtil

import "testing"

func TestMaxInt(t *testing.T) {
	cases := []struct {
		Name      string
		Param1    []int
		Expect    int
		ExpectVal int
	}{
		{
			Name:      "空",
			Param1:    []int{},
			Expect:    -1,
			ExpectVal: 0,
		},
		{
			Name:      "第一个",
			Param1:    []int{0, -2, -1, -2},
			Expect:    0,
			ExpectVal: 0,
		},
		{
			Name:      "最后一个",
			Param1:    []int{0, -2, -1, 2},
			Expect:    3,
			ExpectVal: 2,
		},
		{
			Name:      "其他",
			Param1:    []int{0, 2, -1, 2},
			Expect:    1,
			ExpectVal: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res1, res2 := MaxInt(c.Param1...)
			if res1 != c.Expect || res2 != c.ExpectVal {
				t.Fatalf("input: %v, expected %v, %v, but %v, %v got",
					c.Param1, c.Expect, c.ExpectVal, res1, res2)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	cases := []struct {
		Name      string
		Param1    []int
		Expect    int
		ExpectVal int
	}{
		{
			Name:      "空",
			Param1:    []int{},
			Expect:    -1,
			ExpectVal: 0,
		},
		{
			Name:      "第一个",
			Param1:    []int{-10, -2, -1, -2},
			Expect:    0,
			ExpectVal: -10,
		},
		{
			Name:      "最后一个",
			Param1:    []int{10, 23, 13, 5},
			Expect:    3,
			ExpectVal: 5,
		},
		{
			Name:      "其他",
			Param1:    []int{153, 2, 12, 2},
			Expect:    1,
			ExpectVal: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res1, res2 := MinInt(c.Param1...)
			if res1 != c.Expect || res2 != c.ExpectVal {
				t.Fatalf("input: %v, expected %v, %v, but %v, %v got",
					c.Param1, c.Expect, c.ExpectVal, res1, res2)
			}
		})
	}
}

func TestRoundFloat64(t *testing.T) {
	cases := []struct {
		Name   string
		Param1 float64
		Param2 int
		Expect float64
	}{
		{
			Name:   "零",
			Param1: 0,
			Param2: 2,
			Expect: 0,
		},
		{
			Name:   "0位四舍",
			Param1: 12.4,
			Param2: 0,
			Expect: 12,
		},
		{
			Name:   "1位四舍",
			Param1: 12.44123,
			Param2: 1,
			Expect: 12.4,
		},
		{
			Name:   "2位四舍",
			Param1: 12.444444,
			Param2: 2,
			Expect: 12.44,
		},
		{
			Name:   "0位五入",
			Param1: 12.500,
			Param2: 0,
			Expect: 13,
		},
		{
			Name:   "1位五入",
			Param1: 12.445,
			Param2: 1,
			Expect: 12.5,
		},
		{
			Name:   "2位五入",
			Param1: 12.44500,
			Param2: 2,
			Expect: 12.45,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := RoundFloat64Unsafe(c.Param1, c.Param2)
			if res != c.Expect {
				t.Fatalf("input: %v, %v, expected %v, but %v got",
					c.Param1, c.Param2, c.Expect, res)
			}
		})
	}
}
