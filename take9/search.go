package take9

// Search - 探索部
func Search(pPos *Position) Move {

	// 指し手生成
	// 探索中に削除される指し手を除く
	move_list := GenMoveList(pPos)
	size := len(move_list)

	if size == 0 {
		return ResignMove
	}

	var bestMove Move
	var bestVal int16
	// その手を指してみるぜ（＾～＾）
	for _, move := range move_list {
		pPos.DoMove(move)

		captured := pPos.CapturedList[pPos.OffsetMovesIndex]
		materialVal := EvalMaterial(captured)

		if bestVal < materialVal {
			bestMove = move
			bestVal = materialVal
		}

		pPos.UndoMove()
	}

	// ゲーム向けの軽い乱数
	return bestMove
}
