package take11

import (
	"bytes"
	"fmt"
)

// Print - 局面出力（＾ｑ＾）
func (pPos *Position) Sprint(boardLayer int) string {
	var phase_str string
	switch pPos.GetPhase() {
	case FIRST:
		phase_str = "First"
	case SECOND:
		phase_str = "Second"
	default:
		phase_str = "?"
	}

	// 0段目
	zeroRanks := [10]string{"  9", "  8", "  7", "  6", "  5", "  4", "  3", "  2", "  1", "   "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	if !pPos.IsEmptySq(boardLayer, 90) {
		zeroRanks[0] = pPos.Board[boardLayer][90].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 80) {
		zeroRanks[1] = pPos.Board[boardLayer][80].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 70) {
		zeroRanks[2] = pPos.Board[boardLayer][70].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 60) {
		zeroRanks[3] = pPos.Board[boardLayer][60].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 50) {
		zeroRanks[4] = pPos.Board[boardLayer][50].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 40) {
		zeroRanks[5] = pPos.Board[boardLayer][40].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 30) {
		zeroRanks[6] = pPos.Board[boardLayer][30].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 20) {
		zeroRanks[7] = pPos.Board[boardLayer][20].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 10) {
		zeroRanks[8] = pPos.Board[boardLayer][10].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 0) {
		zeroRanks[9] = pPos.Board[boardLayer][0].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 1) {
		zeroFiles[0] = pPos.Board[boardLayer][1].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 2) {
		zeroFiles[1] = pPos.Board[boardLayer][2].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 3) {
		zeroFiles[2] = pPos.Board[boardLayer][3].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 4) {
		zeroFiles[3] = pPos.Board[boardLayer][4].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 5) {
		zeroFiles[4] = pPos.Board[boardLayer][5].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 6) {
		zeroFiles[5] = pPos.Board[boardLayer][6].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 7) {
		zeroFiles[6] = pPos.Board[boardLayer][7].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 8) {
		zeroFiles[7] = pPos.Board[boardLayer][8].ToCode()
	}
	if !pPos.IsEmptySq(boardLayer, 9) {
		zeroFiles[8] = pPos.Board[boardLayer][9].ToCode()
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPos.StartMovesNum, (pPos.StartMovesNum+pPos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[boardLayer][7], pPos.Hands[boardLayer][8], pPos.Hands[boardLayer][9], pPos.Hands[boardLayer][10], pPos.Hands[boardLayer][11], pPos.Hands[boardLayer][12], pPos.Hands[boardLayer][13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf("%3s%3s%3s%3s%3s%3s%3s%3s%3s%3s\n", zeroRanks[0], zeroRanks[1], zeroRanks[2], zeroRanks[3], zeroRanks[4], zeroRanks[5], zeroRanks[6], zeroRanks[7], zeroRanks[8], zeroRanks[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][91].ToCode(), pPos.Board[boardLayer][81].ToCode(), pPos.Board[boardLayer][71].ToCode(), pPos.Board[boardLayer][61].ToCode(), pPos.Board[boardLayer][51].ToCode(), pPos.Board[boardLayer][41].ToCode(), pPos.Board[boardLayer][31].ToCode(), pPos.Board[boardLayer][21].ToCode(), pPos.Board[boardLayer][11].ToCode(), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][92].ToCode(), pPos.Board[boardLayer][82].ToCode(), pPos.Board[boardLayer][72].ToCode(), pPos.Board[boardLayer][62].ToCode(), pPos.Board[boardLayer][52].ToCode(), pPos.Board[boardLayer][42].ToCode(), pPos.Board[boardLayer][32].ToCode(), pPos.Board[boardLayer][22].ToCode(), pPos.Board[boardLayer][12].ToCode(), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][93].ToCode(), pPos.Board[boardLayer][83].ToCode(), pPos.Board[boardLayer][73].ToCode(), pPos.Board[boardLayer][63].ToCode(), pPos.Board[boardLayer][53].ToCode(), pPos.Board[boardLayer][43].ToCode(), pPos.Board[boardLayer][33].ToCode(), pPos.Board[boardLayer][23].ToCode(), pPos.Board[boardLayer][13].ToCode(), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][94].ToCode(), pPos.Board[boardLayer][84].ToCode(), pPos.Board[boardLayer][74].ToCode(), pPos.Board[boardLayer][64].ToCode(), pPos.Board[boardLayer][54].ToCode(), pPos.Board[boardLayer][44].ToCode(), pPos.Board[boardLayer][34].ToCode(), pPos.Board[boardLayer][24].ToCode(), pPos.Board[boardLayer][14].ToCode(), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][95].ToCode(), pPos.Board[boardLayer][85].ToCode(), pPos.Board[boardLayer][75].ToCode(), pPos.Board[boardLayer][65].ToCode(), pPos.Board[boardLayer][55].ToCode(), pPos.Board[boardLayer][45].ToCode(), pPos.Board[boardLayer][35].ToCode(), pPos.Board[boardLayer][25].ToCode(), pPos.Board[boardLayer][15].ToCode(), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][96].ToCode(), pPos.Board[boardLayer][86].ToCode(), pPos.Board[boardLayer][76].ToCode(), pPos.Board[boardLayer][66].ToCode(), pPos.Board[boardLayer][56].ToCode(), pPos.Board[boardLayer][46].ToCode(), pPos.Board[boardLayer][36].ToCode(), pPos.Board[boardLayer][26].ToCode(), pPos.Board[boardLayer][16].ToCode(), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][97].ToCode(), pPos.Board[boardLayer][87].ToCode(), pPos.Board[boardLayer][77].ToCode(), pPos.Board[boardLayer][67].ToCode(), pPos.Board[boardLayer][57].ToCode(), pPos.Board[boardLayer][47].ToCode(), pPos.Board[boardLayer][37].ToCode(), pPos.Board[boardLayer][27].ToCode(), pPos.Board[boardLayer][17].ToCode(), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][98].ToCode(), pPos.Board[boardLayer][88].ToCode(), pPos.Board[boardLayer][78].ToCode(), pPos.Board[boardLayer][68].ToCode(), pPos.Board[boardLayer][58].ToCode(), pPos.Board[boardLayer][48].ToCode(), pPos.Board[boardLayer][38].ToCode(), pPos.Board[boardLayer][28].ToCode(), pPos.Board[boardLayer][18].ToCode(), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[boardLayer][99].ToCode(), pPos.Board[boardLayer][89].ToCode(), pPos.Board[boardLayer][79].ToCode(), pPos.Board[boardLayer][69].ToCode(), pPos.Board[boardLayer][59].ToCode(), pPos.Board[boardLayer][49].ToCode(), pPos.Board[boardLayer][39].ToCode(), pPos.Board[boardLayer][29].ToCode(), pPos.Board[boardLayer][19].ToCode(), zeroFiles[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[boardLayer][0], pPos.Hands[boardLayer][1], pPos.Hands[boardLayer][2], pPos.Hands[boardLayer][3], pPos.Hands[boardLayer][4], pPos.Hands[boardLayer][5], pPos.Hands[boardLayer][6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_text := pPos.createMovesText()

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}

// CreateMovesList - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func (pPos *Position) createMovesText() string {
	moves_text := make([]byte, 0, MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pPos.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pPos.Moves[i].ToCode()...)
	}
	return string(moves_text)
}

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `layer` - 利き数ボードのレイヤー番号（＾～＾）
func (pPos *Position) SprintControl(phase Phase, layer int) string {
	var board [BOARD_SIZE]int8
	var phase_str string
	var title string

	switch phase {
	case FIRST:
		phase_str = "First"
	case SECOND:
		phase_str = "Second"
	default:
		return "\n"
	}

	var ph = phase - 1
	if 0 <= ph && ph < 2 {
		title = fmt.Sprintf("Control(%d)%s", layer, GetControlLayerName(layer))
		board = pPos.ControlBoards[ph][layer]
	}

	return "\n" +
		//
		fmt.Sprintf("[%s %s]\n", title, phase_str) +
		//
		"\n" +
		//
		"  9  8  7  6  5  4  3  2  1\n" +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| a\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| b\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| c\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| e\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| f\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| g\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| h\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| i\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
}

