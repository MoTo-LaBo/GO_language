package main

import "fmt"

/* 算術演算子 */

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "ED")

	fmt.Println(1 - 2)

	fmt.Println(5 * 4)

	fmt.Println(60 / 5)

	fmt.Println(9 % 4)

	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)

	n++ // n = n + 1
	fmt.Println(n)

	n-- // n = n - 1
	fmt.Println(n)

	s := "ABC"
	s += "DE"
	fmt.Println(s)

	/* -------------------- 比較演算子 -------------------- */

	fmt.Println(1 == 1)

	fmt.Println(1 == 2)

	fmt.Println(4 <= 8)

	fmt.Println(1 >= 6)

	fmt.Println(1 < 2)

	fmt.Println(true == false)
	fmt.Println(true != false)

	/* -------------------- 論理演算子 -------------------- */

	// true かつ true であれば　true
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	// 左もしくは右が　true であれば　true
	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	// 論理値の反転
	fmt.Println(!true)
	fmt.Println(!false)

}
