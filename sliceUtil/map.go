//go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64"
package sliceUtil

// ToMapKeyTypeStruct 函数功能说明
// this method will remove duplicate elements
// @param values
// @return map[KeyType]struct{}
func ToMapKeyTypeStruct(values ...[]KeyType) map[KeyType]struct{} {
	initLength := 0
	if len(values) == 1 {
		initLength = len(values[0])
	} else if len(values) > 0 {
		initLength = len(values[0]) * len(values)
	}

	m := make(map[KeyType]struct{}, initLength)
	for _, arr := range values {
		for _, v := range arr {
			m[v] = struct{}{}
		}
	}
	return m
}
