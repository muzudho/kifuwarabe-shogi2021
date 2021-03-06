// 利きのテスト
package take12

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// TestControl
func TestControl(pPosSys *PositionSystem, pPos *Position) (bool, string) {
	pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY1].Clear()
	pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY2].Clear()

	pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_ERROR1].Clear()
	pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_ERROR2].Clear()

	// 利きをコピー
	copyCb1 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY1]
	sumCb1 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM1]
	copyCb2 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY2]
	sumCb2 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM2]
	for sq := 0; sq < BOARD_SIZE; sq += 1 {
		copyCb1.Board[sq] = sumCb1.Board[sq]
		copyCb2.Board[sq] = sumCb2.Board[sq]
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPosSys, pPos)
	move_total := len(move_list)

	for move_seq, move := range move_list {
		// その手を指してみるぜ（＾～＾）
		pPosSys.DoMove(pPos, move)

		// すぐ戻すぜ（＾～＾）
		pPosSys.UndoMove(pPos)

		// 元に戻っていればOK（＾～＾）
		is_error := checkControl(pPosSys, move_seq, move_total, move)
		if is_error {
			return is_error, fmt.Sprintf("Error! move_seq=(%d/%d) move=%s", move_seq, move_total, move.ToCode())
		}
	}

	return false, ""
}

// Check - 元に戻っていればOK（＾～＾）
func checkControl(pPosSys *PositionSystem, move_seq int, move_total int, move Move) bool {

	is_error := false

	// 誤差調べ
	copyCB1 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY1]
	sumCB1 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM1]
	errorCB1 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_ERROR1]
	copyCB2 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_COPY2]
	sumCB2 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM2]
	errorCB2 := pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_TEST_ERROR2]
	for sq := 0; sq < BOARD_SIZE; sq += 1 {
		diff1 := copyCB1.Board[sq] - sumCB1.Board[sq]
		errorCB1.Board[sq] = diff1
		if diff1 != 0 {
			is_error = true
			break
		}

		diff2 := copyCB2.Board[sq] - sumCB2.Board[sq]
		errorCB2.Board[sq] = diff2
		if diff2 != 0 {
			is_error = true
			break
		}
	}

	return is_error
}

// SumAbsControl - 利きテーブルの各マスを絶対値にし、その総和を返します
func SumAbsControl(pPosSys *PositionSystem, ph1_c ControlLayerT, ph2_c ControlLayerT) [2]int {

	sumList := [2]int{0, 0}

	cb1 := pPosSys.PControlBoardSystem.Boards[ph1_c]
	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 {

			sumList[FIRST-1] += int(math.Abs(float64(cb1.Board[from])))

		}
	}

	cb2 := pPosSys.PControlBoardSystem.Boards[ph2_c]
	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 {

			sumList[SECOND-1] += int(math.Abs(float64(cb2.Board[from])))

		}
	}

	return sumList
}

