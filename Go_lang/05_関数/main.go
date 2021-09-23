package main

import (
	"fmt"
	"os"
)

/* 関数 function, 無名関数, 関数を返す関数,
　　関数を引数に取る関数, クロージャー, ジェネレーター  */

// 単一
func Plus(x, y int) int {
	return x + y
}

// 複数
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値を示す関数 (返り値の変数を指定できる)
func Double(price int) (result int) {
	result = price * 2
	return // 返り値を最初で指定しているので省略できる
}

/* ---------- 無名関数 ---------- */
func Noreturn() {
	fmt.Println("No Return")
	return
}

/* ---------- 関数を返す関数 ---------- */
func Returnfunc() func() {
	return func() {
		fmt.Println("I'm function")
	}
}

/* ---------- 関数を引数に取る関数 ---------- */
func CallFanction(fp func()) {
	fp()
}

/* ---------- 関数閉包(クロージャー : closur) ---------- */
/* storeで空文字を作成 -> s へ格納 -> next で store を上書き ->
返す値は s -> 最初の定義の段階(1つ目)は空文字として入っているので
1個遅れた形で入力された文字列が出力される。それの繰り返しなので
最後の文字列も出力されない */
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

/* ---------- ジェネレーター : generator ---------- */
/* クロージャーを応用する事によって generator のように振る舞う事ができる
変数を変える事で新たに生成される。クロージャー間で i は共有されない */
func intergers() func() int {
	i1 := 0
	return func() int {
		i1++
		return i1
	}
}

func main() {
	/* --------------- defer 最も使用される処理 --------------- */
	/* defer文を使用した resource の解放処理
	file作成 -> 書き込み・追記 -> Close
	*/
	file, err := os.OpenFile("Go_lang/05_関数/05_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語基礎 : 関数", "\n ")
	fmt.Fprintln(file, `
	---------- 関数閉包(クロージャー : closur) ----------

	- storeで空文字を作成 -> s へ格納 -> next で store を上書き ->
	 返す値は s -> 最初の定義の段階(1つ目)は空文字として入っている
		- 1個遅れた形で入力された文字列が出力される
		- それの繰り返しなので最後の文字列も出力されない`, "\n ")

	fmt.Fprintln(file, `
	---------- ジェネレーター : generator ----------

	- クロージャーを応用する事によって generator のように振る舞う事ができる
		- 変数を変える事で新たに生成される
		- クロージャー間で同じ変数は共有されない`)

	i := Plus(1, 2)
	fmt.Println(i)

	// i2, i3 := Div(9, 3)
	// fmt.Println(i2, i3)

	// 返り値を使用したくな(破棄) _ (アンダーバー)を使用する
	i2, _ := Div(9, 3)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()

	/* 無名関数 : 名前が無い所を除けば普通の関数と本質的にはほぼ同じ */
	f1 := func(x, y int) int {
		return x + y
	}
	i5 := f1(1, 2)
	fmt.Println(i5)

	// 変数を定義しなくてもそのまま出力できる
	f2 := func(x, y int) int {
		return x + y
	}(3, 2)
	fmt.Println(f2)

	// 関数を返す関数
	rf := Returnfunc()
	rf() // Returnfunc で返す無名関数の中身が返ってくる

	// 関数を引数に取る関数
	CallFanction(func() {
		fmt.Println("I'm paramerter function")
	})

	//  関数閉包(クロージャー : closur)
	fc := Later()
	fmt.Println(fc("Hello"))
	fmt.Println(fc("Golang"))
	fmt.Println(fc("Let's Go"))
	fmt.Println(fc("Go"))

	// ジェネレーター : generator
	ints := intergers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	// 変数を変える事で新たに生成される
	otherints := intergers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())

}
