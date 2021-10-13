package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

/* ---------- defer ( デファー ) ---------- */
/* 関数の終了時に実行される処理を登録することができる */
func TestDefer() {
	defer fmt.Println("END") // 最初は処理されず、最後に処理(実行)される
	fmt.Println("START")     // こちらが先に処理される
}

// defer文を複数登録した場合
// 3, 2, 1 の順番で出力される : 後から登録した式から実行される
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

/* --------------- 並列処理 go goroutin--------------- */
// sub で走らせる関数
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Microsecond)
	}
}

/* --------------- init ( package 初期化 )--------------- */
func init() {
	fmt.Println("init")
}

func main() {

	/* --------------- defer 最も使用される処理 --------------- */
	/* defer文を使用した resource の解放処理
	file作成 -> 書き込み・追記 -> Close
	*/

	file, err := os.OpenFile("Go_lang/06_制御構文/06_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go言語基礎 : 制御構文\n ")

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
	fmt.Fprintln(file, `
	--------------- 簡易文付き if ---------------

	- 注意点 :
	- if文では内部の変数が優先される`, "\n ")
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
	fmt.Fprintln(file, `
	--------------- switch ( 式スイッチ ) ---------------

	- 式を評価して処理を分岐する : data型は合わせる事`, "\n ")
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
	fmt.Fprintln(file, `
	--------------- switch ( 型スイッチ ) ---------------

	- 型参照
		- 動的に型を参照する。全ての型と互換性がある interface型を使用すると様々な引数をとる関数を作成する事ができる
		- func の引数で渡される値の型は interface型なので失われる。計算はできない`, "\n ")

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

	/* ---------- label付き for ---------- */
	/* start -> END で全ての for を抜けたい
	for {
		for {
			for {
				fmt.Println("start")
				break
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

	start ->　処理をしないの無限 Loop になってしまう */

	// label付き for ( Loopをつける事によって、処理をしないをskipできる )
	// star -> END
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("start")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")

	/* ---------- label付き for & continue ---------- */
	/* j が 1 の時だけ表示
	Loop をつけいないと処理をしないが間に入ってしまい, continue の意味をなさない */
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

	/* ---------- defer ( デファー ) ---------- */
	/* 関数の終了時に実行される処理を登録することができる */
	TestDefer()
	fmt.Fprintln(file, `
	--------------- defer ( デファー ) ---------------

	- 関数の終了時に実行される処理を登録することができる
	- 複数の処理を登録したい場合（無名関数を使用する）
		-main関数が終了後に実行される : 1, 2, 3 の順番で出力される`, "\n ")
	// 複数の処理を登録したい場合（無名関数を使用する）
	// main関数が終了後に実行される : 1, 2, 3 の順番で出力される
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	// main()関数の外で複数定義して main()関数で使用する

	RunDefer()

	/* --------------- defer で最も使用される処理 --------------- */
	// defer文を使用した resource の解放処理
	/* os.Createでfileを作成できる
	1. os.Create で fileを開く
	2. error があれば表示される(errorハンドリング)
	3. defer で file Close する
	file などの Resource 解放処理の漏れなどを防ぐ
	*/
	// file, err := os.Create("Go_lang/06_制御構文/test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// file.Write([]byte(`defer文を使用した resource の解放処理
	// os.Createでfileを作成できる
	// 1. os.Create で fileを開く
	// 2. error があれば表示される(errorハンドリング)
	// 3. defer で file Close する file などの Resource 解放処理の漏れなどを防ぐ
	// `))
	fmt.Fprintln(file, `
	--------------- defer で最も使用される処理 ---------------

	- defer文を使用した resource の解放処理
	- os.Createでfileを作成できる
		- 1. os.Create で fileを開く
		- 2. error があれば表示される(errorハンドリング)
		- 3. defer で file Close する file などの Resource 解放処理の漏れなどを防ぐ`, "\n ")
	/* --------------- panic recover --------------- */
	/* go の例外処理 : runtime を強制的に終了させる機能を持つ為
	  「 panic より推奨されている error  handling 」
		- panic は program を強制終了する為、あまり使用はしない方が良い */
	fmt.Fprintln(file, `
	--------------- panic recover ---------------

	- go の例外処理 : runtime を強制的に終了させる機能を持つ為
	「 panic より推奨されている error  handling 」
		- panic は program を強制終了する為、あまり使用はしない方が良い`, "\n ")

	/* defer で recover を登録する。panic 状態であれば x = recover に
	値が返ってくるので、その値を　x を出力して強制終了を回避
	recover で　panic から復帰できる
	recover は panic で発生した error から復帰する機能
	※　実質 defer文の中でしか機能しないので deferと一緒に使用するのが原則 */
	fmt.Fprintln(file, `
	--------------- recoverの役割 ---------------

	recover は panic で発生した error から復帰する機能
	※　実質 defer文の中でしか機能しないので deferと一緒に使用するのが原則`, "\n ")

	// // recover
	// defer func() {
	// 	if x1 := recover(); x1 != nil {
	// 		fmt.Println(x1)
	// 	}
	// }()
	// // panic : runtime error
	// panic("runtime error")
	// fmt.Println("Start")

	/* --------------- 並列処理 go goroutin--------------- */
	fmt.Fprintln(file, `
	--------------- 並列処理 go goroutin---------------

	- goroutin ( ゴルーチン )
		- go を文に使用する事で簡単に平行処理ができる
		- スレットより小さい単位
		- main()関数, sub()関数で同時に走らせるå`)

	// go sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Microsecond)
	// }

	/* --------------- init ( package 初期化 )--------------- */
	fmt.Fprintln(file, `
	--------------- init ( package 初期化 )---------------

	- main( ) 関数より先に読み込まられて処理される
	- main より先に確実に初期化処理をしたい場合に使用する
		- sub package の読み込みなどetc...
	- 複数のinit( ) を使用する事ができる
		- 記述順に処理されて行く
	`)

	fmt.Println("Main")

}
