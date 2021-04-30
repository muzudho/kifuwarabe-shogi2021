// 利きボード
package take12

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
			pPosSys.PControlBoardSystem.AddControlDiff(pPos, c, from, sign)
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
			pPosSys.PControlBoardSystem.AddControlDiff(pPos, c, from, sign)
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
			pPosSys.PControlBoardSystem.AddControlDiff(pPos, c, from, sign)
		}
	}
}
