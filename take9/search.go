package take8

import (
	"math/rand"
)

// Search - 探索部
func Search(pPos *Position) Move {

	// 指し手生成
	// 探索中に削除される指し手を除く
	move_list := GenMoveList(pPos)
	size := len(move_list)

	if size == 0 {
		return ResignMove
	}

	// その手を指してみるぜ（＾～＾）
	for _, move := range move_list {
		pPos.DoMove(move)

		// captured := pPos.CapturedList[pPos.OffsetMovesIndex]

		pPos.UndoMove()
	}

	// ゲーム向けの軽い乱数
	return move_list[rand.Intn(size)]
}