// ShuffleBoard - 盤上の駒、持ち駒をシャッフルします
// ゲーム中にはできない動きをするので、利きの計算は無視します。
// 最後に利きは再計算します
func ShuffleBoard(pPosSys *PositionSystem, pPos *Position) {

	// 駒の数を数えます
	countList1 := CountAllPieces(pPos)

	// 盤と駒台との移動
	// 適当な回数
	for i := 0; i < 200; i += 1 {

		// 盤から駒台の方向
		for rank := Square(1); rank < 10; rank += 1 {
			for file := Square(9); file > 0; file -= 1 {
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
							case PIECE_TYPE_K:
								pPos.Hands1[HAND_K1_IDX] += 1
								ok = true
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands1[HAND_R1_IDX] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands1[HAND_B1_IDX] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands1[HAND_G1_IDX] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands1[HAND_S1_IDX] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands1[HAND_N1_IDX] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands1[HAND_L1_IDX] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands1[HAND_P1_IDX] += 1
								ok = true
							default:
								// Ignored
							}
						case SECOND:
							switch pieceType {
							case PIECE_TYPE_K:
								pPos.Hands1[HAND_K2_IDX] += 1
								ok = true
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands1[HAND_R2_IDX] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands1[HAND_B2_IDX] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands1[HAND_G2_IDX] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands1[HAND_S2_IDX] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands1[HAND_N2_IDX] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands1[HAND_L2_IDX] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands1[HAND_P2_IDX] += 1
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

		// 駒の数を数えます
		countList2 := CountAllPieces(pPos)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("Shuffle: (1) countError=%d", countError))
		}

		// 駒台から盤の方向
		for hand_index := HAND_IDX_START; hand_index < HAND_IDX_END; hand_index += 1 {
			num := pPos.Hands1[hand_index]
			if num > 0 {
				sq := Square(rand.Intn(100))
				// うまく空マスなら移動成功
				if OnBoard(sq) && pPos.IsEmptySq(sq) {
					pPos.Board[sq] = HandPieceMap1[hand_index]
					pPos.Hands1[hand_index] -= 1
				}
			}
		}

		// 駒の数を数えます
		countList2 = CountAllPieces(pPos)
		countError = CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("Shuffle: (2) countError=%d", countError))
		}
	}

	// 盤上での移動
	// 適当に大きな回数
	for i := 0; i < 81*80; i += 1 {
		sq1 := Square(rand.Intn(100))
		sq2 := Square(rand.Intn(100))
		if OnBoard(sq1) && OnBoard(sq2) && !pPos.IsEmptySq(sq1) {
			piece := pPos.Board[sq1]
			// 位置スワップ
			pPos.Board[sq1] = pPos.Board[sq2]
			pPos.Board[sq2] = piece

			// 成／不成 変更
			promote := Square(rand.Intn(10))
			if promote == 0 {
				pPos.Board[sq2] = Promote(pPos.Board[sq2])
			} else if promote == 1 {
				pPos.Board[sq2] = Demote(pPos.Board[sq2])
			}

			// 駒の先後変更（玉除く）
			piece = pPos.Board[sq2]
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

				pPos.Board[sq2] = PieceFromPhPt(phase, pieceType)
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPos)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("Shuffle: (3) countError=%d", countError))
		}
	}

	// 手番のシャッフル
	switch rand.Intn(2) {
	case 0:
		pPosSys.phase = FIRST
	default:
		pPosSys.phase = SECOND
	}

	// 手目は 1 に戻します
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	G.Chat.Debug(pPos.Sprint(
		pPosSys.phase,
		pPosSys.StartMovesNum,
		pPosSys.OffsetMovesIndex,
		pPosSys.createMovesText()))

	if false {
		var countList [8]int

		if true {
			countList = [8]int{}

			// 盤上
			for rank := Square(1); rank < 10; rank += 1 {
				for file := Square(9); file > 0; file -= 1 {
					sq := SquareFrom(file, rank)

					fmt.Printf("%s,", pPos.Board[sq].ToCode())

					piece := What(pPos.Board[sq])
					switch piece {
					case PIECE_TYPE_K:
						countList[0] += 1
					case PIECE_TYPE_R, PIECE_TYPE_PR:
						countList[1] += 1
					case PIECE_TYPE_B, PIECE_TYPE_PB:
						countList[2] += 1
					case PIECE_TYPE_G:
						countList[3] += 1
					case PIECE_TYPE_S, PIECE_TYPE_PS:
						countList[4] += 1
					case PIECE_TYPE_N, PIECE_TYPE_PN:
						countList[5] += 1
					case PIECE_TYPE_L, PIECE_TYPE_PL:
						countList[6] += 1
					case PIECE_TYPE_P, PIECE_TYPE_PP:
						countList[7] += 1
					default:
						// Ignore
					}
				}
				fmt.Printf("\n")
			}

			// 駒台
			countList[0] += pPos.Hands1[HAND_K1_IDX] + pPos.Hands1[HAND_K2_IDX]
			countList[1] += pPos.Hands1[HAND_R1_IDX] + pPos.Hands1[HAND_R2_IDX]
			countList[2] += pPos.Hands1[HAND_B1_IDX] + pPos.Hands1[HAND_B2_IDX]
			countList[3] += pPos.Hands1[HAND_G1_IDX] + pPos.Hands1[HAND_G2_IDX]
			countList[4] += pPos.Hands1[HAND_S1_IDX] + pPos.Hands1[HAND_S2_IDX]
			countList[5] += pPos.Hands1[HAND_N1_IDX] + pPos.Hands1[HAND_N2_IDX]
			countList[6] += pPos.Hands1[HAND_L1_IDX] + pPos.Hands1[HAND_L2_IDX]
			countList[7] += pPos.Hands1[HAND_P1_IDX] + pPos.Hands1[HAND_P2_IDX]
		} else {
			countList = CountAllPieces(pPos)
		}

		G.Chat.Debug("#Count\n")
		G.Chat.Debug("#-----\n")
		G.Chat.Debug("#King  :%3d\n", countList[0])
		G.Chat.Debug("#Rook  :%3d\n", countList[1])
		G.Chat.Debug("#Bishop:%3d\n", countList[2])
		G.Chat.Debug("#Gold  :%3d\n", countList[3])
		G.Chat.Debug("#Silver:%3d\n", countList[4])
		G.Chat.Debug("#Knight:%3d\n", countList[5])
		G.Chat.Debug("#Lance :%3d\n", countList[6])
		G.Chat.Debug("#Pawn  :%3d\n", countList[7])
		G.Chat.Debug("#----------\n")
		G.Chat.Debug("#Total :%3d\n", countList[0]+countList[1]+countList[2]+countList[3]+countList[4]+countList[5]+countList[6]+countList[7])
	} else {
		ShowAllPiecesCount(pPos)
	}

	// position sfen 文字列を取得
	command := pPosSys.SprintSfenResignation(pPos)
	G.Chat.Debug("#command=%s", command)

	// 利きの再計算もやってくれる
	pPosSys.ReadPosition(pPos, command)

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	G.Chat.Debug(pPos.Sprint(
		pPosSys.phase,
		pPosSys.StartMovesNum,
		pPosSys.OffsetMovesIndex,
		pPosSys.createMovesText()))
	ShowAllPiecesCount(pPos)
	command2 := pPosSys.SprintSfenResignation(pPos)
	G.Chat.Debug("#command2=%s", command2)

	// 駒の数を数えます
	countList2 := CountAllPieces(pPos)
	countError := CountErrorCountLists(countList1, countList2)
	if countError != 0 {
		panic(fmt.Errorf("Shuffle: (4) countError=%d", countError))
	}
}

