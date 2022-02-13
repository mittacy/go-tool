//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64"
package sliceUtil

// DiffKeyType 求差集
// @param array
// @param arrays
// @return []KeyType 出现在第一个数组中但其他输入数组中没有的值，顺序为第一个数组的原顺序
func DiffKeyType(array []KeyType, arrays ...[]KeyType) []KeyType {
	m := ToMapKeyTypeStruct(arrays...)

	res := make([]KeyType, 0)
	for _, v := range array {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

// IntersectKeyType 求交集
// @param array
// @param arrays
// @return []KeyType 第一个数组中出现的且在其他每个输入数组中都出现的值，顺序为第一个数组的原顺序
func IntersectKeyType(array []KeyType, arrays ...[]KeyType) []KeyType {
	m := ToMapKeyTypeStruct(arrays...)

	res := make([]KeyType, 0, len(array))
	for _, v := range array {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return res
}

// TwoDiffKeyType 比较两个数组的差异元素
// @param a a数组
// @param b b数组
// @return inAnoB 在a不在b的元素
// @return inBnoA 在b不在a的元素
func TwoDiffKeyType(a []KeyType, b []KeyType) (inAnoB []KeyType, inBnoA []KeyType) {
	aMap := ToMapKeyTypeStruct(a)
	bMap := ToMapKeyTypeStruct(b)

	inBnoA = []KeyType{}
	for _, v := range b {
		if _, ok := aMap[v]; !ok {
			inBnoA = append(inBnoA, v)
		}
	}

	inAnoB = []KeyType{}
	for _, v := range a {
		if _, ok := bMap[v]; !ok {
			inAnoB = append(inAnoB, v)
		}
	}

	return inAnoB, inBnoA
}
