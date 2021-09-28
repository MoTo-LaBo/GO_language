package main

/* slice, map, channel */
import (
	"fmt"
	"os"
	"time"
)

/* --------------- 可変長引数 --------------- */

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

/* --------------- channel & goroutine --------------- */

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

// func reciever1(name string, ch chan int) {

// 	for {
// 		i3, ok := <-ch
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(name, i3)
// 	}
// 	fmt.Println(name + "END")

// }

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

	// for i := 0; i < len(sls); i++ {
	// 	fmt.Println(sls[i])
	// }

	/* --------------- map / Key : value --------------- */

	fmt.Fprintln(file, `
	--------------- map / Key : value ---------------

	Key : value で値を作成して, Key : value で値を取得できる

	・Go言語は nil = None のように値が入っていないものを取得しようとすると
		- None ではなく初期値が返ってくる
		- 登録されていないモノは初期値が返ってくる

	・Go言語の基本型は nil のように特殊な値を持たないので
		- 何も存在しないまま処理を実行してしまうミスも起こりやす

	・上記を踏まえ map() には error  handling の機能も付いている
		- ok を使用して値が取得出来たかどうかを True or False で判定する
		- if !ok { 処理 fmt.Println("error")} のように
			- if文を使用して、値が取り出せなかった時の処理を記述
	`)

	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make & map
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	// 値の取り出し | Key:value
	fmt.Println(m["A"])

	fmt.Println(m4[2])

	// 登録されていない値は初期値が返ってくる
	fmt.Println(m4[3])

	// error  handling
	s := m4[1]
	fmt.Println(s)

	// ok をつける事によって取得出来たかどうかを True or False で判定できる
	// s, ok := m4[1]
	// fmt.Println(s, ok)

	// index 3番目には何もはいってい無いからの状態なので取得出来なかったとして、 False
	// _, ok := m4[3]
	// if !ok {
	// 	fmt.Println("error")
	// }
	// fmt.Println(s, ok)

	// 値の変更・上書き・追加
	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "EU"
	fmt.Println(m4)

	// 削除 : delete(削除 map, Key)
	delete(m4, 3)
	fmt.Println(m4)

	// len()
	fmt.Println(len(m4))

	/* --------------- map & foe --------------- */
	m5 := map[string]int{
		"Apple": 100,
		"Banan": 200,
		"Pache": 300,
	}
	// for k, v := range m5 {
	// 	fmt.Println(k, v)
	// }

	// for _, v := range m5 {
	// 	fmt.Println(v)
	// }

	for k := range m5 {
		fmt.Println(k)
	}

	/* --------------- channel --------------- */
	fmt.Fprintln(file, `
	/* --------------- channel --------------- */

	複数のゴルーチン間でのdataの受け渡しをする為に設計されたdata構造
		1. 宣言
		2. 操作

	・channel は data の送受信を行うdata構造
		- sub type を使用して送信・受信を明示的に指定する事もできる
		- 指定しない場合は双方向のchannel になる
		- 複数のgoroutine(ゴルーチン)間でのやりとりなので ch1(1つだけ)では deadloc!

	・make(chan 型)関数で channel として機能を持たせる事ができる
		- make によって初期化が行われて、書き込み・読み込みができるようになる

	・cap()
		- 容量(data size / buffer size)を調べる事ができる
		- バッファサイズを指定して作成する事もできる

	・queue (キュー) = buffer size
		- 最初に投入したデータが最初に出力される
		- 「First-in, First-out」（FIFO）型のデータ構造
		- dataを取り出す順番が保証される

		- buffer size を超えた量のdataを送った場合...
			- fatal error : all goroutines are asleep - deadlock!
	`)
	/*
		var ch1 chan int

		ch1 = make(chan int)

		ch2 := make(chan int)

		fmt.Println(cap(ch1))
		fmt.Println(cap(ch2))

		ch3 := make(chan int, 5)
		fmt.Println(cap(ch3))

		// 受信専用の channel
		// var ch2 <-chan int

		// 送信専用の channel
		// var ch3 <-chan int

		// data を送信
		ch3 <- 1
		fmt.Println(len(ch3))

		ch3 <- 2
		ch3 <- 3
		fmt.Println("len", len(ch3))

		// data の受信
		i := <-ch3
		fmt.Println(i)
		fmt.Println("len", len(ch3))

		i2 := <-ch3
		fmt.Println(i2)
		fmt.Println("len", len(ch3))

		fmt.Println(<-ch3)
		fmt.Println("len", len(ch3))

		// buffer size を超えるdataを送ってみる
		ch3 <- 1
		fmt.Println(<-ch3)
		ch3 <- 2
		ch3 <- 3
		ch3 <- 4
		ch3 <- 5
		ch3 <- 6
		fmt.Println("len", len(ch3))
	*/
	/* --------------- channel & goroutine --------------- */
	// channel は複数のgoroutine(ゴルーチン)間で機能する

	ch1 := make(chan int)
	ch2 := make(chan int)
	// // 送信先の channel がないので deadlock !
	// fmt.Println(<-ch1)

	// reciever を起動して並行で走らせながら channel の共有を行う
	go reciever(ch1)
	go reciever(ch2)

	// channel に data を送る : 並行処理なので同じものが２つ出力される
	i := 0
	for i < 3 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

	/* --------------- channel close --------------- */
	// 送受信を終えた channel を明示的に止める事ができる

	fmt.Fprintln(file, `
	--------------- channel close ---------------

	Close された channel から送信はできない。受信はできる
		- ※ 受信は初期値 (0の状態) が返ってくる

	i, ok で Close された ch の値を確認すると
		- i : 初期値 , ok : boolean (True or False)
		- ch の open状態が返ってくる
		- 厳密には...
			channel の buffer内が空で chが Close された状態 = False
			値を送信後 <- Close では buffer に値が受信(入る)ので
			Close されていても True が返されるj
	`)

	// close された channel から送信することはできな
	// ch3 := make(chan int, 2)

	// ch3 <- 1
	// close(ch3)

	// 受信はできる : 初期値(0の状態)が返ってくる
	// fmt.Println(<-ch3)

	// 第二引数は channel の open状態を boolean(True or False)で表す
	// i, ok := <-ch3
	// fmt.Println(i, ok)

	// i2, ok := <-ch3
	// fmt.Println(i2, ok)

	// go reciever1("1.goroutine", ch3)
	// go reciever1("2.goroutine", ch3)
	// go reciever1("3.goroutine", ch3)

	// i3 := 0
	// for i3 < 5 {
	// 	ch3 <- i3
	// 	i3++
	// }
	// close(ch3)
	// time.Sleep(3 * time.Second)

	/* --------------- channel for文 --------------- */
	// for 文で回した時は必ず　close する　deadloc になってしまう

	// ch4 := make(chan int, 3)
	// ch4 <- 1
	// ch4 <- 2
	// ch4 <- 3
	// close(ch4)
	// for i4 := range ch4 {
	// 	fmt.Println(i4)
	// }

	/* --------------- channel select --------------- */

	fmt.Fprintln(file, `
	--------------- channel select ---------------

	複数の channel を処理する時どちらかの ch に error が出た時に他の program が止まってしまう
		- select文はそれを防いでくれる
		- 複数の goroutine を制御してくれる

	select
		- select case の内部は channel の処理
			- ch 出ないと error になってしまう
		- 送受信はどちらでも大丈夫
		- switch文の case とは少し違う
			- 最初に成立した case が優先されるのではなくランダムに実行される
		- default
			- case に当てはまらない時に適用される

		select文を使用する事で適切に非同期処理ができるようになっている
	`)

	ch3 := make(chan int, 2)
	ch4 := make(chan string, 2)

	ch3 <- 1
	ch4 <- "A"

	// v := <-ch3
	// v2 := <-ch4
	// fmt.Println(v)
	// fmt.Println(v2)

	// 分岐してくれる
	select {
	case v1 := <-ch3:
		fmt.Println(v1 + 1000)
	case v2 := <-ch4:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch5 := make(chan int)
	ch6 := make(chan int)
	ch7 := make(chan int)
	// reciever : 始まり(基準)で値を入れて送信
	go func() {
		for {
			i := <-ch5
			ch6 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch6
			ch7 <- i2 - 1
		}
	}()

	// channel に data を渡す
	n := 0
L:
	for {
		select {
		case ch5 <- n:
			n++
		case i3 := <-ch7:
			fmt.Println("recieved", i3)
		default:
			if n > 10 {
				break L
			}
		}
		// if n > 10 {
		// 	break
		// }
	}

}
