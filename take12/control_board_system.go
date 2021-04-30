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
func (pControlBoardSys *ControlBoardSystem) ClearControlLayer1(c ControlLayerT) {
	cb0 := pControlBoardSys.Boards[0][c]
	cb1 := pControlBoardSys.Boards[1][c]
	cb0.Clear()
	cb1.Clear()
}

// DiffControl - 利きテーブルの差分計算
func (pControlBoardSys *ControlBoardSystem) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pControlBoardSys.Boards[FIRST-1][c3].Clear()
	pControlBoardSys.Boards[SECOND-1][c3].Clear()

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

	pControlBoardSys.Boards[FIRST-1][c1].Clear()
	pControlBoardSys.Boards[SECOND-1][c1].Clear()

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
		pControlBoardSys.Boards[FIRST-1][c].Clear()
		pControlBoardSys.Boards[SECOND-1][c].Clear()
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

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pControlBoardSys *ControlBoardSystem) AddControlLance(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外
			pControlBoardSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pControlBoardSys *ControlBoardSystem) AddControlBishop(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pControlBoardSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pControlBoardSys *ControlBoardSystem) AddControlRook(pPos *Position, c ControlLayerT, sign int8, excludeFrom Square) {
	for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pControlBoardSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// 将棋盤の内側をスキャンします。
var scanningLine = []Square{
	82, 72, 62, 52, 42, 32, 22, 12,
	83, 73, 63, 53, 43, 33, 23, 13,
	84, 74, 64, 54, 44, 34, 24, 14,
	85, 75, 65, 55, 45, 35, 25, 15,
	86, 76, 66, 56, 46, 36, 26, 16,
	87, 77, 67, 57, 47, 37, 27, 17,
	88, 78, 68, 58, 48, 38, 28, 18,
}

// ブラシの太さ
var brushingArea = []int32{
	9, -1, -11,
	10, 0, -10,
	11, 1, -9}

// WaterColor - 水で薄めたような評価値にします
// pCB1 - pCB2 = pCB3
func WaterColor(pCB1 *ControlBoard, pCB2 *ControlBoard, pCB3 *ControlBoard) {
	// 将棋盤の内側をスキャンします。

	pCB3.Clear()

	for _, sq1 := range scanningLine {
		// ブラシの面積分の利きを総和します
		var sum int8 = 0
		for _, rel := range brushingArea {
			sq2 := Square(int32(sq1) + rel)
			sum += pCB1.Board[sq2] - pCB2.Board[sq2]
		}
		// 総和したものを、結果表に上乗せします
		for _, rel := range brushingArea {
			sq2 := Square(int32(sq1) + rel)
			pCB3.Board[sq2] += sum
		}
	}
}
