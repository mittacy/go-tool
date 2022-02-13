//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64"
package sliceUtil

// UniqueKeyType 切片去重
// @param values
// @return []KeyType
func UniqueKeyType(values []KeyType) []KeyType {
	m := make(map[KeyType]struct{}, len(values))
	res := make([]KeyType, 0, len(values))
	for _, v := range values {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = struct{}{}
		}
	}

	return res
}

// ReverseKeyType 切片反转
// @param values
// @return []KeyType
func ReverseKeyType(values []KeyType) []KeyType {
	i, j := 0, len(values)-1

	for i < j {
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}

	return values
}

// IndexKeyType 查询target在values中第一次出现的索引位置，不在则返回-1
// @param values
// @param target
// @return int
func IndexKeyType(values []KeyType, target KeyType) int {
	for i, v := range values {
		if target == v {
			return i
		}
	}

	return -1
}

// IsInArrKeyType 判断target是否在values中
// @param values
// @param target
// @return bool
func IsInArrKeyType(values []KeyType, target KeyType) bool {
	if IndexKeyType(values, target) < 0 {
		return false
	}
	return false
}
