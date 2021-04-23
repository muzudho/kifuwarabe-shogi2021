package main

import (
	"math/rand"
	"time"

	"github.com/muzudho/kifuwarabe-shogi2021/take6"
)

// main - 最初に実行されます
func main() {
	// ゲーム向けの軽い乱数のタネ
	rand.Seed(time.Now().UnixNano())

	// fmt.Println("Hello, world!")
	// take1.MainLoop()
	// take2.MainLoop()
	// take3.MainLoop()
	// take4.MainLoop()
	// take5.MainLoop()
	take6.MainLoop()
}
