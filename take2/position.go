package take2

import (
	"fmt"
	"strings"
)

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）時計回りに９０°回転バロック・ダイアゴナルを軸に反転した将棋盤を想像しろだぜ（＾～＾）
	Board []string
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pos *Position) ReadPosition(command string) {
	G.Log.Trace("command=%s\n", command)

	if strings.HasPrefix(command, "position startpos") {
		// 初期局面にします
		pos.Board = []string{
			"", "a", "b", "c", "d", "e", "f", "g", "h", "i",
			"1", "", "", "", "", "", "", "", "", "",
			"2", "", "", "", "", "", "", "", "", "",
			"3", "", "", "", "", "", "", "", "", "",
			"4", "", "", "", "", "", "", "", "", "",
			"5", "", "", "", "", "", "", "", "", "",
			"6", "", "", "", "", "", "", "", "", "",
			"7", "", "", "", "", "", "", "", "", "",
			"8", "", "", "", "", "", "", "", "", "",
			"9", "", "", "", "", "", "", "", "", "",
		}
	}
}

// Print - 局面出力（＾ｑ＾）
func (pos *Position) Sprint() string {
	return "\n" +
		//
		fmt.Sprintf("[? moves / First / ? repeats]\n") +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pos.Board[90], pos.Board[80], pos.Board[70], pos.Board[60], pos.Board[50], pos.Board[40], pos.Board[30], pos.Board[20], pos.Board[10], pos.Board[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[91], pos.Board[81], pos.Board[71], pos.Board[61], pos.Board[51], pos.Board[41], pos.Board[31], pos.Board[21], pos.Board[11], pos.Board[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[92], pos.Board[82], pos.Board[72], pos.Board[62], pos.Board[52], pos.Board[42], pos.Board[32], pos.Board[22], pos.Board[12], pos.Board[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[93], pos.Board[83], pos.Board[73], pos.Board[63], pos.Board[53], pos.Board[43], pos.Board[33], pos.Board[23], pos.Board[13], pos.Board[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[94], pos.Board[84], pos.Board[74], pos.Board[64], pos.Board[54], pos.Board[44], pos.Board[34], pos.Board[24], pos.Board[14], pos.Board[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[95], pos.Board[85], pos.Board[75], pos.Board[65], pos.Board[55], pos.Board[45], pos.Board[35], pos.Board[25], pos.Board[15], pos.Board[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[96], pos.Board[86], pos.Board[76], pos.Board[66], pos.Board[56], pos.Board[46], pos.Board[36], pos.Board[26], pos.Board[16], pos.Board[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[97], pos.Board[87], pos.Board[77], pos.Board[67], pos.Board[57], pos.Board[47], pos.Board[37], pos.Board[27], pos.Board[17], pos.Board[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[98], pos.Board[88], pos.Board[78], pos.Board[68], pos.Board[58], pos.Board[48], pos.Board[38], pos.Board[28], pos.Board[18], pos.Board[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[99], pos.Board[89], pos.Board[79], pos.Board[69], pos.Board[59], pos.Board[49], pos.Board[39], pos.Board[29], pos.Board[19], pos.Board[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
}
