package take13

import "fmt"

// 利きテーブル・インデックス型
type ControlLayerT int

const (
	// 先手用
	CONTROL_LAYER_SUM1             = ControlLayerT(0) // 先手の利きボード
	CONTROL_LAYER_DIFF1_ROOK_OFF   = ControlLayerT(1) // 飛の利き 負(差分)
	CONTROL_LAYER_DIFF1_BISHOP_OFF = ControlLayerT(2) // 角の利き 負(差分)
	CONTROL_LAYER_DIFF1_LANCE_OFF  = ControlLayerT(3) // 香の利き 負(差分)
	CONTROL_LAYER_DIFF1_PUT        = ControlLayerT(4) // 打とか指すとか
	CONTROL_LAYER_DIFF1_REMOVE     = ControlLayerT(5)
	CONTROL_LAYER_DIFF1_CAPTURED   = ControlLayerT(6)
	CONTROL_LAYER_DIFF1_LANCE_ON   = ControlLayerT(7) // 香の利き 正(差分)
	CONTROL_LAYER_DIFF1_BISHOP_ON  = ControlLayerT(8) // 角の利き 正(差分)
	CONTROL_LAYER_DIFF1_ROOK_ON    = ControlLayerT(9) // 飛の利き 正(差分)
	// 後手用
	CONTROL_LAYER_SUM2             = ControlLayerT(10) // 後手の利きボード
	CONTROL_LAYER_DIFF2_ROOK_OFF   = ControlLayerT(11)
	CONTROL_LAYER_DIFF2_BISHOP_OFF = ControlLayerT(12)
	CONTROL_LAYER_DIFF2_LANCE_OFF  = ControlLayerT(13)
	CONTROL_LAYER_DIFF2_PUT        = ControlLayerT(14) // 打とか指すとか
	CONTROL_LAYER_DIFF2_REMOVE     = ControlLayerT(15)
	CONTROL_LAYER_DIFF2_CAPTURED   = ControlLayerT(16)
	CONTROL_LAYER_DIFF2_LANCE_ON   = ControlLayerT(17)
	CONTROL_LAYER_DIFF2_BISHOP_ON  = ControlLayerT(18)
	CONTROL_LAYER_DIFF2_ROOK_ON    = ControlLayerT(19)
	// テスト（先手用）
	CONTROL_LAYER_TEST_COPY1          = ControlLayerT(20) // テスト用
	CONTROL_LAYER_TEST_ERROR1         = ControlLayerT(21) // テスト用
	CONTROL_LAYER_TEST_RECALCULATION1 = ControlLayerT(22) // テスト用 再計算
	// テスト（後手用）
	CONTROL_LAYER_TEST_COPY2          = ControlLayerT(23)
	CONTROL_LAYER_TEST_ERROR2         = ControlLayerT(24)
	CONTROL_LAYER_TEST_RECALCULATION2 = ControlLayerT(25)
	// 評価値用
	CONTROL_LAYER_EVAL1 = ControlLayerT(26) // 評価関数用
	CONTROL_LAYER_EVAL2 = ControlLayerT(27) // 評価関数用
	CONTROL_LAYER_EVAL3 = ControlLayerT(28) // 評価関数用
	// 計測
	CONTROL_LAYER_DIFF_TYPE_SIZE = ControlLayerT(9)
	CONTROL_LAYER_DIFF1_START    = ControlLayerT(1)
	CONTROL_LAYER_DIFF1_END      = CONTROL_LAYER_DIFF1_START + CONTROL_LAYER_DIFF_TYPE_SIZE // この数を含まない。テスト用も含まない
	CONTROL_LAYER_DIFF2_START    = ControlLayerT(11)
	CONTROL_LAYER_DIFF2_END      = CONTROL_LAYER_DIFF2_START + CONTROL_LAYER_DIFF_TYPE_SIZE // この数を含まない。テスト用も含まない
	CONTROL_LAYER_ALL_SIZE       = ControlLayerT(29)                                        // この数を含まない
)

// ControlBoardSystem - 利きボード・システム
type ControlBoardSystem struct {
	// マスへの利き数、または差分が入っています。デバッグ目的で無駄に分けてるんだけどな（＾～＾）
	Boards [CONTROL_LAYER_ALL_SIZE]*ControlBoard
}

func NewControlBoardSystem() *ControlBoardSystem {
	cbsys := new(ControlBoardSystem)

	cbsys.Boards = [CONTROL_LAYER_ALL_SIZE]*ControlBoard{
		// 先手用
		NewControlBoard("Sum1"),
		NewControlBoard("RookOff1"),
		NewControlBoard("BishopOff1"),
		NewControlBoard("LanceOff1"),
		NewControlBoard("Put1"),
		NewControlBoard("Remove1"),
		NewControlBoard("Captured1"),
		NewControlBoard("LanceOn1"),
		NewControlBoard("BishopOn1"),
		NewControlBoard("RookOn1"),
		// 後手用
		NewControlBoard("Sum2"),
		NewControlBoard("RookOff2"),
		NewControlBoard("BishopOff2"),
		NewControlBoard("LanceOff2"),
		NewControlBoard("Put2"),
		NewControlBoard("Remove2"),
		NewControlBoard("Captured2"),
		NewControlBoard("LanceOn2"),
		NewControlBoard("BishopOn2"),
		NewControlBoard("RookOn2"),
		// テスト（先手用）
		NewControlBoard("TestCopy1"),
		NewControlBoard("TestError1"),
		NewControlBoard("TestRecalc1"),
		// テスト（後手用）
		NewControlBoard("TestCopy2"),
		NewControlBoard("TestError2"),
		NewControlBoard("TestRecalc2"),
		// 評価値用
		NewControlBoard("Eval1"),
		NewControlBoard("Eval2"),
		NewControlBoard("Eval3"),
	}

	return cbsys
}

