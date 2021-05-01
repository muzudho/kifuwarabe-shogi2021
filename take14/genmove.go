package take14

import "fmt"

// 盤上の駒、駒台の駒に対して、37個のルールを実装すればいいはず（＾～＾）
var genmv_k1 = []int{8, 9, 10} // 2
var genmv_r1 = []int{13}
var genmv_b1 = []int{14}
var genmv_g1 = []int{8, 10} // 3
var genmv_s1 = []int{19, 21}
var genmv_n1 = []int{15}
var genmv_l1 = []int{17, 22}
var genmv_p1 = []int{17}
var genmv_pr1 = []int{5, 8, 9, 10}
var genmv_pb1 = []int{6, 8, 9, 10}
var genmv_ps1 = []int{8, 10}   // 3
var genmv_pn1 = []int{8, 10}   // 3
var genmv_pl1 = []int{8, 10}   // 3
var genmv_pp1 = []int{8, 10}   // 3
var genmv_k2 = []int{8, 9, 10} // 2
var genmv_r2 = []int{13}
var genmv_b2 = []int{14}
var genmv_g2 = []int{9, 10} // 4
var genmv_s2 = []int{20, 21}
var genmv_n2 = []int{16}
var genmv_l2 = []int{18, 23}
var genmv_p2 = []int{18}
var genmv_pr2 = []int{5, 8, 9, 10}
var genmv_pb2 = []int{6, 8, 9, 10}
var genmv_ps2 = []int{9, 10} // 4
var genmv_pn2 = []int{9, 10} // 4
var genmv_pl2 = []int{9, 10} // 4
var genmv_pp2 = []int{9, 10} // 4

var genmv_dr1 = []int{31}
var genmv_db1 = []int{31}
var genmv_dg1 = []int{31}
var genmv_ds1 = []int{31}
var genmv_dn1 = []int{32}
var genmv_dl1 = []int{34}
var genmv_dp1 = []int{36}
var genmv_dr2 = []int{31}
var genmv_db2 = []int{31}
var genmv_dg2 = []int{31}
var genmv_ds2 = []int{31}
var genmv_dn2 = []int{33}
var genmv_dl2 = []int{35}
var genmv_dp2 = []int{37}

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

