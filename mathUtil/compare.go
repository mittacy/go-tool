//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=int8,int16,int,int32,int64,float32,float64"
package mathUtil

// MaxKeyType 求最大值
// @param nums
// @return int 最大值第一次出现的索引
// @return KeyType 最大值
func MaxKeyType(nums ...KeyType) (int, KeyType) {
	if len(nums) == 0 {
		return -1, 0
	}

	max := nums[0]
	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	return index, max
}

// MinKeyType 求最小值
// @param nums
// @return int 最小值第一次出现的索引
// @return KeyType 最小值
func MinKeyType(nums ...KeyType) (int, KeyType) {
	if len(nums) == 0 {
		return -1, 0
	}

	min := nums[0]
	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			index = i
		}
	}

	return index, min
}