// SprintLocation - あの駒どこにいんの？を表示
func (pPos *Position) SprintLocation() string {
	king1, king2 := pPos.GetKingLocations()
	return "\n" +
		//
		" K   k      R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			king1, king2,
			pPos.RookLocations[0], pPos.RookLocations[1],
			pPos.BishopLocations[0], pPos.BishopLocations[1],
			pPos.LanceLocations[0], pPos.LanceLocations[1],
			pPos.LanceLocations[2], pPos.LanceLocations[3]) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}

// SprintSfen - SFEN文字列返せよ（＾～＾）
func (pPos *Position) SprintSfen(boardLayer int) string {
	// 9x9=81 + 8slash = 89 文字 なんだが成り駒で増えるし めんどくさ（＾～＾）多めに取っとくか（＾～＾）
	// 成り駒２文字なんで、byte型だとめんどくさ（＾～＾）
	buf := make([]byte, 0, 200)

	spaces := 0
	for rank := Square(1); rank < 10; rank += 1 {
		for file := Square(9); file > 0; file -= 1 {
			piece := pPos.Board[boardLayer][SquareFrom(file, rank)]

			if piece != PIECE_EMPTY {
				if spaces > 0 {
					buf = append(buf, OneDigitNumbers[spaces])
					spaces = 0
				}

				pieceString := piece.ToCode()
				length := len(pieceString)
				switch length {
				case 2:
					buf = append(buf, pieceString[0])
					buf = append(buf, pieceString[1])
				case 1:
					buf = append(buf, pieceString[0])
				default:
					panic(fmt.Errorf("LogicError: length=%d", length))
				}
			} else {
				// Space
				spaces += 1
			}

		}

		if spaces > 0 {
			buf = append(buf, OneDigitNumbers[spaces])
			spaces = 0
		}

		if rank < 9 {
			buf = append(buf, '/')
		}
	}

	// 手番
	var phaseStr string
	switch pPos.GetPhase() {
	case FIRST:
		phaseStr = "b"
	case SECOND:
		phaseStr = "w"
	default:
		panic(fmt.Errorf("LogicalError: Unknows phase=[%d]", pPos.GetPhase()))
	}

	// 持ち駒
	hands := ""
	num := pPos.Hands[boardLayer][0]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("R%d", num)
	}

	num = pPos.Hands[boardLayer][1]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("B%d", num)
	}

	num = pPos.Hands[boardLayer][2]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("G%d", num)
	}

	num = pPos.Hands[boardLayer][3]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("S%d", num)
	}

	num = pPos.Hands[boardLayer][4]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("N%d", num)
	}

	num = pPos.Hands[boardLayer][5]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("L%d", num)
	}

	num = pPos.Hands[boardLayer][6]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("P%d", num)
	}

	num = pPos.Hands[boardLayer][7]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("r%d", num)
	}

	num = pPos.Hands[boardLayer][8]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("b%d", num)
	}

	num = pPos.Hands[boardLayer][9]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("g%d", num)
	}

	num = pPos.Hands[boardLayer][10]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("s%d", num)
	}

	num = pPos.Hands[boardLayer][11]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("n%d", num)
	}

	num = pPos.Hands[boardLayer][12]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("l%d", num)
	}

	num = pPos.Hands[boardLayer][13]
	if num == 1 {
		hands += "p"
	} else if num > 1 {
		hands += fmt.Sprintf("p%d", num)
	}

	if hands == "" {
		hands = "-"
	}

	// 手数
	movesNum := pPos.StartMovesNum + pPos.OffsetMovesIndex

	// 指し手
	moves_text := pPos.createMovesText()

	return fmt.Sprintf("position sfen %s %s %s %d moves%s\n", buf, phaseStr, hands, movesNum, moves_text)
}

