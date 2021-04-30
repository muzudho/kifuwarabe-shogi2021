package main

import (
	"math/rand"
	"time"

	"github.com/muzudho/kifuwarabe-shogi2021/take11"
)

// main - 最初に実行されます
func main() {
	// fmt.Printf("(11-12) mod 10=%d\n", (11-12)%10)

	// fmt.Println("Hello, world!")
	// take1.MainLoop()
	// take2.MainLoop()
	// take3.MainLoop()
	// take4.MainLoop()
	// take5.MainLoop()

	// ゲーム向けの軽い乱数のタネ
	rand.Seed(time.Now().UnixNano())

	// take6.MainLoop()
	// take7.MainLoop()
	// take8.MainLoop()
	// take9.MainLoop()
	// take10.MainLoop()
	take11.MainLoop()
}
