// In Go, an _array_ is a numbered sequence of elements of a
// specific length.

package main

import "fmt"

func main() {

    // Here we create an array `a` that will hold exactly
    // 5 `int`s. The type of elements and length are both
    // part of the array's type. By default an array is
    // zero-valued, which for `int`s means `0`s.
    // 声明数组时需要指定内部存储的数据类型，以及需要存储的元素的数量，这个数量也称为数组的长度
    // 声明一个包含5个元素的整型数组
    var a [5]int
    fmt.Println("emp:", a)

    // We can set a value at an index using the
    // `array[index] = value` syntax, and get a value with
    // `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // The builtin `len` returns the length of an array.
    fmt.Println("len:", len(a))

    // Use this syntax to declare and initialize an array
    // in one line.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 声明数组并指定特定元素的值
    // 声明一个有5个元素的数组
    // 用具体初始化索引为1和2的元素
    // 其余元素保持零值
    array := [5]int{1: 10, 2: 20}
    fmt.Println("array:", array)

    // Array types are one-dimensional, but you can
    // compose types to build multi-dimensional data
    // structures.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
