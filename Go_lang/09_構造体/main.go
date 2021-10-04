package main

import (
	"fmt"
	"os"
)

/* --------------- struct 独自型 ----------------- */
// 独自の integer型を生成する事ができる
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

/* --------------- struct 埋め込み--------------- */
/* 構造体はフィールドに構造体を埋め込む事ができる : フィールド名と型名が
　一緒になる事はよく見られる形式 */
type T struct {
	User User
	// User // 型名を省略できる
}

func (u *User) SetName3() {
	u.Name = "A"
}

/* --------------- 構造体の定義 ----------------- */
type User struct {
	Name string
	Age  int
	// X, Y int
}

// func (u User) SayName() {
// 	fmt.Println(u.Name)
// }

/* --------------- struct型 slice ----------------- */
// User を slice として扱いたい場合
type Users []*User

// 下記のようにも記述できるが、上記のような記述の方が望ましい
// type Users struct {
// 	Users []*Users
// }

/* --------------- struct型 コンストラクタ ----------------- */

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

/* --------------- struct method　--------------- */
// func (u User) SetName(name string) {
// 	u.Name = name
// }

// func (u *User) SetName2(name string) {
// 	u.Name = name
// }

/* 構造体を参照渡ししたい場合に pointer を使用する */

// // 普通の関数
// func UpdataUser(user User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// // pointer型の関数
// func UpdataUser2(user *User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

func main() {
	file, err := os.OpenFile("Go_lang/09_構造体/09_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語 : struct()", `
		--------------- defer 最も使用される処理 ---------------

	defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close


			--------------- struct method ---------------

			任意の型に特化した関数を定義する仕組み

	・ 基本原則は method のレシーバーは pointer型(参照渡し)にしておく事がのぞましい

	・ 呼び出し側は pionter にする必要はない
		- 全ては method のレシーバーが参照渡し or 値渡しで挙動が決まるので


			--------------- struct method ---------------

	・ 構造体はフィールドに構造体を埋め込む事ができる
		- フィールド名と型名が一緒になる事はよく見られる形式

	・ 構造体だけ記述して型名を省略できる
		- 型名省略の場合のに直接構造体のフィールドに acssce できる


			----------- struct型 コンストラクタ -------------

			Go言語にはコンストラク機能はないが、パターンとして
			コンストラクタ関数を使用する事が良くある

	・ コンストラクタ関数を使用して User型の pointer を生成する
		- Goでは良く使用される
			--------------- struct 独自型 -----------------

	・　特有の機能を持った独自な型や独自のmethodをtypeを用いて作成する事ができる

 `)

	// // 明示的定義
	// var user1 User
	// fmt.Println(user1)
	// user1.Name = "user1"
	// user1.Age = 10
	// fmt.Println(user1)

	// // 暗黙的定義
	// user2 := User{}
	// fmt.Println(user2)
	// user2.Name = "user2"
	// fmt.Print(user2, "\n")

	// user3 := User{Name: "user3", Age: 30}
	// fmt.Println(user3)

	// // 構造体のフィールドの順番に値を代入する事
	// user4 := User{"user4", 40}
	// fmt.Println(user4)

	// // error になる
	// // user5 := User{30, "user5"}
	// // fmt.Println(user5)

	// user6 := User{Name: "user6"}
	// fmt.Println(user6)

	// /* --------------- pointer型で宣言 --------------- */
	// /* pointer型として使用する場合は、&(アンバサンド)をつけて使用する頻度が高い
	// 関数の引数として構造体の変数を渡す場合に使用される */
	// // new() : 構造体の pointer を返すようになっている
	// user7 := new(User)
	// fmt.Println(user7)

	// // adress演算子を使用する事で = new()と同様な事ができる
	// user8 := &User{}
	// fmt.Println(user8)

	// /* --------------- pointer型関数と関数の違い --------------- */

	// // どちらも初期値のまま
	// UpdataUser(user1)  // 普通の関数
	// UpdataUser2(user8) // pointer型の関数

	// fmt.Println(user1) // 値渡し  : copy生成 -> 新たな構造体
	// fmt.Println(user8) // 参照渡し : adress -> そのまま上書き

	/* --------------- struct method　--------------- */
	// user1 := User{Name: "user1"}
	// user1.SayName()

	// // 普通の関数　値渡しの関数
	// user1.SetName("A")
	// user1.SayName()
	// // pointer型　参照渡しの関数
	// user1.SetName2("A")
	// user1.SayName()

	// user2 := &User{Name: "user2"}
	// user2.SetName2("B")
	// user2.SayName()

	/* --------------- struct 埋め込み--------------- */

	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User.Name)

	// User型を省略した場合のみ : 直接 Nameフィールドにアクセスできる
	// fmt.Println(t.Name)
	t.User.SetName3()
	fmt.Println(t)

	/* --------------- struct型 コンストラクタ ----------------- */

	// user1 := NewUser("user1", 10)
	// fmt.Println(user1)  // pointer型で表示
	// fmt.Println(*user1) // 実態アクセス表示

	/* --------------- struct型 slice ----------------- */

	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	// fmt.Println(users)          // pointer型なので各userの adress が表示される

	// 各data の事態を for loop で表示させる事ができる
	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}

	/* --------------- struct型 map ----------------- */

	// map の中にstruct(構造体)
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	fmt.Println(m)

	// Key として User を使用
	m2 := map[User]string{
		{Name: "user3", Age: 30}: "Tokyo",
		{Name: "user4", Age: 40}: "LA",
	}
	fmt.Println(m2)

	// make を使って生成
	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}

	/* --------------- struct 独自型 ----------------- */

	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi) // 型を調べる

	/* 他のdata型と計算されないようにする為にあえて MyInt を定義して使用 */
	// i := 100
	// fmt.Println(mi + i)

	mi.Print()
}
