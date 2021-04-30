package take12

// ControlBoardSystem - 利きボード・システム
type ControlBoardSystem struct {
	// マスへの利き数、または差分が入っています。デバッグ目的で無駄に分けてるんだけどな（＾～＾）
	// 利きテーブル [0]先手 [1]後手
	// [0] 利き
	// [1] 飛の利き引く(差分)
	// [2] 角の利き引く(差分)
	// [3] 香の利き引く(差分)
	// [4] ムーブ用(差分)
	// [5] ムーブ用(差分)
	// [6] ムーブ用(差分)
	// [7] 香の利き戻す(差分)
	// [8] 角の利き戻す(差分)
	// [9] 飛の利き戻す(差分)
	// [10] テスト用
	// [11] テスト用
	// [12] テスト用(再計算)
	Boards [PHASE_ARRAY_SIZE][CONTROL_LAYER_ALL_SIZE]*ControlBoard
}

func NewControlBoardSystem() *ControlBoardSystem {
	cbsys := new(ControlBoardSystem)

	cbsys.Boards = [PHASE_ARRAY_SIZE][CONTROL_LAYER_ALL_SIZE]*ControlBoard{{
		NewControlBoard("Sum"),
		NewControlBoard("RookOff"),
		NewControlBoard("BishopOff"),
		NewControlBoard("LanceOff"),
		NewControlBoard("Put"),
		NewControlBoard("Remove"),
		NewControlBoard("Captured"),
		NewControlBoard("LanceOn"),
		NewControlBoard("BishopOn"),
		NewControlBoard("RookOn"),
		NewControlBoard("Eval"),
		NewControlBoard("TestCopy"),
		NewControlBoard("TestError"),
		NewControlBoard("TestRecalc"),
	}, {
		NewControlBoard("Sum"),
		NewControlBoard("RookOff"),
		NewControlBoard("BishopOff"),
		NewControlBoard("LanceOff"),
		NewControlBoard("Put"),
		NewControlBoard("Remove"),
		NewControlBoard("Captured"),
		NewControlBoard("LanceOn"),
		NewControlBoard("BishopOn"),
		NewControlBoard("RookOn"),
		NewControlBoard("Eval"),
		NewControlBoard("TestCopy"),
		NewControlBoard("TestError"),
		NewControlBoard("TestRecalc"),
	}}

	return cbsys
}

// ClearControlLayer - 利きボードのクリアー
func (pControlBoardSys *ControlBoardSystem) ClearControlLayer(c ControlLayerT) {
	cb0 := pControlBoardSys.Boards[0][c]
	cb1 := pControlBoardSys.Boards[1][c]
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			cb0.Board[sq] = 0
			cb1.Board[sq] = 0
		}
	}
}
