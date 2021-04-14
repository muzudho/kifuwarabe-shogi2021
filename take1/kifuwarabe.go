package take1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MainLoop - 開始。
func MainLoop() {
	fmt.Println("Take 1")

	// 何か標準入力しろだぜ☆（＾～＾）
	scanner := bufio.NewScanner(os.Stdin)

MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "quit":
			break MainLoop
		}
	}
}
