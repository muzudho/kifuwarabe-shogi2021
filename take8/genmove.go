package take8

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

// GenControl - 利いているマスの一覧を返します。動けるマスではありません。
func GenControl(pPos *Position, from Square) []Square {
	sq_list := []Square{}

	piece := pPos.Board[from]

	// ２つ先のマスから斜めに長い利き
	switch piece {
	case PIECE_B1, PIECE_PB1, PIECE_B2, PIECE_PB2:
		if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
			for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
		if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
			for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
		if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
			for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
		if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
			for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから先手香車の長い利き
	switch piece {
	case PIECE_L1, PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
		if Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
			for to := from - 2; Rank(to) != 0; to -= 1 { // 上
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから後手香車の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_L2, PIECE_R2, PIECE_PR2:
		if Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
			for to := from + 2; Rank(to) != 0; to += 1 { // 下
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
	default:
		// Ignored
	}

	// ２つ横のマスから飛の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
		if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
			for to := from + 20; File(to) != 0; to += 10 { // 左
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
		if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
			for to := from - 20; File(to) != 0; to -= 10 { // 右
				sq_list = append(sq_list, to)
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}
	default:
		// Ignored
	}

	// 先手桂の利き
	if piece == PIECE_N1 {
		if to := from + 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左上桂馬飛び
			sq_list = append(sq_list, to)
		}
		if to := from - 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右上桂馬飛び
			sq_list = append(sq_list, to)
		}
	}

	// 後手桂の利き
	if piece == PIECE_N2 {
		if to := from + 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右下
			sq_list = append(sq_list, to)
		}
	}

	// 先手歩の利き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_L1, PIECE_P1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1, PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from - 1; Rank(to) != 0 { // 上
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 後手歩の利き
	switch piece {
	case PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_L2, PIECE_P2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2, PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1:
		if to := from + 1; Rank(to) != 0 { // 下
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 先手斜め前の利き
	switch piece {
	case PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_PS1, PIECE_PN1, PIECE_PL1,
		PIECE_PP1, PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_S2:
		if to := from + 9; File(to) != 0 && Rank(to) != 0 { // 左上
			sq_list = append(sq_list, to)
		}
		if to := from - 11; File(to) != 0 && Rank(to) != 0 { // 右上
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 後手斜め前の利き
	switch piece {
	case PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_PS2, PIECE_PN2, PIECE_PL2,
		PIECE_PP2, PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_S1:
		if to := from + 11; File(to) != 0 && Rank(to) != 0 { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 9; File(to) != 0 && Rank(to) != 0 { // 右下
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 横１マスの利き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
		PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from + 10; File(to) != 0 { // 左
			sq_list = append(sq_list, to)
		}
		if to := from - 10; File(to) != 0 { // 右
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	return sq_list
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friendKingSq := pPos.KingLocations[pPos.Phase-1]
	opponent := FlipPhase(pPos.Phase)

	if pPos.ControlBoards[opponent-1][friendKingSq] > 0 {
		// 王手されています
		// TODO アタッカーがどの駒か調べたいが……。一手前に動かした駒か、空き王手のどちらかしかなくないか（＾～＾）？
		// 王手されているところが開始局面だと、一手前を調べることができないので、やっぱ調べるしか（＾～＾）
		// 空き王手を利用して、2箇所から 長い利きが飛んでくることはある（＾～＾）

		// 駒を動かしてみて、王手が解除されるか調べるか（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					control_list := GenControl(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, to := range control_list {
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
								pPos.DoMove(move)

								if pPos.ControlBoards[opponent-1][to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPos.UndoMove()
							}
						}
					} else {
						for _, to := range control_list {
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
								pPos.DoMove(move)

								if pPos.ControlBoards[opponent-1][friendKingSq] == 0 {
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPos.UndoMove()
							}
						}
					}
				}
			}
		}
	} else {
		// 王手されていないぜ（＾～＾）
		// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					control_list := GenControl(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, to := range control_list {
							if pPos.Hetero(from, to) && pPos.ControlBoards[opponent-1][to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
							}
						}
					} else {
						for _, to := range control_list {
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
							}
						}
					}
				}
			}
		}
	}

	// TODO 打もやりたい（＾～＾）

	return move_list
}
