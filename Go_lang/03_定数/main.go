package main

import (
	"fmt"
	"os"
)

/* 定数 */
/*
   定数は変更できない : 関数内にはほとんど記述しない
   他の package からも読み込めるように グローバル環境に記述し
   頭文字を大文字にする事により他の package からも読み込める
   ・パブリック, プライベート
   ・グローバルに宣言して小文字で定数を設定してしまうと
	※ 現在の package内でしか使用できなくなる

	・大文字, 小文字の使い分け (頭文字)
		- 他で呼び出す定数　＝　大文字
		- 現在packageだけ　＝　小文字
*/

const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

/* overflows int で error になる
   var Big int = 9223372036854775807 + 1
   定数の場合は int型の最大値を定義して表示する事ができる
   ※ 変数の場合は型の最大値超えて定義する事はできない
*/
const Big = 9223372036854775807 + 1

// 連続する整数の連番を表示させる
const (
	c0 = iota
	c1
	c2
)

func main() {

	/* --------------- defer 最も使用される処理 --------------- */
	/* defer文を使用した resource の解放処理
	file作成 -> 書き込み・追記 -> Close
	*/
	file, err := os.OpenFile("Go_lang/03_定数/03_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語基礎 : 定数\n ")

	fmt.Fprintln(file, `
	--------------- 定数 ---------------

	- 定数は変更できない : 関数内にはほとんど記述しない
		- 他の package からも読み込めるように グローバル環境に記述し
		- 頭文字を大文字にする事により他の package からも読み込める
	・パブリック, プライベート
	・グローバルに宣言して小文字で定数を設定してしまうと
	 	※ 現在の package内でしか使用できなくなる

	 ・大文字, 小文字の使い分け (頭文字)
		 - 他で呼び出す定数　＝　大文字
		 - 現在packageだけ　＝　小文字`, "\n ")

	fmt.Fprintln(file, `
	---------------  overflows int ---------------

	- var Big int = 9223372036854775807 + 1
		定数の場合は int型の最大値を定義して表示する事ができる
		※ 変数の場合は型の最大値超えて定義する事はできない`)

	fmt.Println(Pi)

	// 一度定義した定数は上書きはできない
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL)
	fmt.Println(SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)

}
