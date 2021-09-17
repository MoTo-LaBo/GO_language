package main

/* 型 : int, float, uint, coplex, bool, string, byte */
import (
	"fmt"
)

func main() {

	/* ---------- int　(integer : 整数) ---------- */
	// 環境依存 :　pc環境によって　int8,16,32,64 という風に数値が変わる
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)
	// fmt.Println(i + i2) : int値が合わないと計算できない
	// 型変換
	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2)) // 型変換：引数に渡す
	// 型変換後
	fmt.Println(i + int(i2))

	/* ----------- 浮動小数点型 ( float ) ---------- */
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2 // 暗黙的な定義の場合は自動で float64 ※ 環境依存では無い
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32) // float32 は基本的には使用しない

	// 型変換
	fmt.Printf("%T\n", float64(fl32))

	// 非数特殊な値を保持する場合がある
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	var u8 uint = 255 // uint型 (+の整数型)
	fmt.Println(u8)

	var c64 complex64 = -5 + 12i // complex型 (複素数型)
	fmt.Println(c64)

	/* ---------- bool　(boolean : 論理値型) ---------- */

	var t, f bool = true, false
	fmt.Println(t, f)

	/* ---------- string　(string : 文字列型) ---------- */
	// 文字列はバイト配列の集まり

	var s string = "Hello Golang"
	fmt.Printf(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Printf(si)
	fmt.Printf("%T\n", si)

	fmt.Printf(`test
	test
		test`)

	fmt.Printf("\"\n")
	fmt.Printf(`""`)

	fmt.Println("\n", s[0])         // バイト型　： 72 と表示される
	fmt.Println("\n", string(s[0])) // バイト型 -> str型に変換　: H と表示される

	/* ---------- byte　(uint8 : byte型) ---------- */

	byteA := []byte{72, 73} // slice : 配列を表現できる型
	fmt.Println(byteA)

	// byte型 -> 文字列へ変換 : ASCII で表している
	fmt.Println(string(byteA))
	// 文字列:ASCII -> byte型
	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

	/* -------------------- 配列型 -------------------- */

	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0] の指定した要素数の配列が作成される
	fmt.Printf("%T\n", arr1) // [3]int -> 型

	// 明示的な定義 : 空文字が入る = [ A B  ]
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的な定義
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数の宣言を省略 : 要素の数を自動で計算してくれる = ... (ドット３つ)
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4) // [2]string : 入れる要素数を変えると数も変わる

	// 値の取り出し : index を指定
	fmt.Println(arr1[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2-1])

	// 値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	/* 中の要素が同じでも int = int -> 要素数 3 x 要素数　4
	要素数の数が合わないと代入はできない error になる*/
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 配列の要素数を調べる : lengs
	fmt.Println(len(arr1))

	/* ---------- interface & nill ---------- */

	/* {}を含めて１つの型。
	あらゆる型と互換性がある特殊な型
	int, string, float, bool... 全て
	初期値 : nil -> python の NaN にあたる(何も値を持たない)
	あくまで全ての型を汎用的に表す手段 */

	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "Golang"
	fmt.Println(x)
	x = true
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// data 特有の計算や演算はできないので注意
	// x = 2
	// fmt.Println(x + 3)

}
