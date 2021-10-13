package main

import (
	"fmt"
	"os"
)

/* ---------------- interface型 アサーション ---------------- */
// 最もポピュラーな使い方, 異なる型に共通の性質を付与する

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Nname=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

/* ---------------- カスタム error (interface) ---------------- */
// type error interface {
// 	Error() string
// }
// 上記のように package に表記されている

// 独自の error を作成する
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

/* 上記のように記述する事により, Go の Error と言う interface型と共通の
性質を持つものとして扱えるようになる */

// error を発生させる関数
func RaiseError() error {
	return &MyError{Message: "カスタムerrorが発生しました", ErrCode: 1234}
}

/* ---------------- Stringer (interface) ---------------- */
// type Stringer interface {
// 	String() string
// }
// 上記のように package に表記されている

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {

	// Stringer
	p := &Point{100, "ABC"}
	fmt.Println(p)

	// 独自の error
	err := RaiseError()
	fmt.Println(err.Error())

	// 型参照で復元してフィールドに access する
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	// file の書き出し
	file, err := os.OpenFile("Go_lang/10_interface/10_doucument.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go言語基礎 : interface\n", `
	--------------- defer 最も使用される処理 ---------------

	defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close


	--------------- interface型 ---------------

	・ Stringfy と言う interface型でまとめて扱うことができる
		- for文で回して共通の method を使用できる

		- 型の性質を抽出した interface を使用すると Go の厳密な型の system に
		柔軟性を与える事ができる


	--------------- カスタムerror (interface) ---------------

	・ Go の Error() と言う interface型と共通の性質を持つものとして扱えるようになる

	・ 変数の error は program 上はあくまで Error型の変数なので, 独自でカスタムした
		フィールドには access できない
		- フィールドに access する場合は, 型参照を使用して復元してフィールドに access する


	--------------- Stringer (interface) ---------------

	・ fmt package に定義されている Stringer型
		- simple 文字列を返す

	`)

	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}
