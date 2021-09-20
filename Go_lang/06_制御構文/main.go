package main

import (
	"fmt"
	"strconv"
)

/* if文(条件分岐), error handling, for文 */

func main() {
	// a := 0
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易文付きif
	/* 注意点 :
	if文では内部の変数が優先される */
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 内部の変数が優先
	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)

	/* ---------- error handling ---------- */
	// var s string = "A"
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	/* ---------- for文 ( 繰り返し文 )---------- */
	/*
		i1 := 0
		for {
			i1++
			if i1 == 3 {
				break
			}
			fmt.Println("Loop")
		}
	*/
	/*
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/
	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue
			}
			if i == 6 {
				break
			}
			fmt.Println(i)
		}
	*/
	/*
		arr := [3]int{1, 2, 3}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/
	// index,要素を取り出す
	/*
		arr := [3]int{1, 2, 3}
		for i, v := range arr {
			fmt.Println(i, v)
		}
	*/
	/*
		// slice
		sl := []string{"python", "php", "Go"}
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	// map : python の dictionary型 / key:value
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
