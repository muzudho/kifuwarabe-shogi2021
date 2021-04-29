// 利きのテスト
package take11

import (
	"fmt"
	"math"
	"math/rand"
)

// TestControl
func TestControl(pPosSys *PositionSystem, b PosLayerT) (bool, string) {
	pPosSys.ClearControlLayer(CONTROL_LAYER_TEST_COPY)
	pPosSys.ClearControlLayer(CONTROL_LAYER_TEST_ERROR)

	// 利きをコピー
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			pPosSys.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] = pPosSys.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
		}
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPosSys, b)
	move_total := len(move_list)

	for move_seq, move := range move_list {
		// その手を指してみるぜ（＾～＾）
		pPosSys.DoMove(b, move)

		// すぐ戻すぜ（＾～＾）
		pPosSys.UndoMove(b)

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
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			diff := pPosSys.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] - pPosSys.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
			pPosSys.ControlBoards[phase][CONTROL_LAYER_TEST_ERROR][sq] = diff
			if diff != 0 {
				is_error = true
			}
		}
	}

	return is_error
}

// SumAbsControl - 利きテーブルの各マスを絶対値にし、その総和を返します
func SumAbsControl(pPosSys *PositionSystem, layer1 int) [2]int {

	sumList := [2]int{0, 0}

	for phase := 0; phase < 2; phase += 1 {
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				sumList[phase] += int(math.Abs(float64(pPosSys.ControlBoards[phase][layer1][from])))

			}
		}
	}

	return sumList
}