// ClearControlLayer - 利きボードのクリアー
func (pCtrlBrdSys *ControlBoardSystem) ClearControlLayer1(ph1_c ControlLayerT, ph2_c ControlLayerT) {
	pCtrlBrdSys.Boards[ph1_c].Clear()
	pCtrlBrdSys.Boards[ph2_c].Clear()
}

// DiffControl - 利きテーブルの差分計算
func (pCtrlBrdSys *ControlBoardSystem) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pCtrlBrdSys.Boards[c3].Clear()

	cb3 := pCtrlBrdSys.Boards[c3]
	cb1 := pCtrlBrdSys.Boards[c1]
	cb2 := pCtrlBrdSys.Boards[c2]
	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 {

			cb3.Board1[from] = cb1.Board1[from] - cb2.Board1[from]

		}
	}
}

// RecalculateControl - 利きの再計算
func (pCtrlBrdSys *ControlBoardSystem) RecalculateControl(
	pPos *Position, ph1_c1 ControlLayerT, ph2_c1 ControlLayerT) {

	pCtrlBrdSys.Boards[ph1_c1].Clear()
	pCtrlBrdSys.Boards[ph2_c1].Clear()

	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 && !pPos.IsEmptySq(from) {
			piece := pPos.Board[from]
			phase := Who(piece)
			sq_list := GenControl(pPos, from)

			pCB := ControllBoardFromPhase(phase, pCtrlBrdSys.Boards[ph1_c1], pCtrlBrdSys.Boards[ph2_c1])

			for _, to := range sq_list {
				pCB.Board1[to] += 1
			}
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pCtrlBrdSys *ControlBoardSystem) MergeControlDiff() {
	cb0sum := pCtrlBrdSys.Boards[CONTROL_LAYER_SUM1]
	cb1sum := pCtrlBrdSys.Boards[CONTROL_LAYER_SUM2]
	for sq := Square(11); sq < BOARD_SIZE; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			// c=0 を除く
			for c1 := CONTROL_LAYER_DIFF1_START; c1 < CONTROL_LAYER_DIFF1_END; c1 += 1 {
				cb0sum.Board1[sq] += pCtrlBrdSys.Boards[c1].Board1[sq]
			}
			for c2 := CONTROL_LAYER_DIFF2_START; c2 < CONTROL_LAYER_DIFF2_END; c2 += 1 {
				cb1sum.Board1[sq] += pCtrlBrdSys.Boards[c2].Board1[sq]
			}
		}
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pCtrlBrdSys *ControlBoardSystem) ClearControlDiff() {
	// c=0 を除く
	for c1 := CONTROL_LAYER_DIFF1_START; c1 < CONTROL_LAYER_DIFF1_END; c1 += 1 {
		pCtrlBrdSys.Boards[c1].Clear()
	}
	for c2 := CONTROL_LAYER_DIFF2_START; c2 < CONTROL_LAYER_DIFF2_END; c2 += 1 {
		pCtrlBrdSys.Boards[c2].Clear()
	}
}

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pCtrlBrdSys *ControlBoardSystem) AddControlLance(pPos *Position,
	ph1_c ControlLayerT, ph2_c ControlLayerT, sign int16, excludeFrom Square) {
	for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外

			piece := pPos.Board[from]
			ValidateThereArePieceIn(pPos, from)
			phase := Who(piece)
			pCB := ControllBoardFromPhase(phase, pCtrlBrdSys.Boards[ph1_c], pCtrlBrdSys.Boards[ph2_c])
			pCB.AddControlDiff(GenControl(pPos, from), from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pCtrlBrdSys *ControlBoardSystem) AddControlBishop(pPos *Position,
	ph1_c ControlLayerT, ph2_c ControlLayerT, sign int16, excludeFrom Square) {
	for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外

			piece := pPos.Board[from]
			ValidateThereArePieceIn(pPos, from)
			phase := Who(piece)
			pCB := ControllBoardFromPhase(phase, pCtrlBrdSys.Boards[ph1_c], pCtrlBrdSys.Boards[ph2_c])
			pCB.AddControlDiff(GenControl(pPos, from), from, sign)
		}
	}
}

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pCtrlBrdSys *ControlBoardSystem) AddControlRook(pPos *Position,
	pPh1_CB *ControlBoard, pPh2_CB *ControlBoard, sign int16, excludeFrom Square) {
	for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外

			piece := pPos.Board[from]
			ValidateThereArePieceIn(pPos, from)
			phase := Who(piece)
			pCB := ControllBoardFromPhase(phase, pPh1_CB, pPh2_CB)
			pCB.AddControlDiff(GenControl(pPos, from), from, sign)
		}
	}
}

func ControllBoardFromPhase(
	phase Phase, pPh1_CB *ControlBoard, pPh2_CB *ControlBoard) *ControlBoard {

	// fmt.Printf("Debug: phase=%d\n", phase)
	switch phase {
	case FIRST:
		return pPh1_CB
	case SECOND:
		return pPh2_CB
	default:
		panic(fmt.Errorf("Unknown phase=%d", phase))
	}
}