// CountAllPieces - 駒の数を確認するぜ（＾～＾）
func CountAllPieces(pPos *Position) [8]int {

	countList := [8]int{}

	// 盤上
	for rank := Square(1); rank < 10; rank += 1 {
		for file := Square(9); file > 0; file -= 1 {
			sq := SquareFrom(file, rank)

			piece := What(pPos.Board[sq])
			switch piece {
			case PIECE_TYPE_K:
				countList[0] += 1
			case PIECE_TYPE_R, PIECE_TYPE_PR:
				countList[1] += 1
			case PIECE_TYPE_B, PIECE_TYPE_PB:
				countList[2] += 1
			case PIECE_TYPE_G:
				countList[3] += 1
			case PIECE_TYPE_S, PIECE_TYPE_PS:
				countList[4] += 1
			case PIECE_TYPE_N, PIECE_TYPE_PN:
				countList[5] += 1
			case PIECE_TYPE_L, PIECE_TYPE_PL:
				countList[6] += 1
			case PIECE_TYPE_P, PIECE_TYPE_PP:
				countList[7] += 1
			default:
				// Ignore
			}
		}
	}

	// 駒台
	countList[0] += pPos.Hands1[HAND_K1_IDX] + pPos.Hands1[HAND_K2_IDX]
	countList[1] += pPos.Hands1[HAND_R1_IDX] + pPos.Hands1[HAND_R2_IDX]
	countList[2] += pPos.Hands1[HAND_B1_IDX] + pPos.Hands1[HAND_B2_IDX]
	countList[3] += pPos.Hands1[HAND_G1_IDX] + pPos.Hands1[HAND_G2_IDX]
	countList[4] += pPos.Hands1[HAND_S1_IDX] + pPos.Hands1[HAND_S2_IDX]
	countList[5] += pPos.Hands1[HAND_N1_IDX] + pPos.Hands1[HAND_N2_IDX]
	countList[6] += pPos.Hands1[HAND_L1_IDX] + pPos.Hands1[HAND_L2_IDX]
	countList[7] += pPos.Hands1[HAND_P1_IDX] + pPos.Hands1[HAND_P2_IDX]

	return countList
}

// CountErrorCountLists - 数えた駒の枚数を比較します
func CountErrorCountLists(countList1 [8]int, countList2 [8]int) int {
	sum := 0
	for i := 0; i < 8; i += 1 {
		sum += int(math.Abs(float64(countList1[i] - countList2[i])))
	}
	return sum
}

