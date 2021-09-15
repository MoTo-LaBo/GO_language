package main

import (
	"fmt"
)

/* 変数 */

// i5 := 500
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	/* 明示的な定義 */
	// var 変数名　型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	/* 変数だけを作成 */
	var i3 int          // 初期値 0
	var s3 string       // 初期値　空文字(表示はされていないが値はある)
	fmt.Println(i3, s3) // 型だけの場合は初期値が入る

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	/* 値の更新 */
	i = 150
	fmt.Println(i)

	/* 暗黙的な定義 */
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	// 値の更新
	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)
	// 変数は int 型なので string を割り当てる事ができない

	// 明示的な定義であれば関数外でも定義できる
	fmt.Println(i5)

	outer()

	// s4 は outer 関数の中だけ有効
	// fmt.Println(s4)

	// var s5 string = "Not use"
}
