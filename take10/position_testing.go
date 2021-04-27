// 利きのテスト
package take10

import "fmt"

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
