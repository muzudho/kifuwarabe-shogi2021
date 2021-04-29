// 利きボード
package take11

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
	CONTROL_LAYER_TEST_COPY          = ControlLayerT(10) // テスト用
	CONTROL_LAYER_TEST_ERROR         = ControlLayerT(11) // テスト用
	CONTROL_LAYER_TEST_RECALCULATION = ControlLayerT(12) // テスト用 再計算
	CONTROL_LAYER_DIFF_START         = ControlLayerT(1)
	CONTROL_LAYER_DIFF_END           = ControlLayerT(10) // この数を含まない。テスト用も含まない
	CONTROL_LAYER_ALL_SIZE           = 13                // この数を含まない
)

// GetControlLayerName - 利きボードのレイヤーの名前
func GetControlLayerName(c ControlLayerT) string {
	switch c {
	case CONTROL_LAYER_SUM:
		return "Sum"
	case CONTROL_LAYER_DIFF_ROOK_OFF:
		return "RookOff"
	case CONTROL_LAYER_DIFF_BISHOP_OFF:
		return "BishopOff"
	case CONTROL_LAYER_DIFF_LANCE_OFF:
		return "LanceOff"
	case CONTROL_LAYER_DIFF_PUT:
		return "Put"
	case CONTROL_LAYER_DIFF_REMOVE:
		return "Remove"
	case CONTROL_LAYER_DIFF_CAPTURED:
		return "Captured"
	case CONTROL_LAYER_DIFF_LANCE_ON:
		return "LanceOn"
	case CONTROL_LAYER_DIFF_BISHOP_ON:
		return "BishopOn"
	case CONTROL_LAYER_DIFF_ROOK_ON:
		return "RookOn"
	case CONTROL_LAYER_TEST_COPY:
		return "TestCopy"
	case CONTROL_LAYER_TEST_ERROR:
		return "TestError"
	case CONTROL_LAYER_TEST_RECALCULATION:
		return "TestRecalc"
	default:
		panic(fmt.Errorf("Unknown controlLayer=%d", c))
	}
}

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlRook(b BoardLayerT, c ControlLayerT, sign int8, excludeFrom Square) {
	for _, from := range pPos.RookLocations {
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(b, from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPos.AddControlDiff(b, c, from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlBishop(b BoardLayerT, c ControlLayerT, sign int8, excludeFrom Square) {
	for _, from := range pPos.BishopLocations {
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(b, from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPos.AddControlDiff(b, c, from, sign)
		}
	}
}

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlLance(b BoardLayerT, c ControlLayerT, sign int8, excludeFrom Square) {
	for _, from := range pPos.LanceLocations {

		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(b, from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[b][from]) { // 杏は除外
			pPos.AddControlDiff(b, c, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiff(b BoardLayerT, c ControlLayerT, from Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[b][from]
	if piece == PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(Who(piece)) - 1
	// fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenControl(pPos, b, from)

	for _, to := range sq_list {
		// fmt.Printf("Debug: to=%d\n", to)
		// 差分の方のテーブルを更新（＾～＾）
		pPos.ControlBoards[ph][c][to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPos *Position) ClearControlDiff() {
	// c=0 を除く
	for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
		pPos.ClearControlLayer(c)
	}
}

func (pPos *Position) ClearControlLayer(c ControlLayerT) {
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			pPos.ControlBoards[0][c][sq] = 0
			pPos.ControlBoards[1][c][sq] = 0
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPos *Position) MergeControlDiff() {
	for sq := Square(11); sq < BOARD_SIZE; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			// c=0 を除く
			for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
				pPos.ControlBoards[0][CONTROL_LAYER_SUM][sq] += pPos.ControlBoards[0][c][sq]
				pPos.ControlBoards[1][CONTROL_LAYER_SUM][sq] += pPos.ControlBoards[1][c][sq]
			}
		}
	}
}

// RecalculateControl - 利きの再計算
func (pPos *Position) RecalculateControl(b BoardLayerT, c1 ControlLayerT) {

	pPos.ClearControlLayer(c1)

	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 && !pPos.IsEmptySq(b, from) {
			piece := pPos.Board[b][from]
			phase := Who(piece)
			sq_list := GenControl(pPos, b, from)

			for _, to := range sq_list {
				pPos.ControlBoards[phase-1][c1][to] += 1
			}

		}
	}
}

// DiffControl - 利きテーブルの差分計算
func (pPos *Position) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pPos.ClearControlLayer(c3)

	for phase := 0; phase < 2; phase += 1 {
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				pPos.ControlBoards[phase][c3][from] = pPos.ControlBoards[phase][c1][from] - pPos.ControlBoards[phase][c2][from]

			}
		}
	}
}