// SprintRecord - 棋譜表示（＾～＾）
func (pPos *Position) SprintRecord() string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, MOVES_SIZE*9)
	for i := 0; i < pPos.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pPos.Moves[i].ToCode()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pPos.CapturedList[i].ToCode()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pPos *Position) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	buffer.WriteString("Board:")
	for i := 0; i < BOARD_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Board[i]))
	}
	buffer.WriteString("\n")

	king1, king2 := pPos.GetKingLocations()
	buffer.WriteString(fmt.Sprintf("KingLocations:%d,%d,\n", king1, king2))

	buffer.WriteString("BishopLocations:")
	for i := 0; i < 2; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.BishopLocations[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("LanceLocations:")
	for i := 0; i < 2; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.LanceLocations[i]))
	}
	buffer.WriteString("\n")

	for phase := 0; phase < 2; phase += 1 {
		// 利きボード
		for layer := 0; layer < CONTROL_LAYER_ALL_SIZE; layer += 1 {
			buffer.WriteString(pPos.SprintControl(Phase(phase+1), layer))
		}
	}

	buffer.WriteString("Hands:")
	for i := 0; i < 14; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Hands[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pPos.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pPos.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pPos.OffsetMovesIndex))

	buffer.WriteString("Moves:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Moves[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("CapturedList:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
