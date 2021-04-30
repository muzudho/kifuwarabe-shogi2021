package take13

// WaterColor - 水で薄めたような評価値にします
// pCB3 = 0
// pCB4 = 0
// pCB5 = 0
// pCB1 - pCB2 = pCB3
// pCB3 - pCB4 = pCB5
func WaterColor(pCB1 *ControlBoard, pCB2 *ControlBoard, pCB3 *ControlBoard, pCB4 *ControlBoard, pCB5 *ControlBoard) {
	// 将棋盤の内側をスキャンします。

	pCB3.Clear()
	pCB4.Clear()
	pCB5.Clear()

	pW := pCB1
	pX := pCB2
	pY := pCB3
	waterColor2(pW, pX, pY)

	pW = pCB3
	pX = pCB4
	pY = pCB5
	waterColor2(pW, pX, pY)
}

// 81マス・スキャン
func waterColor2(pCB1 *ControlBoard, pCB2 *ControlBoard, pCB3 *ControlBoard) {
	for rank := 1; rank < 10; rank += 1 {
		for file := 9; file > 0; file -= 1 {
			sum := waterColor3(rank, file, pCB1, pCB2)
			sq := SquareFrom(Square(file), Square(rank))
			pCB3.Board[sq] = sum
		}
	}
}

// チェビシェフ距離で 2マス離れたところ、16マス・スキャン
func waterColor3(rank int, file int, pCB1 *ControlBoard, pCB2 *ControlBoard) int8 {

	var sum int8 = 0
	var squares int8 = 0

	// 上辺
	relRank := -2
	for relFile := 2; relFile > -3; relFile -= 1 {
		sum, squares = waterColor4(sum, squares, rank+relRank, file+relFile, pCB1, pCB2)
	}

	// 下辺
	relRank = 2
	for relFile := 2; relFile > -3; relFile -= 1 {
		sum, squares = waterColor4(sum, squares, rank+relRank, file+relFile, pCB1, pCB2)
	}

	// 右辺
	relFile := -2
	for relRank = -1; relRank < 2; relRank += 1 {
		sum, squares = waterColor4(sum, squares, rank+relRank, file+relFile, pCB1, pCB2)
	}

	// 左辺
	relFile = 2
	for relRank = -1; relRank < 2; relRank += 1 {
		sum, squares = waterColor4(sum, squares, rank+relRank, file+relFile, pCB1, pCB2)
	}

	sum /= squares

	return sum
}

func waterColor4(sum int8, squares int8, rank int, file int, pCB1 *ControlBoard, pCB2 *ControlBoard) (int8, int8) {
	// ブラシの面積分の利きを総和します

	sq := SquareFrom(Square(file), Square(rank))
	if OnBoard(sq) {
		sum += pCB1.Board[sq] - pCB2.Board[sq]
		squares += 1
	}

	return sum, squares
}
