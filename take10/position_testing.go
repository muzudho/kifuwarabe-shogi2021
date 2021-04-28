// 利きのテスト
package take10

import (
	"fmt"
	"math"
	"math/rand"
)

// TestControl
func TestControl(pPos *Position) (bool, string) {
	pPos.ClearControlLayer(CONTROL_LAYER_TEST_COPY)
	pPos.ClearControlLayer(CONTROL_LAYER_TEST_ERROR)

	// 利きをコピー
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			pPos.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] = pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
		}
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPos)
	move_total := len(move_list)

	for move_seq, move := range move_list {
		// その手を指してみるぜ（＾～＾）
		pPos.DoMove(move)

		// すぐ戻すぜ（＾～＾）
		pPos.UndoMove()

		// 元に戻っていればOK（＾～＾）
		is_error := checkControl(pPos, move_seq, move_total, move)
		if is_error {
			return is_error, fmt.Sprintf("Error! move_seq=(%d/%d) move=%s", move_seq, move_total, move.ToCode())
		}
	}

	return false, ""
}

// Check - 元に戻っていればOK（＾～＾）
func checkControl(pPos *Position, move_seq int, move_total int, move Move) bool {

	is_error := false

	// 誤差調べ
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			diff := pPos.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] - pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
			pPos.ControlBoards[phase][CONTROL_LAYER_TEST_ERROR][sq] = diff
			if diff != 0 {
				is_error = true
			}
		}
	}

	return is_error
}

// SumAbsControl - 利きテーブルの各マスを絶対値にし、その総和を返します
func SumAbsControl(pPos *Position, layer1 int) [2]int {

	sumList := [2]int{0, 0}

	for phase := 0; phase < 2; phase += 1 {
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				sumList[phase] += int(math.Abs(float64(pPos.ControlBoards[phase][layer1][from])))

			}
		}
	}

	return sumList
}

// ShuffleBoard - 盤上の駒、持ち駒をシャッフルします
func ShuffleBoard(pPos *Position) {

	// 盤と駒台との移動
	// 適当な回数
	for i := 0; i < 100; i += 1 {
		// 盤から駒台の方向
		for rank := Square(0); rank < 10; rank += 1 {
			for file := Square(0); file < 10; file += 1 {
				sq := SquareFrom(file, rank)

				// 10マスに1マスは駒台へ
				change := Square(rand.Intn(10))
				if change == 0 {
					piece := pPos.Board[sq]
					if piece != PIECE_EMPTY {
						phase := Who(piece)
						pieceType := What(piece)

						ok := false
						switch phase {
						case FIRST:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands[HAND_R1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands[HAND_B1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands[HAND_G1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands[HAND_S1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands[HAND_N1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands[HAND_L1-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands[HAND_P1-HAND_ORIGIN] += 1
								ok = true
							default:
								// Ignored
							}
						case SECOND:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands[HAND_R2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands[HAND_B2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands[HAND_G2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands[HAND_S2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands[HAND_N2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands[HAND_L2-HAND_ORIGIN] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands[HAND_P2-HAND_ORIGIN] += 1
								ok = true
							default:
								// Ignored
							}
						default:
							panic(fmt.Errorf("Uknown phase=%d", phase))
						}

						if ok {
							pPos.Board[sq] = PIECE_EMPTY
						}
					}

				}
			}
		}

		// 駒台から盤の方向
		for phase := 0; phase < 2; phase += 1 {
			for j := 0; j < 7; j += 1 {
				num := pPos.Hands[phase*int(HAND_TYPE_SIZE)+j]
				if num > 0 {
					sq := Square(rand.Intn(100))
					if File(sq) != 0 && Rank(sq) != 0 {
						// うまく空マスなら移動成功
						if pPos.IsEmptySq(sq) {
							pPos.Board[sq] = HandPieceMap[phase*int(HAND_TYPE_SIZE)+j]
						}
					}
				}
			}
		}
	}

	// 適当に大きな回数
	for i := 0; i < 81*80; i += 1 {
		square1 := Square(rand.Intn(100))
		square2 := Square(rand.Intn(100))
		if File(square1) != 0 && Rank(square1) != 0 && File(square2) != 0 && Rank(square2) != 0 {
			piece := pPos.Board[square1]

			// 位置スワップ
			pPos.Board[square1] = pPos.Board[square2]
			pPos.Board[square2] = piece

			// 成／不成 変更
			promote := Square(rand.Intn(10))
			if promote == 0 {
				pPos.Board[square1] = Promote(pPos.Board[square1])
			} else if promote == 1 {
				pPos.Board[square1] = Demote(pPos.Board[square1])
			}

			// TODO 駒の先後変更（玉除く）
			piece = pPos.Board[square1]
			switch What(piece) {
			case PIECE_TYPE_K, PIECE_TYPE_EMPTY:
				// Ignored
			default:
				phase := Who(piece)
				pieceType := What(piece)

				change := Square(rand.Intn(10))
				if change == 0 {
					phase = FlipPhase(phase)
				}

				pPos.Board[square1] = PieceFromPhPt(phase, pieceType)
			}

		}
	}

}
