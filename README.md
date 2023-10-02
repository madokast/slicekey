# slicekey

Make a slice that can be used as a key of a map

可以当作 map key 的 slice

The slice-key is an array object, created by reflect in runtime, holding the same elements as the slice. 

原理是利用反射动态构建一个数组对象，将 slice 元素复制到数组，从而可以作为 key 使用。

## usage

使用 Of() 函数创建一个 slicekey

```go
	m := map[Slice[int]]string{}

	key := make([]int, 0)
	m[Of(key...)] = "empty"

	key1 := make([]int, 1)
	key1[0] = 1
	m[Of(key1...)] = "1-"

	key2 := make([]int, 2)
	key2[0] = 1
	key2[1] = 3
	m[Of(key2...)] = "1-3"

	key2[0] = 3
	key2[1] = 1
	m[Of(key2...)] = "3-1"

	m[Of(1)] = "1-modify"

	fmt.Printf("%v\n", m) // map[{[1]}:1-modify {[]}:empty {[1 3]}:1-3 {[3 1]}:3-1]

	fmt.Printf("m[1, 3] = %s\n", m[Of(1, 3)]) // m[1, 3] = 1-3
	fmt.Printf("m[3, 1] = %s\n", m[Of(3, 1)]) // m[3, 1] = 3-1
```

另一段实例

```go
	m := make(map[Slice[int]]int)

	for i := 100; i < 102; i++ {
		m[Of(i)] = i
		for j := 20; j < 23; j++ {
			m[Of(i, j)] = i + j
		}
	}

	fmt.Printf("%v\n", m) // map[{[100]}:100 {[101]}:101 {[100 20]}:120 {[100 21]}:121 {[100 22]}:122 {[101 20]}:121 {[101 21]}:122 {[101 22]}:123]
```