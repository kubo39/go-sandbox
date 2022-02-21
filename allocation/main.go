package main

import "fmt"

type S struct { a int32 }

func main() {
	// Goにはコンストラクタがない。
	// そのためコピーコンストラクタといったものについては考える必要はない。

	// スタック上の値として
	s := S{ 42 }
	fmt.Println(s.a)  // 42

	// 束縛はない。
	// newはヒープアロケーションということではなくallocates storage for a variable
	// https://go.dev/ref/spec#Allocation
	// *Tを返すので型も異なる。
	// s := new(S) <-- 2つの意味で間違い
	s2 := new(S)
	fmt.Println(s2.a) // newの場合は値の初期値以外の方法での初期化はできないようだ。
}
