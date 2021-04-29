package take11

import (
	"bytes"
	"fmt"
)

// Print - 局面出力（＾ｑ＾）
func (pPosSys *PositionSystem) Sprint(b PosLayerT) string {
	var phase_str string
	switch pPosSys.GetPhase() {
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
	if !pPosSys.PPosition[b].IsEmptySq(90) {
		zeroRanks[0] = pPosSys.PPosition[b].Board[90].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(80) {
		zeroRanks[1] = pPosSys.PPosition[b].Board[80].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(70) {
		zeroRanks[2] = pPosSys.PPosition[b].Board[70].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(60) {
		zeroRanks[3] = pPosSys.PPosition[b].Board[60].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(50) {
		zeroRanks[4] = pPosSys.PPosition[b].Board[50].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(40) {
		zeroRanks[5] = pPosSys.PPosition[b].Board[40].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(30) {
		zeroRanks[6] = pPosSys.PPosition[b].Board[30].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(20) {
		zeroRanks[7] = pPosSys.PPosition[b].Board[20].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(10) {
		zeroRanks[8] = pPosSys.PPosition[b].Board[10].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(0) {
		zeroRanks[9] = pPosSys.PPosition[b].Board[0].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(1) {
		zeroFiles[0] = pPosSys.PPosition[b].Board[1].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(2) {
		zeroFiles[1] = pPosSys.PPosition[b].Board[2].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(3) {
		zeroFiles[2] = pPosSys.PPosition[b].Board[3].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(4) {
		zeroFiles[3] = pPosSys.PPosition[b].Board[4].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(5) {
		zeroFiles[4] = pPosSys.PPosition[b].Board[5].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(6) {
		zeroFiles[5] = pPosSys.PPosition[b].Board[6].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(7) {
		zeroFiles[6] = pPosSys.PPosition[b].Board[7].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(8) {
		zeroFiles[7] = pPosSys.PPosition[b].Board[8].ToCode()
	}
	if !pPosSys.PPosition[b].IsEmptySq(9) {
		zeroFiles[8] = pPosSys.PPosition[b].Board[9].ToCode()
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPosSys.StartMovesNum, (pPosSys.StartMovesNum+pPosSys.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPosSys.PPosition[b].Hands[7], pPosSys.PPosition[b].Hands[8], pPosSys.PPosition[b].Hands[9], pPosSys.PPosition[b].Hands[10], pPosSys.PPosition[b].Hands[11], pPosSys.PPosition[b].Hands[12], pPosSys.PPosition[b].Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf("%3s%3s%3s%3s%3s%3s%3s%3s%3s%3s\n", zeroRanks[0], zeroRanks[1], zeroRanks[2], zeroRanks[3], zeroRanks[4], zeroRanks[5], zeroRanks[6], zeroRanks[7], zeroRanks[8], zeroRanks[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[91].ToCode(), pPosSys.PPosition[b].Board[81].ToCode(), pPosSys.PPosition[b].Board[71].ToCode(), pPosSys.PPosition[b].Board[61].ToCode(), pPosSys.PPosition[b].Board[51].ToCode(), pPosSys.PPosition[b].Board[41].ToCode(), pPosSys.PPosition[b].Board[31].ToCode(), pPosSys.PPosition[b].Board[21].ToCode(), pPosSys.PPosition[b].Board[11].ToCode(), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[92].ToCode(), pPosSys.PPosition[b].Board[82].ToCode(), pPosSys.PPosition[b].Board[72].ToCode(), pPosSys.PPosition[b].Board[62].ToCode(), pPosSys.PPosition[b].Board[52].ToCode(), pPosSys.PPosition[b].Board[42].ToCode(), pPosSys.PPosition[b].Board[32].ToCode(), pPosSys.PPosition[b].Board[22].ToCode(), pPosSys.PPosition[b].Board[12].ToCode(), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[93].ToCode(), pPosSys.PPosition[b].Board[83].ToCode(), pPosSys.PPosition[b].Board[73].ToCode(), pPosSys.PPosition[b].Board[63].ToCode(), pPosSys.PPosition[b].Board[53].ToCode(), pPosSys.PPosition[b].Board[43].ToCode(), pPosSys.PPosition[b].Board[33].ToCode(), pPosSys.PPosition[b].Board[23].ToCode(), pPosSys.PPosition[b].Board[13].ToCode(), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[94].ToCode(), pPosSys.PPosition[b].Board[84].ToCode(), pPosSys.PPosition[b].Board[74].ToCode(), pPosSys.PPosition[b].Board[64].ToCode(), pPosSys.PPosition[b].Board[54].ToCode(), pPosSys.PPosition[b].Board[44].ToCode(), pPosSys.PPosition[b].Board[34].ToCode(), pPosSys.PPosition[b].Board[24].ToCode(), pPosSys.PPosition[b].Board[14].ToCode(), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[95].ToCode(), pPosSys.PPosition[b].Board[85].ToCode(), pPosSys.PPosition[b].Board[75].ToCode(), pPosSys.PPosition[b].Board[65].ToCode(), pPosSys.PPosition[b].Board[55].ToCode(), pPosSys.PPosition[b].Board[45].ToCode(), pPosSys.PPosition[b].Board[35].ToCode(), pPosSys.PPosition[b].Board[25].ToCode(), pPosSys.PPosition[b].Board[15].ToCode(), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[96].ToCode(), pPosSys.PPosition[b].Board[86].ToCode(), pPosSys.PPosition[b].Board[76].ToCode(), pPosSys.PPosition[b].Board[66].ToCode(), pPosSys.PPosition[b].Board[56].ToCode(), pPosSys.PPosition[b].Board[46].ToCode(), pPosSys.PPosition[b].Board[36].ToCode(), pPosSys.PPosition[b].Board[26].ToCode(), pPosSys.PPosition[b].Board[16].ToCode(), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[97].ToCode(), pPosSys.PPosition[b].Board[87].ToCode(), pPosSys.PPosition[b].Board[77].ToCode(), pPosSys.PPosition[b].Board[67].ToCode(), pPosSys.PPosition[b].Board[57].ToCode(), pPosSys.PPosition[b].Board[47].ToCode(), pPosSys.PPosition[b].Board[37].ToCode(), pPosSys.PPosition[b].Board[27].ToCode(), pPosSys.PPosition[b].Board[17].ToCode(), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[98].ToCode(), pPosSys.PPosition[b].Board[88].ToCode(), pPosSys.PPosition[b].Board[78].ToCode(), pPosSys.PPosition[b].Board[68].ToCode(), pPosSys.PPosition[b].Board[58].ToCode(), pPosSys.PPosition[b].Board[48].ToCode(), pPosSys.PPosition[b].Board[38].ToCode(), pPosSys.PPosition[b].Board[28].ToCode(), pPosSys.PPosition[b].Board[18].ToCode(), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPosSys.PPosition[b].Board[99].ToCode(), pPosSys.PPosition[b].Board[89].ToCode(), pPosSys.PPosition[b].Board[79].ToCode(), pPosSys.PPosition[b].Board[69].ToCode(), pPosSys.PPosition[b].Board[59].ToCode(), pPosSys.PPosition[b].Board[49].ToCode(), pPosSys.PPosition[b].Board[39].ToCode(), pPosSys.PPosition[b].Board[29].ToCode(), pPosSys.PPosition[b].Board[19].ToCode(), zeroFiles[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPosSys.PPosition[b].Hands[0], pPosSys.PPosition[b].Hands[1], pPosSys.PPosition[b].Hands[2], pPosSys.PPosition[b].Hands[3], pPosSys.PPosition[b].Hands[4], pPosSys.PPosition[b].Hands[5], pPosSys.PPosition[b].Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_text := pPosSys.createMovesText()

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}

// Print - ２局面の比較用画面出力（＾ｑ＾）
func (pPosSys *PositionSystem) SprintDiff(b1 PosLayerT, b2 PosLayerT) string {
	var phase_str string
	switch pPosSys.GetPhase() {
	case FIRST:
		phase_str = "First"
	case SECOND:
		phase_str = "Second"
	default:
		phase_str = "?"
	}

	// 0段目
	zeroRanks := [10]string{"    9", "    8", "    7", "    6", "    5", "    4", "    3", "    2", "    1", "     "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	for file := 9; file > -1; file -= 1 {
		if !pPosSys.PPosition[b1].IsEmptySq(Square(file*10)) || !pPosSys.PPosition[b2].IsEmptySq(Square(file*10)) {
			zeroRanks[10-file] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[file*10].ToCode(), pPosSys.PPosition[b2].Board[file*10].ToCode())
		}
	}

	// 0筋目
	for rank := Square(1); rank < 10; rank += 1 {
		if !pPosSys.PPosition[b1].IsEmptySq(rank) || !pPosSys.PPosition[b2].IsEmptySq(rank) {
			zeroFiles[rank-1] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[rank].ToCode(), pPosSys.PPosition[b2].Board[rank].ToCode())
		}
	}

	lines := []string{}
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPosSys.StartMovesNum, (pPosSys.StartMovesNum+pPosSys.OffsetMovesIndex), phase_str))
	lines = append(lines, "\n")
	lines = append(lines, "    r    b    g    s    n    l    p\n")
	lines = append(lines, "+----+----+----+----+----+----+----+\n")

	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buf bytes.Buffer
	for i := 0; i < 7; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands[i+7], pPosSys.PPosition[b2].Hands[i+7]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+\n")

	buf.Reset()
	for i := 0; i < 10; i += 1 {
		buf.WriteString(zeroRanks[i])
	}
	buf.WriteString("\n")
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank := 1
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 2
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 3
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 4
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 5
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 6
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 7
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 8
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 9
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[0]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "          R    B    G    S    N    L    P\n")
	lines = append(lines, "      +----+----+----+----+----+----+----+\n")

	buf.Reset()
	buf.WriteString("      ")
	for i := 0; i < 7; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands[i], pPosSys.PPosition[b2].Hands[i]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, "      +----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "moves")

	lines = append(lines, pPosSys.createMovesText())
	lines = append(lines, "\n")

	buf.Reset()
	for _, line := range lines {
		buf.WriteString(line)
	}
	return buf.String()
}

// CreateMovesList - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func (pPosSys *PositionSystem) createMovesText() string {
	moves_text := make([]byte, 0, MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pPosSys.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pPosSys.Moves[i].ToCode()...)
	}
	return string(moves_text)
}

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `c` - 利き数ボードのレイヤー番号（＾～＾）
func (pPosSys *PositionSystem) SprintControl(phase Phase, c ControlLayerT) string {
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
		title = fmt.Sprintf("Control(%d)%s", c, GetControlLayerName(c))
		board = pPosSys.ControlBoards[ph][c]
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
func (pPosSys *PositionSystem) SprintLocation(b PosLayerT) string {
	return "\n" +
		//
		" K   k      R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			pPosSys.PPosition[b].PieceLocations[PCLOC_K1], pPosSys.PPosition[b].PieceLocations[PCLOC_K2],
			pPosSys.PPosition[b].PieceLocations[PCLOC_R1], pPosSys.PPosition[b].PieceLocations[PCLOC_R2],
			pPosSys.PPosition[b].PieceLocations[PCLOC_B1], pPosSys.PPosition[b].PieceLocations[PCLOC_B2],
			pPosSys.PPosition[b].PieceLocations[PCLOC_L1], pPosSys.PPosition[b].PieceLocations[PCLOC_L2],
			pPosSys.PPosition[b].PieceLocations[PCLOC_L3], pPosSys.PPosition[b].PieceLocations[PCLOC_L4]) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}

// SprintSfen - SFEN文字列返せよ（＾～＾）
func (pPosSys *PositionSystem) SprintSfen(b PosLayerT) string {
	// 9x9=81 + 8slash = 89 文字 なんだが成り駒で増えるし めんどくさ（＾～＾）多めに取っとくか（＾～＾）
	// 成り駒２文字なんで、byte型だとめんどくさ（＾～＾）
	buf := make([]byte, 0, 200)

	spaces := 0
	for rank := Square(1); rank < 10; rank += 1 {
		for file := Square(9); file > 0; file -= 1 {
			piece := pPosSys.PPosition[b].Board[SquareFrom(file, rank)]

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
	switch pPosSys.GetPhase() {
	case FIRST:
		phaseStr = "b"
	case SECOND:
		phaseStr = "w"
	default:
		panic(fmt.Errorf("LogicalError: Unknows phase=[%d]", pPosSys.GetPhase()))
	}

	// 持ち駒
	hands := ""
	num := pPosSys.PPosition[b].Hands[0]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("R%d", num)
	}

	num = pPosSys.PPosition[b].Hands[1]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("B%d", num)
	}

	num = pPosSys.PPosition[b].Hands[2]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("G%d", num)
	}

	num = pPosSys.PPosition[b].Hands[3]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("S%d", num)
	}

	num = pPosSys.PPosition[b].Hands[4]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("N%d", num)
	}

	num = pPosSys.PPosition[b].Hands[5]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("L%d", num)
	}

	num = pPosSys.PPosition[b].Hands[6]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("P%d", num)
	}

	num = pPosSys.PPosition[b].Hands[7]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("r%d", num)
	}

	num = pPosSys.PPosition[b].Hands[8]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("b%d", num)
	}

	num = pPosSys.PPosition[b].Hands[9]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("g%d", num)
	}

	num = pPosSys.PPosition[b].Hands[10]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("s%d", num)
	}

	num = pPosSys.PPosition[b].Hands[11]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("n%d", num)
	}

	num = pPosSys.PPosition[b].Hands[12]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("l%d", num)
	}

	num = pPosSys.PPosition[b].Hands[13]
	if num == 1 {
		hands += "p"
	} else if num > 1 {
		hands += fmt.Sprintf("p%d", num)
	}

	if hands == "" {
		hands = "-"
	}

	// 手数
	movesNum := pPosSys.StartMovesNum + pPosSys.OffsetMovesIndex

	// 指し手
	moves_text := pPosSys.createMovesText()

	return fmt.Sprintf("position sfen %s %s %s %d moves%s\n", buf, phaseStr, hands, movesNum, moves_text)
}

// SprintRecord - 棋譜表示（＾～＾）
func (pPosSys *PositionSystem) SprintRecord() string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, MOVES_SIZE*9)
	for i := 0; i < pPosSys.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pPosSys.Moves[i].ToCode()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pPosSys.CapturedList[i].ToCode()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pPosSys *PositionSystem) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	for b := PosLayerT(0); b < 2; b += 1 {
		buffer.WriteString(fmt.Sprintf("Board[%d]:", b))
		for i := 0; i < BOARD_SIZE; i += 1 {
			buffer.WriteString(fmt.Sprintf("%d,", pPosSys.PPosition[i].Board))
		}
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("KingLocations[%d]:%d,%d\n", b, pPosSys.PPosition[b].PieceLocations[PCLOC_K1], pPosSys.PPosition[b].PieceLocations[PCLOC_K2]))
		buffer.WriteString(fmt.Sprintf("RookLocations[%d]:%d,%d\n", b, pPosSys.PPosition[b].PieceLocations[PCLOC_R1], pPosSys.PPosition[b].PieceLocations[PCLOC_R2]))
		buffer.WriteString(fmt.Sprintf("BishopLocations[%d]:%d,%d\n", b, pPosSys.PPosition[b].PieceLocations[PCLOC_B1], pPosSys.PPosition[b].PieceLocations[PCLOC_B2]))
		buffer.WriteString(fmt.Sprintf("LanceLocations[%d]:%d,%d,%d,%d\n", b, pPosSys.PPosition[b].PieceLocations[PCLOC_L1], pPosSys.PPosition[b].PieceLocations[PCLOC_L2], pPosSys.PPosition[b].PieceLocations[PCLOC_L3], pPosSys.PPosition[b].PieceLocations[PCLOC_L4]))
	}

	for phase := 0; phase < 2; phase += 1 {
		// 利きボード
		for c := ControlLayerT(0); c < CONTROL_LAYER_ALL_SIZE; c += 1 {
			buffer.WriteString(pPosSys.SprintControl(Phase(phase+1), c))
		}
	}

	buffer.WriteString("Hands:")
	for i := 0; i < 14; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPosSys.PPosition[i].Hands))
	}
	buffer.WriteString("\n")

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pPosSys.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pPosSys.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pPosSys.OffsetMovesIndex))

	buffer.WriteString("Moves:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPosSys.Moves[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("CapturedList:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPosSys.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