// ShuffleBoard - 盤上の駒、持ち駒をシャッフルします
// ゲーム中にはできない動きをするので、利きの計算は無視します。
// 最後に利きは再計算します
func ShuffleBoard(pPosSys *PositionSystem, b PosLayerT) {

	// 駒の数を数えます
	countList1 := CountAllPieces(pPosSys, b)

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
					piece := pPosSys.PPosition[b].Board[sq]
					if piece != PIECE_EMPTY {
						phase := Who(piece)
						pieceType := What(piece)

						ok := false
						switch phase {
						case FIRST:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPosSys.PPosition[b].Hands[HAND_R1_IDX] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPosSys.PPosition[b].Hands[HAND_B1_IDX] += 1
								ok = true
							case PIECE_TYPE_G:
								pPosSys.PPosition[b].Hands[HAND_G1_IDX] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPosSys.PPosition[b].Hands[HAND_S1_IDX] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPosSys.PPosition[b].Hands[HAND_N1_IDX] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPosSys.PPosition[b].Hands[HAND_L1_IDX] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPosSys.PPosition[b].Hands[HAND_P1_IDX] += 1
								ok = true
							default:
								// Ignored
							}
						case SECOND:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPosSys.PPosition[b].Hands[HAND_R2_IDX] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPosSys.PPosition[b].Hands[HAND_B2_IDX] += 1
								ok = true
							case PIECE_TYPE_G:
								pPosSys.PPosition[b].Hands[HAND_G2_IDX] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPosSys.PPosition[b].Hands[HAND_S2_IDX] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPosSys.PPosition[b].Hands[HAND_N2_IDX] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPosSys.PPosition[b].Hands[HAND_L2_IDX] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPosSys.PPosition[b].Hands[HAND_P2_IDX] += 1
								ok = true
							default:
								// Ignored
							}
						default:
							panic(fmt.Errorf("Uknown phase=%d", phase))
						}

						if ok {
							pPosSys.PPosition[b].Board[sq] = PIECE_EMPTY
						}
					}

				}
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPosSys, b)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("Shuffle: (1) countError=%d", countError))
		}

		// 駒台から盤の方向
		for hand_index := HAND_IDX_START; hand_index < HAND_IDX_END; hand_index += 1 {
			num := pPosSys.PPosition[b].Hands[hand_index]
			if num > 0 {
				sq := Square(rand.Intn(100))
				// うまく空マスなら移動成功
				if OnBoard(sq) && pPosSys.PPosition[b].IsEmptySq(sq) {
					pPosSys.PPosition[b].Board[sq] = HandPieceMap[hand_index]
					pPosSys.PPosition[b].Hands[hand_index] -= 1
				}
			}
		}

		// 駒の数を数えます
		countList2 = CountAllPieces(pPosSys, b)
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
		if OnBoard(sq1) && OnBoard(sq2) && !pPosSys.PPosition[b].IsEmptySq(sq1) {
			piece := pPosSys.PPosition[b].Board[sq1]
			// 位置スワップ
			pPosSys.PPosition[b].Board[sq1] = pPosSys.PPosition[b].Board[sq2]
			pPosSys.PPosition[b].Board[sq2] = piece

			// 成／不成 変更
			promote := Square(rand.Intn(10))
			if promote == 0 {
				pPosSys.PPosition[b].Board[sq2] = Promote(pPosSys.PPosition[b].Board[sq2])
			} else if promote == 1 {
				pPosSys.PPosition[b].Board[sq2] = Demote(pPosSys.PPosition[b].Board[sq2])
			}

			// 駒の先後変更（玉除く）
			piece = pPosSys.PPosition[b].Board[sq2]
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

				pPosSys.PPosition[b].Board[sq2] = PieceFromPhPt(phase, pieceType)
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPosSys, b)
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
	G.Chat.Debug(pPosSys.Sprint(b))

	if false {
		var countList [8]int

		if true {
			countList = [8]int{}

			// 盤上
			for rank := Square(1); rank < 10; rank += 1 {
				for file := Square(9); file > 0; file -= 1 {
					sq := SquareFrom(file, rank)

					fmt.Printf("%s,", pPosSys.PPosition[b].Board[sq].ToCode())

					piece := What(pPosSys.PPosition[b].Board[sq])
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
			countList[1] += pPosSys.PPosition[b].Hands[0] + pPosSys.PPosition[b].Hands[7]
			countList[2] += pPosSys.PPosition[b].Hands[1] + pPosSys.PPosition[b].Hands[8]
			countList[3] += pPosSys.PPosition[b].Hands[2] + pPosSys.PPosition[b].Hands[9]
			countList[4] += pPosSys.PPosition[b].Hands[3] + pPosSys.PPosition[b].Hands[10]
			countList[5] += pPosSys.PPosition[b].Hands[4] + pPosSys.PPosition[b].Hands[11]
			countList[6] += pPosSys.PPosition[b].Hands[5] + pPosSys.PPosition[b].Hands[12]
			countList[7] += pPosSys.PPosition[b].Hands[6] + pPosSys.PPosition[b].Hands[13]
		} else {
			countList = CountAllPieces(pPosSys, b)
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
		ShowAllPiecesCount(pPosSys, b)
	}

	// position sfen 文字列を取得
	command := pPosSys.SprintSfen(b)
	G.Chat.Debug("#command=%s", command)

	// 利きの再計算もやってくれる
	pPosSys.ReadPosition(b, command)

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	G.Chat.Debug(pPosSys.Sprint(b))
	ShowAllPiecesCount(pPosSys, b)
	command2 := pPosSys.SprintSfen(b)
	G.Chat.Debug("#command2=%s", command2)

	// 駒の数を数えます
	countList2 := CountAllPieces(pPosSys, b)
	countError := CountErrorCountLists(countList1, countList2)
	if countError != 0 {
		panic(fmt.Errorf("Shuffle: (4) countError=%d", countError))
	}
}

// CountAllPieces - 駒の数を確認するぜ（＾～＾）
func CountAllPieces(pPosSys *PositionSystem, b PosLayerT) [8]int {

	countList := [8]int{}

	// 盤上
	for rank := Square(1); rank < 10; rank += 1 {
		for file := Square(9); file > 0; file -= 1 {
			sq := SquareFrom(file, rank)

			piece := What(pPosSys.PPosition[b].Board[sq])
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
	countList[1] += pPosSys.PPosition[b].Hands[0] + pPosSys.PPosition[b].Hands[7]
	countList[2] += pPosSys.PPosition[b].Hands[1] + pPosSys.PPosition[b].Hands[8]
	countList[3] += pPosSys.PPosition[b].Hands[2] + pPosSys.PPosition[b].Hands[9]
	countList[4] += pPosSys.PPosition[b].Hands[3] + pPosSys.PPosition[b].Hands[10]
	countList[5] += pPosSys.PPosition[b].Hands[4] + pPosSys.PPosition[b].Hands[11]
	countList[6] += pPosSys.PPosition[b].Hands[5] + pPosSys.PPosition[b].Hands[12]
	countList[7] += pPosSys.PPosition[b].Hands[6] + pPosSys.PPosition[b].Hands[13]

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
func copyBoard(pPosSys *PositionSystem, b0 PosLayerT, b1 PosLayerT) {
	for sq := 0; sq < 100; sq += 1 {
		pPosSys.PPosition[b1].Board[sq] = pPosSys.PPosition[b0].Board[sq]
	}

	pPosSys.PPosition[b1].Hands = pPosSys.PPosition[b0].Hands
	for i := PCLOC_START; i < PCLOC_END; i += 1 {
		pPosSys.PPosition[b1].PieceLocations[i] = pPosSys.PPosition[b0].PieceLocations[i]
	}
}

// copyBoard - 盤[0] を 盤[1] で異なるマスを 盤[2] 盤[3] にセットします
func diffBoard(pPosSys *PositionSystem, b0 PosLayerT, b1 PosLayerT, b2 PosLayerT, b3 PosLayerT) {
	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPosSys.PPosition[b1].Board[sq] == pPosSys.PPosition[b0].Board[sq] {
			// 等しければ空マス
			pPosSys.PPosition[b2].Board[sq] = PIECE_EMPTY
			pPosSys.PPosition[b3].Board[sq] = PIECE_EMPTY

		} else {
			// 異なったら
			pPosSys.PPosition[b2].Board[sq] = pPosSys.PPosition[b0].Board[sq]
			pPosSys.PPosition[b3].Board[sq] = pPosSys.PPosition[b1].Board[sq]
		}
	}

	// 駒台
	for i := HAND_IDX_START; i < HAND_IDX_END; i += 1 {
		if pPosSys.PPosition[b0].Hands[i] == pPosSys.PPosition[b1].Hands[i] {
			// 等しければゼロ
			pPosSys.PPosition[b2].Hands[i] = 0
			pPosSys.PPosition[b3].Hands[i] = 0
		} else {
			// 異なればその数
			pPosSys.PPosition[b2].Hands[i] = pPosSys.PPosition[b0].Hands[i]
			pPosSys.PPosition[b3].Hands[i] = pPosSys.PPosition[b1].Hands[i]
		}
	}

	// 位置
	for i := PCLOC_START; i < PCLOC_END; i += 1 {
		if pPosSys.PPosition[b0].PieceLocations[i] == pPosSys.PPosition[b1].PieceLocations[i] {
			// 等しければゼロ
			pPosSys.PPosition[b2].PieceLocations[i] = 0
			pPosSys.PPosition[b3].PieceLocations[i] = 0
		} else {
			// 異なればその数
			pPosSys.PPosition[b2].PieceLocations[i] = pPosSys.PPosition[b0].PieceLocations[i]
			pPosSys.PPosition[b3].PieceLocations[i] = pPosSys.PPosition[b1].PieceLocations[i]
		}
	}
}

// ２つのボードの違いを数えるぜ（＾～＾）
func errorBoard(pPosSys *PositionSystem, b0 PosLayerT, b1 PosLayerT, b2 PosLayerT, b3 PosLayerT) int {
	diffBoard(pPosSys, b0, b1, b2, b3)

	errorNum := 0

	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPosSys.PPosition[b2].Board[sq] != pPosSys.PPosition[b3].Board[sq] {
			errorNum += 1
		}
	}

	// 駒台
	for i := HAND_IDX_START; i < HAND_IDX_END; i += 1 {
		if pPosSys.PPosition[b2].Hands[i] != pPosSys.PPosition[b3].Hands[i] {
			errorNum += 1
		}
	}

	// 位置
	for i := PCLOC_START; i < PCLOC_END; i += 1 {
		if pPosSys.PPosition[b2].PieceLocations[i] != pPosSys.PPosition[b3].PieceLocations[i] {
			errorNum += 1
		}
	}

	return errorNum
}
