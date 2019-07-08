package main

// A goroutine is a lightweight thread of execution

import "fmt"

func f(from string)  {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main()  {
	/*
	Suppose we have a function call f(s).
	Here’s how we’d call that in the usual way, running it synchronously.
	 */
	f("direct1")
	f("direct2")
	f("direct3")
	f("direct4")
	f("direct5")

	/*
	To invoke this function in a goroutine, use go f(s).
	This new goroutine will execute concurrently with the calling one.
	*/
	go  f("goroutines1")
	go  f("goroutines2")
	go  f("goroutines3")
	go  f("goroutines4")
	go  f("goroutines5")

	/*
	You can also start a goroutine for an anonymous function call.
	*/
	// 匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}

