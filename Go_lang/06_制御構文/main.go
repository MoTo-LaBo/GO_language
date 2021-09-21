package main

import (
	"fmt"
	"strconv"
)

/* if文(条件分岐), error handling, for文, switch(式・型) */

/* ---------- switch ( 型スイッチ )---------- */
/* 型参照
・動的に型を参照する。全ての型と互換性がある inter face型を
 使用すると様々な引数を取る関数を作成する事ができる
・ anything func の引数で渡された値の型は interface型なので
  失われる。計算などはできない
*/
func anything(a interface{}) {
	// fmt.Println(a)  // 型による switch で値が使用できる
	switch v := a.(type) {
	case string:
		fmt.Println(v, "!?")
	case int:
		fmt.Println(v + 1000)
	}
}

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
	// x := 0
	// if x := 2; true {
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

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

	/* ---------- switch ( 式スイッチ )---------- */
	// 式を評価して処理を分岐する : data型は合わせる事
	/*
		// n := 1
		n := 3
		switch n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("I don't Know")
		}
	*/

	// 変数の局所性を使用できる : switch 内だけで使用できる変数
	/*
		switch n := 2; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("(I don't know")
		}
	*/

	// 式を使用して switchを記述できる
	/* 下記のように記述する場合は、swich文に与える式は省略できる
	上記の列挙型と下記の演算型の式は混在できないので注意する error になる */
	/*
		n := 6
		switch {
		case n > 0 && n < 4:
			fmt.Println("0 < n < 4")
		case n > 3 && n < 7:
			fmt.Println("3 < n < 7")
		}
	*/

	/* ---------- switch ( 型スイッチ )---------- */
	/* 型参照
	・動的に型を参照する。全ての型と互換性がある inter face型を
	 使用すると様々な引数を取る関数を作成する事ができる
	・ anything func の引数で渡された値の型は interface型なので
	 失われる。計算などはできない
	*/
	anything("aaa")
	anything(1)
	// 型参照
	var x interface{} = 3
	i2, isInt := x.(int)
	fmt.Println(i2, isInt)
	// fmt.Println(i2 + 2)
	// fmt.Println(x + 2)  // interface型とint型なので計算できない
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// 型参照 : True になった条件文で出力される
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	// switch文 && 型参照
	/* 上記の if文を使用するより簡略的かつ分かり易いので
	switch文を使用した記述の仕方が推奨されている */
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}
	// 値も使用したい場合 (上記は復元した値を使用する事ができない)
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

}
