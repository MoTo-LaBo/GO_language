package main

/* slice, map, channel */
import (
	"fmt"
	"os"
)

/* --------------- 可変長引数 --------------- */

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	file, err := os.OpenFile("Go_lang/07_参照型/07_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語基礎 : 参照型\n ", `
	--------------- defer 最も使用される処理 ---------------

	defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close


	--------------- slice ---------------

	・配列と同様に中の各括弧に記述できる
		- 配列との違いは [ ] の中に要素を記述しない

	・ 参照型は最初に定義した値と後から代入した値は
	　　同じメモリを共有するという特徴がある

	make([]型, 数)
		- 指定したdata型を指定した数で出力


	--------------- append, make, len, cap ---------------

	append(追加slice, 追加data)
			- slice の拡張(dataが追加される)

	・sliceは要素数が可変調になっている
		- 拡張性が高い(要素数制限がなくてサイズ数も変更できる)

	len(要素)
		- 要素数を調べる

	cap(要素)
		- capacity : 要素数を調べる
		- program の容量を気にする場合に使用
		- パフォーマンスを気にする開発場合に細かく指定する

	【capacity（要領）について】
		- 要領以上の要素が追加されるとメモリの消費が倍になってしまう
		- メモリーを気にするような開発の場合は、容量にも気をつける
		- 過剰にメモリを確保してしまうと実行速度が落ちたりする
		- 良質なパフォーマンスを実現するには、要領の管理も気にする


	--------------- 可変長引数 ---------------

	引数に数を指定しなで渡すことができる



	`)
	/* --------------- 可変長引数 --------------- */

	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8))

	fmt.Println(Sum())

	slsum := []int{1, 2, 3}
	fmt.Println(Sum(slsum...))

	/* --------------- slice --------------- */

	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	// 暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make() :
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し : index番号指定
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	fmt.Println(sl5[2:4])

	fmt.Println(sl5[:2])

	fmt.Println(sl5[2:])

	fmt.Println(sl5[:])

	fmt.Println(sl5[len(sl5)-1])

	fmt.Println(sl5[1 : len(sl5)-1])

	/* --------------- append, make, len, cap --------------- */

	// sl := []int{100, 200}
	// fmt.Println(sl)

	// sl = append(sl, 300)
	// fmt.Println(sl)

	// sl = append(sl, 400, 500, 600)
	// fmt.Println(sl)

	// // make
	// sl2 := make([]int, 5)
	// fmt.Println(sl2)

	// // len() : 要素数を調べる
	// fmt.Println(len(sl2))

	// // cap :　capacity
	// fmt.Println(cap(sl2))

	// // make
	// sl3 := make([]int, 5, 10)
	// // len() : 要素数を調べる
	// fmt.Println(len(sl3))

	// // cap :　capacity メモリ容量
	// fmt.Println(cap(sl3))

	// sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	// // len() : 要素数を調べる
	// fmt.Println(len(sl3))

	// // cap :　capacity メモリ容量
	// fmt.Println(cap(sl3))

	/* --------------- cap --------------- */
	// 参照型は最初に定義した値と後から代入した値は同じメモリを共有するという特徴がある
	// slice, map, channel も同じでメモリが共有される
	/*
		sl := []int{100, 200}
		sl2 := sl

		sl2[0] = 1000
		fmt.Println(sl)

		// 基本型の場合は値を渡してもメモリは共有されないので、値は変わらない
		var i int = 10
		i2 := i
		i2 = 100
		fmt.Println(i, i2)
	*/

	// copy(cp先, cp元) :
	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)
	// n := copy(sl2, sl)
	// fmt.Println(n, sl2)

	// slice : 文字列型
	sls := []string{"A", "B", "C"}
	fmt.Println(sls)

	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	for i := 0; i < len(sls); i++ {
		fmt.Println(sls[i])
	}

}
