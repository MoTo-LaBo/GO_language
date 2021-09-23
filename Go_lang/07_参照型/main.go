package main

/* slice, map, channel */
import (
	"fmt"
	"os"
)

/* --------------- defer 最も使用される処理 --------------- */
/* defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close */

func init() {
	file, err := os.OpenFile("Go_lang/07_参照型/07_document.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Go 言語基礎 : 参照型\n ", `
	--------------- defer 最も使用される処理 ---------------

	defer文を使用した resource の解放処理 : file作成 -> 書き込み・追記 -> Close`)

}

func main() {

}
