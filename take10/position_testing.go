// 利きのテスト
package take10

import "fmt"

// PositionTesting - 局面オブジェクトをテストするためのものです
type PositionTesting struct {
	// 利きテーブル [0]先手 [1]後手
	// マスへの利き数が入っています
	ControlBoards [2][BOARD_SIZE]int8
}

// NewPositionTesting - テストの作成
func NewPositionTesting() *PositionTesting {
	var pPosT = new(PositionTesting)
	pPosT.ControlBoards = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}}
	return pPosT
}

func (pPosT *PositionTesting) Test(pPos *Position) (bool, string) {
	// 利きをコピー
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			pPosT.ControlBoards[phase][sq] = pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
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
		ok, message := pPosT.Check(pPos, move_seq, move_total, move)
		if !ok {
			return ok, message
		}
	}

	return true, ""
}

// Check - 元に戻っていればOK（＾～＾）
func (pPosT *PositionTesting) Check(pPos *Position, move_seq int, move_total int, move Move) (bool, string) {
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			if pPosT.ControlBoards[phase][sq] != pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq] {
				return false, fmt.Sprintf("Error! move_seq=(%d/%d) move=%s phase=%d sq=%d", move_seq, move_total, move.ToCode(), phase, sq)
			}
		}
	}

	return true, ""
}
