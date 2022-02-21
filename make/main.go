package main

import "fmt"

func main() {
	// https://go.dev/ref/spec#Making_slices_maps_and_channels

	// スライスの長さ・キャパシティはintの範囲
	s := make([]int, 10, 10)

	// どうでもいいがPrintlnの出力にコンマは出ない
	fmt.Println(s)  // [0 0 0 0 0 0 0 0 0 0]

	// mapを作るとき
	m := make(map[string]int)
	m["a"] = 42
	fmt.Println(m)
}
