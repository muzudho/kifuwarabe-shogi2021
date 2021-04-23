package take6

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
	for rank := 1; rank < 10; rank += 1 {
		for file := 1; file < 10; file += 1 {
			from := uint32(file*10 + rank)
			piece := pPos.Board[from]

			switch piece {
			case "K", "k": // 先手玉, 後手玉
				if to := from + 9; to/10 < 10 && to%10 > 0 { // 左上
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from - 1; to%10 > 0 { // 上
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from - 11; to/10 > 0 && to%10 > 0 { // 右上
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from + 10; to/10 < 10 { // 左
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from - 10; to/10 > 0 { // 右
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from + 11; to/10 < 10 && to%10 < 10 { // 左下
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from + 1; to%10 < 10 { // 下
					move_list = append(move_list, NewMoveValue2(from, to))
				}
				if to := from - 9; to/10 > 0 && to%10 < 10 { // 右下
					move_list = append(move_list, NewMoveValue2(from, to))
				}
			}
		}
	}

	return move_list
}