// copyBoard - 盤[b0] を 盤[b1] にコピーします
func copyBoard(pPos0 *Position, pPos1 *Position) {
	for sq := 0; sq < 100; sq += 1 {
		pPos1.Board[sq] = pPos0.Board[sq]
	}

	pPos1.Hands1 = pPos0.Hands1
	for i := PCLOC_START; i < PCLOC_END; i += 1 {
		pPos1.PieceLocations[i] = pPos0.PieceLocations[i]
	}
}

// copyBoard - 盤[0] を 盤[1] で異なるマスを 盤[2] 盤[3] にセットします
func diffBoard(pPos0 *Position, pPos1 *Position, pPos2 *Position, pPos3 *Position) {
	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPos1.Board[sq] == pPos0.Board[sq] {
			// 等しければ空マス
			pPos2.Board[sq] = PIECE_EMPTY
			pPos3.Board[sq] = PIECE_EMPTY

		} else {
			// 異なったら
			pPos2.Board[sq] = pPos0.Board[sq]
			pPos3.Board[sq] = pPos1.Board[sq]
		}
	}

	// 駒台
	for i := HAND_IDX_START; i < HAND_IDX_END; i += 1 {
		if pPos0.Hands1[i] == pPos1.Hands1[i] {
			// 等しければゼロ
			pPos2.Hands1[i] = 0
			pPos3.Hands1[i] = 0
		} else {
			// 異なればその数
			pPos2.Hands1[i] = pPos0.Hands1[i]
			pPos3.Hands1[i] = pPos1.Hands1[i]
		}
	}

	// 位置
	for i := PCLOC_START; i < PCLOC_END; i += 1 {
		if pPos0.PieceLocations[i] == pPos1.PieceLocations[i] {
			// 等しければゼロ
			pPos2.PieceLocations[i] = 0
			pPos3.PieceLocations[i] = 0
		} else {
			// 異なればその数
			pPos2.PieceLocations[i] = pPos0.PieceLocations[i]
			pPos3.PieceLocations[i] = pPos1.PieceLocations[i]
		}
	}
}

// ２つのボードの違いを数えるぜ（＾～＾）
func errorBoard(pPos0 *Position, pPos1 *Position, pPos2 *Position, pPos3 *Position) int {
	diffBoard(pPos0, pPos1, pPos2, pPos3)

	errorNum := 0

	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPos2.Board[sq] != pPos3.Board[sq] {
			errorNum += 1
		}
	}

	// 駒台
	for i := HAND_IDX_START; i < HAND_IDX_END; i += 1 {
		if pPos2.Hands1[i] != pPos3.Hands1[i] {
			errorNum += 1
		}
	}

	// 位置
	if pPos2.PieceLocations[PCLOC_K1] != pPos3.PieceLocations[PCLOC_K1] {
		errorNum += 1
	}
	if pPos2.PieceLocations[PCLOC_K2] != pPos3.PieceLocations[PCLOC_K2] {
		errorNum += 1
	}

	// 位置（不安定注意）
	rook2 := []int{}
	rook3 := []int{}
	for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
		rook2 = append(rook2, int(pPos2.PieceLocations[i]))
		rook3 = append(rook3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(rook2)
	sort.Ints(rook3)
	for i := 0; i < len(rook2); i += 1 {
		if rook2[i] != rook3[i] {
			errorNum += 1
		}
	}

	// 位置（不安定注意）
	bishop2 := []int{}
	bishop3 := []int{}
	for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
		bishop2 = append(bishop2, int(pPos2.PieceLocations[i]))
		bishop3 = append(bishop3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(bishop2)
	sort.Ints(bishop3)
	for i := 0; i < len(bishop2); i += 1 {
		if bishop2[i] != bishop3[i] {
			errorNum += 1
		}
	}

	// 位置（不安定注意）
	lance2 := []int{}
	lance3 := []int{}
	for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
		lance2 = append(lance2, int(pPos2.PieceLocations[i]))
		lance3 = append(lance3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(lance2)
	sort.Ints(lance3)
	for i := 0; i < len(lance2); i += 1 {
		if lance2[i] != lance3[i] {
			errorNum += 1
		}
	}

	return errorNum
}