// GenControl - 利いているマスの一覧を返します。動けるマスではありません。
// 成らないと移動できないが、成れば移動できるマスがあるので、移動先と成りの２つセットで返します。
// TODO 成る、成らないも入れたいぜ（＾～＾）
func GenControl(pPos *Position, from Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	var genmv_list []int

	if from == SQUARE_EMPTY {
		panic(fmt.Errorf("GenControl has empty square"))
	} else if OnHands(from) {
		// 打なら

		nifuCheck := false

		switch from {
		case SQ_R1:
			genmv_list = genmv_dr1
		case SQ_B1:
			genmv_list = genmv_db1
		case SQ_G1:
			genmv_list = genmv_dg1
		case SQ_S1:
			genmv_list = genmv_ds1
		case SQ_N1:
			genmv_list = genmv_dn1
		case SQ_L1:
			genmv_list = genmv_dl1
		case SQ_P1:
			genmv_list = genmv_dp1
		case SQ_R2:
			genmv_list = genmv_dr1
		case SQ_B2:
			genmv_list = genmv_db1
		case SQ_G2:
			genmv_list = genmv_dg1
		case SQ_S2:
			genmv_list = genmv_ds1
		case SQ_N2:
			genmv_list = genmv_dn1
		case SQ_L2:
			genmv_list = genmv_dl2
		case SQ_P2:
			genmv_list = genmv_dp2
		default:
			panic(fmt.Errorf("Unknown hand from=%d", from))
		}

		for _, step := range genmv_list {
			switch step {
			case 24:
				makeDrop(pPos, 1, moveEndList, nifuCheck)
			case 25:
				makeDrop(pPos, 2, moveEndList, nifuCheck)
			case 26:
				makeDrop(pPos, 3, moveEndList, nifuCheck)
			case 27:
				makeDrop(pPos, 4, moveEndList, nifuCheck)
				makeDrop(pPos, 5, moveEndList, nifuCheck)
				makeDrop(pPos, 6, moveEndList, nifuCheck)
			case 28:
				makeDrop(pPos, 7, moveEndList, nifuCheck)
			case 29:
				makeDrop(pPos, 8, moveEndList, nifuCheck)
			case 30:
				makeDrop(pPos, 9, moveEndList, nifuCheck)
			case 31:
				genmv_list = append(genmv_list, 24)
				genmv_list = append(genmv_list, 25)
				genmv_list = append(genmv_list, 26)
				genmv_list = append(genmv_list, 27)
				genmv_list = append(genmv_list, 28)
				genmv_list = append(genmv_list, 29)
				genmv_list = append(genmv_list, 30)
			case 32:
				genmv_list = append(genmv_list, 26)
				genmv_list = append(genmv_list, 27)
				genmv_list = append(genmv_list, 28)
				genmv_list = append(genmv_list, 29)
				genmv_list = append(genmv_list, 30)
			case 33:
				genmv_list = append(genmv_list, 24)
				genmv_list = append(genmv_list, 25)
				genmv_list = append(genmv_list, 26)
				genmv_list = append(genmv_list, 27)
				genmv_list = append(genmv_list, 28)
			case 34:
				genmv_list = append(genmv_list, 25)
				genmv_list = append(genmv_list, 26)
				genmv_list = append(genmv_list, 27)
				genmv_list = append(genmv_list, 28)
				genmv_list = append(genmv_list, 29)
				genmv_list = append(genmv_list, 30)
			case 35:
				genmv_list = append(genmv_list, 24)
				genmv_list = append(genmv_list, 25)
				genmv_list = append(genmv_list, 26)
				genmv_list = append(genmv_list, 27)
				genmv_list = append(genmv_list, 28)
				genmv_list = append(genmv_list, 29)
			case 36:
				genmv_list = append(genmv_list, 34)
			case 37:
				genmv_list = append(genmv_list, 35)
			default:
				panic(fmt.Errorf("Unknown step=%d", step))
			}
		}

	} else {
		// 打でないなら
		piece := pPos.Board[from]
		phase := Who(piece)

		switch piece {
		case PIECE_EMPTY:
			panic(fmt.Errorf("Piece empty"))
		case PIECE_K1:
			genmv_list = genmv_k1
		case PIECE_R1:
			genmv_list = genmv_r1
		case PIECE_B1:
			genmv_list = genmv_b1
		case PIECE_G1:
			genmv_list = genmv_g1
		case PIECE_S1:
			genmv_list = genmv_s1
		case PIECE_N1:
			genmv_list = genmv_n1
		case PIECE_L1:
			genmv_list = genmv_l1
		case PIECE_P1:
			genmv_list = genmv_p1
		case PIECE_PR1:
			genmv_list = genmv_pr1
		case PIECE_PB1:
			genmv_list = genmv_pb1
		case PIECE_PS1:
			genmv_list = genmv_ps1
		case PIECE_PN1:
			genmv_list = genmv_pn1
		case PIECE_PL1:
			genmv_list = genmv_pl1
		case PIECE_PP1:
			genmv_list = genmv_pp1
		case PIECE_K2:
			genmv_list = genmv_k2
		case PIECE_R2:
			genmv_list = genmv_r2
		case PIECE_B2:
			genmv_list = genmv_b2
		case PIECE_G2:
			genmv_list = genmv_g2
		case PIECE_S2:
			genmv_list = genmv_s2
		case PIECE_N2:
			genmv_list = genmv_n2
		case PIECE_L2:
			genmv_list = genmv_l2
		case PIECE_P2:
			genmv_list = genmv_p2
		case PIECE_PR2:
			genmv_list = genmv_pr2
		case PIECE_PB2:
			genmv_list = genmv_pb2
		case PIECE_PS2:
			genmv_list = genmv_ps2
		case PIECE_PN2:
			genmv_list = genmv_pn2
		case PIECE_PL2:
			genmv_list = genmv_pl2
		case PIECE_PP2:
			genmv_list = genmv_pp2
		default:
			panic(fmt.Errorf("Unknown piece=%d", piece))
		}

		for _, step := range genmv_list {
			switch step {
			case 1:
				// 成れない (Ignored)
			case 2:
				// genmv_list = append(genmv_list, 1)
			case 3:
				// genmv_list = append(genmv_list, 1)
			case 4:
				// genmv_list = append(genmv_list, 1)
			case 5:
				// ２つ先のマスからの上への長い利き
				makeFrontLong(pPos, from, false, moveEndList)
				// ２つ先のマスからの下への長い利き
				makeBackLong(pPos, from, false, moveEndList)
				// ２つ先のマスからの横への長い利き
				makeSideLong(pPos, from, false, moveEndList)
				// genmv_list = append(genmv_list, 1)
			case 6:
				// ２つ先のマスからの斜めへの長い利き
				makeDiagonalLong(pPos, from, moveEndList)
				// genmv_list = append(genmv_list, 1)
			case 7:
				genmv_list = append(genmv_list, 2)
				genmv_list = append(genmv_list, 5)
				genmv_list = append(genmv_list, 6)
			case 8:
				// 先手から見て斜め前の利き
				makeFrontDiagonal(from, moveEndList)
				genmv_list = append(genmv_list, 3)
				genmv_list = append(genmv_list, 7)
			case 9:
				genmv_list = append(genmv_list, 4)
				genmv_list = append(genmv_list, 7)
			case 10:
				// 先手から見て１つ上への利き
				makeFront(from, moveEndList)

				// 先手から見て１つ後ろへの利き
				makeBack(from, moveEndList)

				// 先手から見て１つ横への利き
				makeSide(from, false, moveEndList)
				genmv_list = append(genmv_list, 3)
				genmv_list = append(genmv_list, 4)
				genmv_list = append(genmv_list, 7)
			case 11:
				// Ignored
			case 12:
				// Ignored
			case 13:
				// 移動元または移動先が敵陣なら成れる
				var promote bool
				switch phase {
				case FIRST:
					promote = File(from) < 4
				case SECOND:
					promote = File(from) > 6
				default:
					panic(fmt.Errorf("Unknown phase=%d", phase))
				}

				// ２つ先のマスからの横への長い利き
				makeSideLong(pPos, from, promote, moveEndList)

				// 先手から見て１つ横への利き
				makeSide(from, promote, moveEndList)

				genmv_list = append(genmv_list, 12)
			case 14:
				// ２つ先のマスからの斜めへの長い利き
				makeDiagonalLong(pPos, from, moveEndList)
				genmv_list = append(genmv_list, 12)
			case 15:
				// 先手桂の利き
				makeFrontKnight(from, moveEndList)
				genmv_list = append(genmv_list, 12)
			case 16:
				// 後手桂の利き
				makeBackKnight(from, moveEndList)
				genmv_list = append(genmv_list, 12)
			case 17:
				// 先手から見て１つ上への利き
				makeFront(from, moveEndList)
				genmv_list = append(genmv_list, 12)
			case 18:
				// 先手から見て１つ後ろへの利き
				makeBack(from, moveEndList)
				genmv_list = append(genmv_list, 12)
			case 19:
				// 先手から見て１つ上への利き
				makeFront(from, moveEndList)
				genmv_list = append(genmv_list, 13)
			case 20:
				// 先手から見て１つ後ろへの利き
				makeBack(from, moveEndList)
				genmv_list = append(genmv_list, 13)
			case 21:
				// 先手から見て斜め前の利き
				makeFrontDiagonal(from, moveEndList)
				// 先手から見て斜め後ろの利き
				makeBackDiagonal(from, moveEndList)
				genmv_list = append(genmv_list, 13)
			case 22:
				// ２つ先のマスからの上への長い利き
				makeFrontLong(pPos, from, true, moveEndList)
				genmv_list = append(genmv_list, 13)
			case 23:
				genmv_list = append(genmv_list, 13)
			default:
				panic(fmt.Errorf("Unknown step=%d", step))
			}
		}
	}

	return moveEndList
}

