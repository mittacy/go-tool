//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64 ValueType=string,rune,int8,int16,int,int32,int64,float32,float64"
package mapUtil

// GetKeysKeyTypeValueType 获取所有键
// @param m
// @return []KeyType
func GetKeysKeyTypeValueType(m map[KeyType]ValueType) []KeyType {
	keys := make([]KeyType, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// GetValuesKeyTypeValueType 获取所有值
// @param m
// @return []ValueType
func GetValuesKeyTypeValueType(m map[KeyType]ValueType) []ValueType {
	values := make([]ValueType, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// GetKVKeyTypeValueType 获取所有键和值
// @param m
// @return []ValueType
func GetKVKeyTypeValueType(m map[KeyType]ValueType) ([]KeyType, []ValueType) {
	keys := make([]KeyType, len(m))
	values := make([]ValueType, len(m))

	i := 0
	for k, v := range m {
		keys[i] = k
		values[i] = v
		i++
	}

	return keys, values
}
