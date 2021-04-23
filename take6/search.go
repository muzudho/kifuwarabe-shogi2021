package take6

import (
	"math/rand"
)

// Search - 探索部
func Search() Move {

	// 指し手生成
	legal_move_list := GenMoveList()
	size := len(legal_move_list)

	if size == 0 {
		return ResignMove
	}

	// ゲーム向けの軽い乱数
	return legal_move_list[rand.Intn(size)]
}
