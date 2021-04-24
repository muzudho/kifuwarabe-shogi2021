package take7

func GenControl(pPos *Position, from Square) []Square {
	sq_list := []Square{}

	piece := pPos.Board[from]

	// ２つ先のマスから斜めに長い利き
	switch piece {
	case PIECE_B1, PIECE_PB1, PIECE_B2, PIECE_PB2:
		if from/10%10 != 9 && from%10 != 1 {
			if to := from + 9; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左上
				sq_list = append(sq_list, to)
			}
		}
		if from/10%10 != 1 && from%10 != 1 {
			if to := from - 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右上
				sq_list = append(sq_list, to)
			}
		}
		if from/10%10 != 9 && from%10 != 9 {
			if to := from + 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左下
				sq_list = append(sq_list, to)
			}
		}
		if from/10%10 != 1 && from%10 != 9 {
			if to := from - 9; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右下
				sq_list = append(sq_list, to)
			}
		}
	default: // Ignored
	}

	// ２つ先のマスから先手香車の長い利き
	switch piece {
	case PIECE_L1, PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
		if from%10 != 1 {
			for to := from - 1; to%10 != 0 && pPos.Hetero(to); to -= 1 { // 上
				sq_list = append(sq_list, to)
			}
		}
	default: // Ignored
	}

	// ２つ先のマスから後手香車の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_L2, PIECE_R2, PIECE_PR2:
		if from%10 != 9 {
			for to := from + 1; to%10 != 0 && pPos.Hetero(to); to += 1 { // 下
				sq_list = append(sq_list, to)
			}
		}
	default: // Ignored
	}

	// ２つ横のマスから飛の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
	default: // Ignored
		if from/10%10 != 9 {
			for to := from + 20; to/10%10 != 0 && pPos.Hetero(to); to += 10 { // 左
				sq_list = append(sq_list, to)
			}
		}
		if from/10%10 != 1 {
			for to := from - 20; to/10%10 != 0 && pPos.Hetero(to); to -= 10 { // 右
				sq_list = append(sq_list, to)
			}
		}
	}

	// 先手桂の動き
	if piece == PIECE_N1 {
		if to := from + 8; to/10%10 != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左上桂馬飛び
			sq_list = append(sq_list, to)
		}
		if to := from - 12; to/10%10 != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右上桂馬飛び
			sq_list = append(sq_list, to)
		}
	}

	// 後手桂の動き
	if piece == PIECE_N2 {
		if to := from + 12; to/10%10 != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 8; to/10%10 != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右下
			sq_list = append(sq_list, to)
		}
	}

	// 先手歩の動き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_L1, PIECE_P1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1, PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from - 1; to%10 != 0 && pPos.Hetero(to) { // 上
			sq_list = append(sq_list, to)
		}
	default: // Ignored
	}

	// 後手歩の動き
	switch piece {
	case PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_L2, PIECE_P2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2, PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1:
		if to := from + 1; to%10 != 0 && pPos.Hetero(to) { // 下
			sq_list = append(sq_list, to)
		}
	default: // Ignored
	}

	// 先手斜め前の動き
	switch piece {
	case PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_PS1, PIECE_PN1, PIECE_PL1,
		PIECE_PP1, PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_S2:
		if to := from + 9; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左上
			sq_list = append(sq_list, to)
		}
		if to := from - 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右上
			sq_list = append(sq_list, to)
		}
	default: // Ignored
	}

	// 後手斜め前の動き
	switch piece {
	case PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_PS2, PIECE_PN2, PIECE_PL2,
		PIECE_PP2, PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_S1:
		if to := from + 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 9; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右下
			sq_list = append(sq_list, to)
		}
	default: // Ignored
	}

	// 横１マスの動き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
		PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from + 10; to/10%10 != 0 && pPos.Hetero(to) { // 左
			sq_list = append(sq_list, to)
		}
		if to := from - 10; to/10%10 != 0 && pPos.Hetero(to) { // 右
			sq_list = append(sq_list, to)
		}
	default: // Ignored
	}

	return sq_list
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
	for rank := 1; rank < 10; rank += 1 {
		for file := 1; file < 10; file += 1 {
			from := Square(file*10 + rank)
			if pPos.Homo(from) {
				sq_list := GenControl(pPos, from)

				for _, to := range sq_list {
					move_list = append(move_list, NewMoveValue2(from, to))
				}
			}
		}
	}

	return move_list
}