func makeDrop(pPos *Position, rank Square, moveEndList []MoveEnd, nifuCheck bool) {
	for file := Square(9); file > 0; file-- {
		if nifuCheck && NifuFirst(pPos, file) { // ２歩禁止
			continue
		}
		to := SquareFrom(file, rank)
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
	}
}

// ２つ先のマスからの上への長い利き
func makeFrontLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	if Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
		for to := from - 2; Rank(to) != 0; to -= 1 { // 上
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// ２つ先のマスからの下への長い利き
func makeBackLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	if Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
		for to := from + 2; Rank(to) != 0; to += 1 { // 下
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// ２つ先のマスからの横への長い利き
func makeSideLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	// ２つ先のマスからの左への長い利き
	if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
		for to := from + 20; File(to) != 0; to += 10 { // 左
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

	// ２つ先のマスからの右への長い利き
	if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
		for to := from - 20; File(to) != 0; to -= 10 { // 右
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 先手から見て斜め前の利き
func makeFrontDiagonal(from Square, moveEndList []MoveEnd) {
	if to := from + 9; File(to) != 0 && Rank(to) != 0 { // 左上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
	}
	if to := from - 11; File(to) != 0 && Rank(to) != 0 { // 右上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
	}
}

// 先手から見て斜め後ろの利き
func makeBackDiagonal(from Square, moveEndList []MoveEnd) {
	// 先手から見て斜め後ろの利き
	if to := from + 11; File(to) != 0 && Rank(to) != 0 { // 左下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
	}
	if to := from - 9; File(to) != 0 && Rank(to) != 0 { // 右下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
	}
}

// 先手から見て１つ上への利き
func makeFront(from Square, moveEndList []MoveEnd) {
	// 移動元が２段目のときは必ずならなければならない
	var keepGoing = File(from) != 2
	// 移動元または移動先が敵陣なら成れる
	var promote = File(from) < 5

	if to := from - 1; Rank(to) != 0 { // 上
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 先手から見て１つ横への利き
func makeSide(from Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 10; File(to) != 0 { // 左
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 10; File(to) != 0 { // 右
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 先手から見て１つ後ろへの利き
func makeBack(from Square, moveEndList []MoveEnd) {
	// 移動元が８段目のときは必ずならなければならない
	var keepGoing = File(from) != 8
	// 移動元または移動先が敵陣なら成れる
	var promote = File(from) > 5

	if to := from + 1; Rank(to) != 0 { // 下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 先手桂の利き
func makeFrontKnight(from Square, moveEndList []MoveEnd) {
	// 移動元が３段目のときは必ずならなければならない
	var keepGoing = File(from) != 3
	// 移動元または移動先が敵陣なら成れる
	var promote = File(from) < 6

	if 2 < Rank(from) && Rank(from) < 10 {
		if 0 < File(from) && File(from) < 9 { // 左上桂馬飛び
			to := from + 8
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
		}
		if 1 < File(from) && File(from) < 10 { // 右上桂馬飛び
			to := from - 12
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
		}
	}
}

// 後手桂の利き
func makeBackKnight(from Square, moveEndList []MoveEnd) {
	// 移動元が７段目のときは必ずならなければならない
	var keepGoing = File(from) != 7
	// 移動元または移動先が敵陣なら成れる
	var promote = File(from) > 4

	if to := from + 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// ２つ先のマスからの斜めへの長い利き
func makeDiagonalLong(pPos *Position, from Square, moveEndList []MoveEnd) {
	if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file Square) bool {
	for rank := Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file Square) bool {
	for rank := Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P2 {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPosSys *PositionSystem, pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friend := pPosSys.GetPhase()
	var friendKingSq Square
	var hand_start int
	var hand_end int
	// var opponentKingSq Square
	var pOpponentSumCB *ControlBoard
	if friend == FIRST {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K1)
		hand_start = HAND_IDX_START
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2]
	} else if friend == SECOND {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K2)
		hand_start = HAND_IDX_START + HAND_TYPE_SIZE
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1]
	} else {
		panic(fmt.Errorf("Unknown phase=%d", friend))
	}
	hand_end = hand_start + HAND_TYPE_SIZE

	if !OnBoard(friendKingSq) {
		// 自玉が盤上にない場合は、指し手を返しません

	} else if pOpponentSumCB.Board1[friendKingSq] > 0 {
		// 相手の利きテーブルの自玉のマスに利きがあるか
		// 王手されています
		// fmt.Printf("Debug: Checked friendKingSq=%d opponentKingSq=%d friend=%d opponent=%d\n", friendKingSq, opponentKingSq, friend, opponent)
		// TODO アタッカーがどの駒か調べたいが……。一手前に動かした駒か、空き王手のどちらかしかなくないか（＾～＾）？
		// 王手されているところが開始局面だと、一手前を調べることができないので、やっぱ調べるしか（＾～＾）
		// 空き王手を利用して、2箇所から 長い利きが飛んでくることはある（＾～＾）

		// 盤上の駒を動かしてみて、王手が解除されるか調べるか（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					control_list := GenControl(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board1[to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPosSys.UndoMove(pPos)
							}
						}
					} else {
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board1[friendKingSq] == 0 {
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPosSys.UndoMove(pPos)
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := Square(hand_index) + SQ_HAND_START
				control_list := GenControl(pPos, hand_sq)

				for _, moveEnd := range control_list {
					to := moveEnd.GetDestination()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := NewMoveValue2(hand_sq, to)
						pPosSys.DoMove(pPos, move)

						if pOpponentSumCB.Board1[friendKingSq] == 0 {
							// 王手が解除されてるから採用（＾～＾）
							move_list = append(move_list, move)
						}

						pPosSys.UndoMove(pPos)

					}
				}
			}
		}

	} else {
		// 王手されていないぜ（＾～＾）
		// fmt.Printf("Debug: Not checked\n")

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
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) && pOpponentSumCB.Board1[to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
							}
						}
					} else {
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := Square(hand_index) + SQ_HAND_START
				control_list := GenControl(pPos, hand_sq)

				for _, moveEnd := range control_list {
					to := moveEnd.GetDestination()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move_list = append(move_list, NewMoveValue2(hand_sq, to))
					}
				}
			}
		}
	}

	return move_list
}
