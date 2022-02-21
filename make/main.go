package main

import "fmt"

func main() {
	// https://go.dev/ref/spec#Making_slices_maps_and_channels

	// スライスの長さ・キャパシティはintの範囲
	// intの定義は32 or 64bitとのこと(アーキテクチャ依存)
	// ただしptrdiff_tと同じかは不明。
	// https://go.dev/ref/spec#Numeric_types
	// capを1<<33にしたところout of memoryになったため、
	// x86_64ではintは64bitなのでptrdiff_t相当でよさそう？
	s := make([]int, 10, 10)

	// どうでもいいがPrintlnの出力にコンマは出ない
	fmt.Println(s)  // [0 0 0 0 0 0 0 0 0 0]

	// mapを作るとき
	m := make(map[string]int)
	m["a"] = 42
	fmt.Println(m)
}
