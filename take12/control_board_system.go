package take12

import "fmt"

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

// DiffControl - 利きテーブルの差分計算
func (pControlBoardSys *ControlBoardSystem) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pControlBoardSys.ClearControlLayer(c3)

	for phase := 0; phase < 2; phase += 1 {
		cb3 := pControlBoardSys.Boards[phase][c3]
		cb1 := pControlBoardSys.Boards[phase][c1]
		cb2 := pControlBoardSys.Boards[phase][c2]
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				cb3.Board[from] = cb1.Board[from] - cb2.Board[from]

			}
		}
	}
}

// RecalculateControl - 利きの再計算
func (pControlBoardSys *ControlBoardSystem) RecalculateControl(pPos *Position, c1 ControlLayerT) {

	pControlBoardSys.ClearControlLayer(c1)

	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 && !pPos.IsEmptySq(from) {
			piece := pPos.Board[from]
			phase := Who(piece)
			sq_list := GenControl(pPos, from)

			cb1 := pControlBoardSys.Boards[phase-1][c1]
			for _, to := range sq_list {
				cb1.Board[to] += 1
			}

		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pControlBoardSys *ControlBoardSystem) MergeControlDiff() {
	cb0sum := pControlBoardSys.Boards[0][CONTROL_LAYER_SUM]
	cb1sum := pControlBoardSys.Boards[1][CONTROL_LAYER_SUM]
	for sq := Square(11); sq < BOARD_SIZE; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			// c=0 を除く
			for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
				cb0sum.Board[sq] += pControlBoardSys.Boards[0][c].Board[sq]
				cb1sum.Board[sq] += pControlBoardSys.Boards[1][c].Board[sq]
			}
		}
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pControlBoardSys *ControlBoardSystem) ClearControlDiff() {
	// c=0 を除く
	for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
		pControlBoardSys.ClearControlLayer(c)
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pControlBoardSys *ControlBoardSystem) AddControlDiff(pPos *Position, c ControlLayerT, from Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(Who(piece)) - 1
	// fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenControl(pPos, from)

	cb := pControlBoardSys.Boards[ph][c]
	for _, to := range sq_list {
		// fmt.Printf("Debug: ph=%d c=%d to=%d\n", ph, c, to)
		// 差分の方のテーブルを更新（＾～＾）
		cb.Board[to] += sign * 1
	}
}
