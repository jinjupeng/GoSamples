package main

import "fmt"
// range iterates over elements in a variety of data structures.
func main()  {

	// 初始化切片
	nums := []int{2, 3, 4}
	sum := 0
	// 下划线表示的是返回索引下标值
	// range to sum ths numbers in s slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides
	// both the index and value for each entry.
	// Above we didn’t need the index,
	// so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for i, num := range  nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range  kvs {
		fmt.Println("key:value->", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first values is starting byte index of the rune
	// and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i,c)
	}
}
