// 利きボード
package take12

import "fmt"

// 利きテーブル・インデックス型
type ControlLayerT int

const (
	CONTROL_LAYER_SUM                = ControlLayerT(0)
	CONTROL_LAYER_DIFF_ROOK_OFF      = ControlLayerT(1)
	CONTROL_LAYER_DIFF_BISHOP_OFF    = ControlLayerT(2)
	CONTROL_LAYER_DIFF_LANCE_OFF     = ControlLayerT(3)
	CONTROL_LAYER_DIFF_PUT           = ControlLayerT(4) // 打とか指すとか
	CONTROL_LAYER_DIFF_REMOVE        = ControlLayerT(5)
	CONTROL_LAYER_DIFF_CAPTURED      = ControlLayerT(6)
	CONTROL_LAYER_DIFF_LANCE_ON      = ControlLayerT(7)
	CONTROL_LAYER_DIFF_BISHOP_ON     = ControlLayerT(8)
	CONTROL_LAYER_DIFF_ROOK_ON       = ControlLayerT(9)
	CONTROL_LAYER_EVAL               = ControlLayerT(10) // 評価関数用
	CONTROL_LAYER_TEST_COPY          = ControlLayerT(11) // テスト用
	CONTROL_LAYER_TEST_ERROR         = ControlLayerT(12) // テスト用
	CONTROL_LAYER_TEST_RECALCULATION = ControlLayerT(13) // テスト用 再計算
	CONTROL_LAYER_DIFF_START         = ControlLayerT(1)
	CONTROL_LAYER_DIFF_END           = ControlLayerT(10) // この数を含まない。テスト用も含まない
	CONTROL_LAYER_ALL_SIZE           = 14                // この数を含まない
)

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlRook(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlBishop(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlLance(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlDiff(pPos *Position, c ControlLayerT, from Square, sign int8) {
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

	cb := pPosSys.PControlBoardSystem.Boards[ph][c]
	for _, to := range sq_list {
		// fmt.Printf("Debug: ph=%d c=%d to=%d\n", ph, c, to)
		// 差分の方のテーブルを更新（＾～＾）
		cb.Board[to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPosSys *PositionSystem) ClearControlDiff() {
	// c=0 を除く
	for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
		pPosSys.ClearControlLayer(c)
	}
}

func (pPosSys *PositionSystem) ClearControlLayer(c ControlLayerT) {
	cb0 := pPosSys.PControlBoardSystem.Boards[0][c]
	cb1 := pPosSys.PControlBoardSystem.Boards[1][c]
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			cb0.Board[sq] = 0
			cb1.Board[sq] = 0
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPosSys *PositionSystem) MergeControlDiff() {
	cb0sum := pPosSys.PControlBoardSystem.Boards[0][CONTROL_LAYER_SUM]
	cb1sum := pPosSys.PControlBoardSystem.Boards[1][CONTROL_LAYER_SUM]
	for sq := Square(11); sq < BOARD_SIZE; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			// c=0 を除く
			for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
				cb0sum.Board[sq] += pPosSys.PControlBoardSystem.Boards[0][c].Board[sq]
				cb1sum.Board[sq] += pPosSys.PControlBoardSystem.Boards[1][c].Board[sq]
			}
		}
	}
}

// RecalculateControl - 利きの再計算
func (pPosSys *PositionSystem) RecalculateControl(pPos *Position, c1 ControlLayerT) {

	pPosSys.ClearControlLayer(c1)

	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 && !pPos.IsEmptySq(from) {
			piece := pPos.Board[from]
			phase := Who(piece)
			sq_list := GenControl(pPos, from)

			cb1 := pPosSys.PControlBoardSystem.Boards[phase-1][c1]
			for _, to := range sq_list {
				cb1.Board[to] += 1
			}

		}
	}
}

// DiffControl - 利きテーブルの差分計算
func (pPosSys *PositionSystem) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pPosSys.ClearControlLayer(c3)

	for phase := 0; phase < 2; phase += 1 {
		cb3 := pPosSys.PControlBoardSystem.Boards[phase][c3]
		cb1 := pPosSys.PControlBoardSystem.Boards[phase][c1]
		cb2 := pPosSys.PControlBoardSystem.Boards[phase][c2]
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				cb3.Board[from] = cb1.Board[from] - cb2.Board[from]

			}
		}
	}
}
