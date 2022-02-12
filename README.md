# go-tool

go工具函数库

## 目录结构与规范

+ 包：`业务名Util`
+ 文件名：
    + `功能.go` 函数模板文件，需要生成不同类型参数的函数，可利用genny生成，生成后文件名为：`gen_功能.go`
    + `custom_功能.go`  不需要生成不同类型参数的函数，可自行编写
+ 函数名：`DiffKeyTypeValueType()` 
    + KeyType：入参1类型
    + ValueType：入参2类型
    + ResType：返回值类型
    + ……

### 包名文件名举例

```go
├── sliceUtil			// 业务包名 + Util
│   ├── map.go			// 模板函数文件
│   ├── gen_map.go		// 模板函数生成的不同参数文件
│   ├── custom_map.go	// 自定义函数文件
│   └── map_test.go		// 测试文件
└── timeUtil
	├── time.go
	├── gen_time.go
	├── custom_time.go
    └── time_test.go
```

### 函数命名举例

一般我们调用方式时，会先输入业务，然后输入功能，最后回想入参是什么类型

例子：先要把一个字符串切片转为map[string]struct{}，然后用来循环判断某个元素是否存在于该map中

```go
sliceUtil.ToMapStringStruct
```

+ sliceUtil为业务名
+ ToMap为我们想要的功能名
+ StringStruct分别为输入类型String、输出类型Struct

## 模板文件

使用 [genny](https://github.com/cheekybits/genny) 编写模板文件，然后运行命令生成各种类型的函数来模拟实现泛型

例子：字符串切片转为map[keyType]valueType

1. 创建 sliceUtil 目录

2. 创建 sliceUtil/map.go 文件

3. 编写 sliceUtil/map.go 文件函数模板

    **注意顶部的语句，KeyType为我们模板代码中使用的变量**

    ```go
    //go:generate genny -in=$GOFILE -out=gen_$GOFILE gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64"
    package sliceUtil
    
    import "github.com/cheekybits/genny/generic"
    
    type KeyType generic.Type
    
    // ToMapKeyTypeStruct KeyType slice convert as map[KeyType]struct{}
    // this method will remove duplicate elements
    func ToMapKeyTypeStruct(values []KeyType) map[KeyType]struct{} {
    	m := make(map[KeyType]struct{}, len(values))
    
    	for _, v := range values {
    		if _, ok := m[v]; !ok {
    			m[v] = struct{}{}
    		}
    	}
    
    	return m
    }
    ```

4. 在项目根目录运行命令

    ```shell
    # 执行所有文件生成代码
    $ go generate ./...
    
    # 或者使用genny
    $ genny -in=sliceUtil/map.go -out=sliceUtil/gen-map.go gen "KeyType=string,rune,int8,int16,int,int32,int64,float32,float64"
    ```

    将会生成可用的代码函数：

    ```go
    // This file was automatically generated by genny.
    // Any changes will be lost if this file is regenerated.
    // see https://github.com/cheekybits/genny
    
    package sliceUtil
    
    // ToMapStringStruct string slice convert as map[String]struct{}
    // this method will remove duplicate elements
    func ToMapStringStruct(values []string) map[string]struct{} {
    	m := make(map[string]struct{}, len(values))
    
    	for _, v := range values {
    		if _, ok := m[v]; !ok {
    			m[v] = struct{}{}
    		}
    	}
    
    	return m
    }
    
    // ToMapRuneStruct rune slice convert as map[Rune]struct{}
    // this method will remove duplicate elements
    func ToMapRuneStruct(values []rune) map[rune]struct{} {
    	m := make(map[rune]struct{}, len(values))
    
    	for _, v := range values {
    		if _, ok := m[v]; !ok {
    			m[v] = struct{}{}
    		}
    	}
    
    	return m
    }
    
    ……
    ```

## Functions
