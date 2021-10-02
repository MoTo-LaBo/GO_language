package main

/* pointter (ポインタ) */
import (
	"fmt"
	"os"
)

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func Double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	file, err := os.OpenFile("Go_lang/08_ポインタ/08_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語　:　pointer(ポインター)", `
		--------------- defer 最も使用される処理 ---------------

	defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close


			--------------- pointer とは? ---------------

			pointer とは、値型に分類される data 構造(基本型, 参照型, 構造体)
			のメモリ上のアドレスと型の情報

		- Goではこれを使用して data構造を間接的に参照、操作ができるようになっている
		- 主に値型の操作に使用され、参照型は元からその機能を含んでいる為に基本的には不要


			--------------- pointer型 ---------------

			アドレス・dataを共有する事ができる

	・ & （アンバサンド）はアドレスを指す
		- adress 演算子
		- 値に & を付けることによって adress演算子になる（参照渡し）

	・ ＊（アスタリスク）をつける事で pointer型を宣言する事ができる
		- 代入するのは, 値のadress　=　&n
		- ＊値　=　n　は ＝(equal)になる　※ 参照渡し

	・　＊（値）を書き換えると n に反映される
		- 逆もできる。 n を書き換えると, ＊（値）の実態 data も変わる

	・　＊（アスタリスク）を付ける = 実態参照
		- 現在 data を確認できる
			- pointer が保持している adress を経由して data 本体を参照できる

	・ 参照型はもともと参照渡しという機能を持っているので pointer をつかわなくても値を書き換える事ができる

	`)

	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	// pointer型
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p) // * (アスタリスク)を付けることにより,アドレスが指し示す実態を表示

	// 上記で参照渡しをしているので, *pを書き換えると n に反映される
	// *p = 300
	// fmt.Println(n)
	// n = 200
	// fmt.Println(*p)

	Double2(&n)
	fmt.Println(n)

	Double2(p)
	fmt.Println(*p)

	// slice : 参照型はもともと参照渡しという機能を持っているので pointer をつかわなくても良い
	var sl []int = []int{1, 2, 3}
	Double3(sl)
	fmt.Println(sl)

}
