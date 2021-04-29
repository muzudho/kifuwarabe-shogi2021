package take11

import (
	"math/rand"
)

const RESIGN_VALUE = -32768
const MAX_VALUE = 32767

var nodesNum int
var depthEnd int = 1

// Search - 探索部
func Search(pPos *Position) Move {

	nodesNum = 0
	curDepth := 0
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	bestmove, bestVal := search2(pPos, curDepth)

	// 評価値出力（＾～＾）
	G.Chat.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
		curDepth, nodesNum, bestVal, bestmove.ToCode(), bestmove.ToCode())

	// ゲーム向けの軽い乱数
	return bestmove
}

// search2 - 探索部
func search2(pPos *Position, curDepth int) (Move, int16) {
	//fmt.Printf("Search2: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPos)
	move_length := len(move_list)
	//fmt.Printf("%d/%d move_length=%d\n", curDepth, depthEnd, move_length)

	if move_length == 0 {
		return RESIGN_MOVE, RESIGN_VALUE
	}

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var bestMoveList []Move
	// 最初に最低値を入れておけば、更新されるだろ（＾～＾）
	var bestVal int16 = RESIGN_VALUE

	// 相手の評価値
	var opponentWorstVal int16 = MAX_VALUE

	// その手を指してみるぜ（＾～＾）
	for _, move := range move_list {
		pPos.DoMove(move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pPos.CapturedList[pPos.OffsetMovesIndex-1]
		materialVal := EvalMaterial(captured)

		if curDepth < depthEnd {
			// 再帰
			_, opponentVal := search2(pPos, curDepth+1)

			if opponentVal < opponentWorstVal {
				// より低い価値が見つかったら更新
				bestMoveList = nil
				bestMoveList = append(bestMoveList, move)
				opponentWorstVal = opponentVal
			} else if bestVal == materialVal {
				// 最高値が並んだら配列の要素として追加
				bestMoveList = append(bestMoveList, move)
			}

		} else {
			// 葉ノードでは、相手の手ではなく、自分の局面に点数を付けます

			if bestVal < materialVal {
				// より高い価値が見つかったら更新
				bestMoveList = nil
				bestMoveList = append(bestMoveList, move)
				bestVal = materialVal
			} else if bestVal == materialVal {
				// 最高値が並んだら配列の要素として追加
				bestMoveList = append(bestMoveList, move)
			}
		}

		pPos.UndoMove()
	}

	if curDepth < depthEnd {
		// 葉以外のノードでは、相手の評価値の逆が、自分の評価値
		bestVal = -opponentWorstVal
	}

	var bestmove = RESIGN_MOVE
	bestmove_length := len(bestMoveList)
	//fmt.Printf("%d/%d bestmove_length=%d\n", curDepth, depthEnd, bestmove_length)
	if bestmove_length > 0 {
		bestmove = bestMoveList[rand.Intn(bestmove_length)]
	}

	// 評価値出力（＾～＾）
	// G.Chat.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestmove.ToCode(), bestmove.ToCode())

	// 0件にはならないはず（＾～＾）
	return bestmove, bestVal
}
