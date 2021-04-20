package take2

import "strings"

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）左右反転した将棋盤を想像しろだぜ（＾～＾）
	Board []string
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pos *Position) ReadPosition(command string) {
	G.Log.Trace("command=%s\n", command)

	if strings.HasPrefix(command, "position startpos") {
		// 初期局面にします
		pos.Board = []string{
			"l", "n", "s", "g", "k", "g", "s", "n", "l", "", // 0段目
			"", "b", "", "", "", "", "", "", "r", "",
			"p", "p", "p", "p", "p", "p", "p", "p", "p", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"P", "P", "P", "P", "P", "P", "P", "P", "P", "",
			"", "R", "", "", "", "", "", "B", "", "",
			"L", "N", "S", "G", "K", "G", "S", "N", "L", "",
			"", "", "", "", "", "", "", "", "", "",
		}
	}
}

// Print - 局面出力（＾ｑ＾）
func (pos *Position) Sprint() string {
	return "TODO position\n"
}
