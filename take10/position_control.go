// 利きボード
package take10

import "fmt"

const (
	CONTROL_LAYER_SUM = iota
	CONTROL_LAYER_DIFF_ROOK_OFF
	CONTROL_LAYER_DIFF_BISHOP_OFF
	CONTROL_LAYER_DIFF_LANCE_OFF
	CONTROL_LAYER_DIFF_PUT // 打とか指すとか
	CONTROL_LAYER_DIFF_REMOVE
	CONTROL_LAYER_DIFF_CAPTURED
	CONTROL_LAYER_DIFF_LANCE_ON
	CONTROL_LAYER_DIFF_BISHOP_ON
	CONTROL_LAYER_DIFF_ROOK_ON
	CONTROL_LAYER_TEST_COPY          // テスト用
	CONTROL_LAYER_TEST_ERROR         // テスト用
	CONTROL_LAYER_TEST_RECALCULATION // テスト用 再計算
	CONTROL_LAYER_DIFF_START         = 1
	CONTROL_LAYER_DIFF_END           = 10 // この数を含まない。テスト用も含まない
	CONTROL_LAYER_ALL_SIZE           = 13 // この数を含まない
)

// GetControlLayerName - 利きボードのレイヤーの名前
func GetControlLayerName(layer int) string {
	switch layer {
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
		panic(fmt.Errorf("Unknown layer=%d", layer))
	}
}

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlRook(layer int, sign int8, excludeFrom Square) {
	for _, from := range pPos.RookLocations {
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPos.AddControlDiff(layer, from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlBishop(layer int, sign int8, excludeFrom Square) {
	for _, from := range pPos.BishopLocations {
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPos.AddControlDiff(layer, from, sign)
		}
	}
}

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlLance(layer int, sign int8, excludeFrom Square) {
	for _, from := range pPos.LanceLocations {

		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外
			pPos.AddControlDiff(layer, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiff(layer int, from Square, sign int8) {
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

	for _, to := range sq_list {
		// fmt.Printf("Debug: to=%d\n", to)
		// 差分の方のテーブルを更新（＾～＾）
		pPos.ControlBoards[ph][layer][to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPos *Position) ClearControlDiff() {
	// layer 0 を除く
	for layer := CONTROL_LAYER_DIFF_START; layer < CONTROL_LAYER_DIFF_END; layer += 1 {
		pPos.ClearControlLayer(layer)
	}
}

func (pPos *Position) ClearControlLayer(layer int) {
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			pPos.ControlBoards[0][layer][sq] = 0
			pPos.ControlBoards[1][layer][sq] = 0
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPos *Position) MergeControlDiff() {
	for sq := Square(11); sq < BOARD_SIZE; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			// layer 0 を除く
			for layer := CONTROL_LAYER_DIFF_START; layer < CONTROL_LAYER_DIFF_END; layer += 1 {
				pPos.ControlBoards[0][CONTROL_LAYER_SUM][sq] += pPos.ControlBoards[0][layer][sq]
				pPos.ControlBoards[1][CONTROL_LAYER_SUM][sq] += pPos.ControlBoards[1][layer][sq]
			}
		}
	}
}

// RecalculateControl - 利きの再計算
func (pPos *Position) RecalculateControl(layer1 int) {

	pPos.ClearControlLayer(layer1)

	for from := Square(11); from < BOARD_SIZE; from += 1 {
		if File(from) != 0 && Rank(from) != 0 && !pPos.IsEmptySq(from) {
			piece := pPos.Board[from]
			phase := Who(piece)
			sq_list := GenControl(pPos, from)

			for _, to := range sq_list {
				pPos.ControlBoards[phase-1][layer1][to] += 1
			}

		}
	}
}

// DiffControl - 利きテーブルの差分計算
func (pPos *Position) DiffControl(layer1 int, layer2 int, layer3 int) {

	pPos.ClearControlLayer(layer3)

	for phase := 0; phase < 2; phase += 1 {
		for from := Square(11); from < BOARD_SIZE; from += 1 {
			if File(from) != 0 && Rank(from) != 0 {

				pPos.ControlBoards[phase][layer3][from] = pPos.ControlBoards[phase][layer1][from] - pPos.ControlBoards[phase][layer2][from]

			}
		}
	}
}
