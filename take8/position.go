package take8

import (
	"fmt"
	"strconv"
	"strings"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 00～99
const BOARD_SIZE = 100

// 1:先手 2:後手
type Phase byte

// マス番号 00～99,100～113
type Square uint32

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

const (
	// 空マス
	ZEROTH = Phase(0)
	// 先手
	FIRST = Phase(1)
	// 後手
	SECOND = Phase(2)
)

// 駒
const (
	PIECE_EMPTY = ""
	PIECE_K1    = "K"
	PIECE_R1    = "R"
	PIECE_B1    = "B"
	PIECE_G1    = "G"
	PIECE_S1    = "S"
	PIECE_N1    = "N"
	PIECE_L1    = "L"
	PIECE_P1    = "P"
	PIECE_PR1   = "+R"
	PIECE_PB1   = "+B"
	PIECE_PS1   = "+S"
	PIECE_PN1   = "+N"
	PIECE_PL1   = "+L"
	PIECE_PP1   = "+P"
	PIECE_K2    = "k"
	PIECE_R2    = "r"
	PIECE_B2    = "b"
	PIECE_G2    = "g"
	PIECE_S2    = "s"
	PIECE_N2    = "n"
	PIECE_L2    = "l"
	PIECE_P2    = "p"
	PIECE_PR2   = "+r"
	PIECE_PB2   = "+b"
	PIECE_PS2   = "+s"
	PIECE_PN2   = "+n"
	PIECE_PL2   = "+l"
	PIECE_PP2   = "+p"
)

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]string
	// 飛車の場所。長い利きを消すために必要（＾～＾）
	RookLocations [2]Square
	// 角の場所。長い利きを消すために必要（＾～＾）
	BishopLocations [2]Square
	// 香の場所。長い利きを消すために必要（＾～＾）
	LanceLocations [4]Square
	// 利きテーブル [0]先手 [1]後手
	// マスへの利き数が入っています
	ControlBoards [2][BOARD_SIZE]int8
	// マスへの利き数の差分が入っています
	ControlBoardsDiff [2][BOARD_SIZE]int8

	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	Phase Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]string
}

func NewPosition() *Position {
	var ins = new(Position)
	ins.ResetToStartpos()
	return ins
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPos *Position) ResetToZero() {
	// 筋、段のラベルだけ入れとくぜ（＾～＾）
	pPos.Board = [BOARD_SIZE]string{
		"", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"1", "", "", "", "", "", "", "", "", "",
		"2", "", "", "", "", "", "", "", "", "",
		"3", "", "", "", "", "", "", "", "", "",
		"4", "", "", "", "", "", "", "", "", "",
		"5", "", "", "", "", "", "", "", "", "",
		"6", "", "", "", "", "", "", "", "", "",
		"7", "", "", "", "", "", "", "", "", "",
		"8", "", "", "", "", "", "", "", "", "",
		"9", "", "", "", "", "", "", "", "", "",
	}
	pPos.ControlBoards = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}}
	pPos.ControlBoardsDiff = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}}
	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.RookLocations = [2]Square{0, 0}
	pPos.BishopLocations = [2]Square{0, 0}
	pPos.LanceLocations = [4]Square{0, 0, 0, 0}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.Phase = FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPos.CapturedList = [MOVES_SIZE]string{}
}

// ResetToStartpos - 初期局面にします。
func (pPos *Position) ResetToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]string{
		"", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"1", "l", "", "p", "", "", "", "P", "", "L",
		"2", "n", "b", "p", "", "", "", "P", "R", "N",
		"3", "s", "", "p", "", "", "", "P", "", "S",
		"4", "g", "", "p", "", "", "", "P", "", "G",
		"5", "k", "", "p", "", "", "", "P", "", "K",
		"6", "g", "", "p", "", "", "", "P", "", "G",
		"7", "s", "", "p", "", "", "", "P", "", "S",
		"8", "n", "r", "p", "", "", "", "P", "B", "N",
		"9", "l", "", "p", "", "", "", "P", "", "L",
	}
	pPos.ControlBoards = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 2, 2, 1,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1, 3, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 2,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 2,
		0, 0, 0, 0, 0, 0, 1, 2, 3, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 2, 0,
		0, 0, 0, 0, 0, 0, 1, 3, 1, 1,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 1, 3, 1, 0, 0, 0, 0, 0,
		0, 0, 2, 0, 1, 0, 0, 0, 0, 0,
		0, 2, 3, 2, 1, 0, 0, 0, 0, 0,
		0, 1, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 2, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 1, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 1, 3, 1, 1, 0, 0, 0, 0, 0,
		0, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 2, 2, 1, 0, 0, 0, 0, 0,
	}}
	pPos.ControlBoardsDiff = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}}
	pPos.RookLocations = [2]Square{28, 82}
	pPos.BishopLocations = [2]Square{22, 88}
	pPos.LanceLocations = [4]Square{11, 19, 91, 99}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.Phase = FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPos.CapturedList = [MOVES_SIZE]string{}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pPos *Position) ReadPosition(command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.ResetToStartpos()
		i = 17

		if i >= len || command[i] != ' ' {
			return
		}
		i += 1
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.ResetToZero()
		i = 14
		var rank = 1
		var file = 9

	BoardLoop:
		for {
			promoted := false
			switch pc := command[i]; pc {
			case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
				pPos.Board[file*10+rank] = string(pc)
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = ""
					file -= 1
				}
				i += 1
			case '+':
				i += 1
				promoted = true
			case '/':
				file = 9
				rank += 1
				i += 1
			case ' ':
				i += 1
				break BoardLoop
			default:
				panic("Undefined sfen board")
			}

			if promoted {
				switch pc2 := command[i]; pc2 {
				case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
					pPos.Board[file*10+rank] = "+" + string(pc2)
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i, sq := range pPos.RookLocations {
					if sq == SQUARE_EMPTY {
						pPos.RookLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			case 'B', 'b':
				for i, sq := range pPos.BishopLocations {
					if sq == SQUARE_EMPTY {
						pPos.BishopLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			case 'L', 'l':
				for i, sq := range pPos.LanceLocations {
					if sq == SQUARE_EMPTY {
						pPos.LanceLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pPos.Phase = FIRST
			i += 1
		case 'w':
			pPos.Phase = SECOND
			i += 1
		default:
			panic("Fatal: Unknown phase")
		}

		if command[i] != ' ' {
			// 手番の後ろにスペースがない（＾～＾）
			panic("Fatal: Nothing space")
		}
		i += 1

		// 持ち駒
		if command[i] == '-' {
			i += 1
			if command[i] != ' ' {
				// 持ち駒 - の後ろにスペースがない（＾～＾）
				panic("Fatal: Nothing space after -")
			}
			i += 1
		} else {
		HandLoop:
			for {
				var drop_index Square
				var piece = command[i]
				switch piece {
				case 'R':
					drop_index = DROP_R1
				case 'B':
					drop_index = DROP_B1
				case 'G':
					drop_index = DROP_G1
				case 'S':
					drop_index = DROP_S1
				case 'N':
					drop_index = DROP_N1
				case 'L':
					drop_index = DROP_L1
				case 'P':
					drop_index = DROP_P1
				case 'r':
					drop_index = DROP_R2
				case 'b':
					drop_index = DROP_B2
				case 'g':
					drop_index = DROP_G2
				case 's':
					drop_index = DROP_S2
				case 'n':
					drop_index = DROP_N2
				case 'l':
					drop_index = DROP_L2
				case 'p':
					drop_index = DROP_P2
				case ' ':
					i += 1
					break HandLoop
				default:
					panic(fmt.Errorf("Fatal: Unknown piece=%c", piece))
				}
				i += 1

				var number = 0
			NumberLoop:
				for {
					switch figure := command[i]; figure {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						num, err := strconv.Atoi(string(figure))
						if err != nil {
							panic(err)
						}
						i += 1
						number *= 10
						number += num
					case ' ':
						i += 1
						break HandLoop
					default:
						break NumberLoop
					}
				}

				pPos.Hands[drop_index] = number

				// 長い利きの駒は位置を覚えておくぜ（＾～＾）
				switch drop_index {
				case DROP_R1, DROP_R2:
					for i, sq := range pPos.RookLocations {
						if sq == SQUARE_EMPTY {
							pPos.RookLocations[i] = drop_index
							break
						}
					}
				case DROP_B1, DROP_B2:
					for i, sq := range pPos.BishopLocations {
						if sq == SQUARE_EMPTY {
							pPos.BishopLocations[i] = drop_index
							break
						}
					}
				case DROP_L1, DROP_L2:
					for i, sq := range pPos.LanceLocations {
						if sq == SQUARE_EMPTY {
							pPos.LanceLocations[i] = drop_index
							break
						}
					}
				}

			}
		}

		// 手数
		pPos.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pPos.StartMovesNum *= 10
				pPos.StartMovesNum += num
			case ' ':
				i += 1
				break MovesNumLoop
			default:
				break MovesNumLoop
			}
		}

	} else {
		fmt.Printf("Error: Unknown command=[%s]", command)
	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	start_phase := pPos.Phase
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = ParseMove(command, &i, pPos.Phase)
			if err != nil {
				fmt.Println(err)
				fmt.Println(pPos.Sprint())
				panic(err)
			}
			pPos.Moves[pPos.OffsetMovesIndex] = move
			pPos.OffsetMovesIndex += 1
			pPos.Phase = pPos.Phase%2 + 1
		}
	}

	// 開始局面の利きを計算（＾～＾）
	fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				fmt.Printf("Debug: sq=%d\n", sq)
				pPos.AddControl(sq, 1)
			}
		}
	}
	fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPos.OffsetMovesIndex = 0
	pPos.Phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPos.DoMove(pPos.Moves[i])
	}
}

// ParseMove
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var move = NewMoveValue()

	var hand1 = Square(0)

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		hand1 = DROP_R1
	case 'B':
		*i += 1
		hand1 = DROP_B1
	case 'G':
		*i += 1
		hand1 = DROP_G1
	case 'S':
		*i += 1
		hand1 = DROP_S1
	case 'N':
		*i += 1
		hand1 = DROP_N1
	case 'L':
		*i += 1
		hand1 = DROP_L1
	case 'P':
		*i += 1
		hand1 = DROP_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand1 != 0 {
		switch phase {
		case FIRST:
			move = move.ReplaceSource(hand1)
		case SECOND:
			move = move.ReplaceSource(hand1 + DROP_TYPE_SIZE)
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("Fatal: not *")
		}
		*i += 1
		count = 1
	}

	// file, rank
	for count < 2 {
		switch ch := command[*i]; ch {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*i += 1
			file, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}

			var rank int
			switch ch2 := command[*i]; ch2 {
			case 'a':
				rank = 1
			case 'b':
				rank = 2
			case 'c':
				rank = 3
			case 'd':
				rank = 4
			case 'e':
				rank = 5
			case 'f':
				rank = 6
			case 'g':
				rank = 7
			case 'h':
				rank = 8
			case 'i':
				rank = 9
			default:
				return *new(Move), fmt.Errorf("Fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				move = move.ReplaceSource(sq)
			} else if count == 1 {
				move = move.ReplaceDestination(sq)
			} else {
				return *new(Move), fmt.Errorf("Fatal: Unknown count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		move = move.ReplacePromotion(true)
	}

	return move, nil
}

// Print - 局面出力（＾ｑ＾）
func (pPos *Position) Sprint() string {
	var phase_str = "?"
	if pPos.Phase == FIRST {
		phase_str = "First"
	} else if pPos.Phase == SECOND {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPos.StartMovesNum, (pPos.StartMovesNum+pPos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[7], pPos.Hands[8], pPos.Hands[9], pPos.Hands[10], pPos.Hands[11], pPos.Hands[12], pPos.Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pPos.Board[90], pPos.Board[80], pPos.Board[70], pPos.Board[60], pPos.Board[50], pPos.Board[40], pPos.Board[30], pPos.Board[20], pPos.Board[10], pPos.Board[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[91], pPos.Board[81], pPos.Board[71], pPos.Board[61], pPos.Board[51], pPos.Board[41], pPos.Board[31], pPos.Board[21], pPos.Board[11], pPos.Board[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[92], pPos.Board[82], pPos.Board[72], pPos.Board[62], pPos.Board[52], pPos.Board[42], pPos.Board[32], pPos.Board[22], pPos.Board[12], pPos.Board[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[93], pPos.Board[83], pPos.Board[73], pPos.Board[63], pPos.Board[53], pPos.Board[43], pPos.Board[33], pPos.Board[23], pPos.Board[13], pPos.Board[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[94], pPos.Board[84], pPos.Board[74], pPos.Board[64], pPos.Board[54], pPos.Board[44], pPos.Board[34], pPos.Board[24], pPos.Board[14], pPos.Board[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[95], pPos.Board[85], pPos.Board[75], pPos.Board[65], pPos.Board[55], pPos.Board[45], pPos.Board[35], pPos.Board[25], pPos.Board[15], pPos.Board[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[96], pPos.Board[86], pPos.Board[76], pPos.Board[66], pPos.Board[56], pPos.Board[46], pPos.Board[36], pPos.Board[26], pPos.Board[16], pPos.Board[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[97], pPos.Board[87], pPos.Board[77], pPos.Board[67], pPos.Board[57], pPos.Board[47], pPos.Board[37], pPos.Board[27], pPos.Board[17], pPos.Board[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[98], pPos.Board[88], pPos.Board[78], pPos.Board[68], pPos.Board[58], pPos.Board[48], pPos.Board[38], pPos.Board[28], pPos.Board[18], pPos.Board[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[99], pPos.Board[89], pPos.Board[79], pPos.Board[69], pPos.Board[59], pPos.Board[49], pPos.Board[39], pPos.Board[29], pPos.Board[19], pPos.Board[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[0], pPos.Hands[1], pPos.Hands[2], pPos.Hands[3], pPos.Hands[4], pPos.Hands[5], pPos.Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_text := make([]byte, 0, MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pPos.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pPos.Moves[i].ToCode()...)
	}

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `flag` - 1:利き数の差分, それ以外: 利き数ボード
func (pPos *Position) SprintControl(phase Phase, flag int) string {
	var board [BOARD_SIZE]int8
	var phase_str string
	switch phase {
	case FIRST:
		phase_str = "First"
		if flag == 1 { // 利き数の差分
			board = pPos.ControlBoardsDiff[0]
		} else {
			board = pPos.ControlBoards[0]
		}
	case SECOND:
		phase_str = "Second"
		if flag == 1 {
			board = pPos.ControlBoardsDiff[1]
		} else {
			board = pPos.ControlBoards[1]
		}
	default:
		return "\n"
	}

	return "\n" +
		//
		fmt.Sprintf("[Control %s]\n", phase_str) +
		//
		"\n" +
		//
		"  9  8  7  6  5  4  3  2  1\n" +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| a\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| b\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| c\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| e\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| f\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| g\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| h\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| i\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
}

// SprintLocation - あの駒どこにいんの？を表示
func (pPos *Position) SprintLocation() string {
	return "\n" +
		//
		" R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			pPos.RookLocations[0], pPos.RookLocations[1],
			pPos.BishopLocations[0], pPos.BishopLocations[1],
			pPos.LanceLocations[0], pPos.LanceLocations[1],
			pPos.LanceLocations[2], pPos.LanceLocations[3]) +
		//
		"+---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}

// DoMove - 一手指すぜ（＾～＾）
func (pPos *Position) DoMove(move Move) {
	// 作業前に、長い利きの駒の利きを -1 します
	pPos.AddControlAllSlidingPiece(-1)

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	mov_src_sq := move.GetSource()
	mov_dst_sq := move.GetDestination()
	var cap_src_sq Square

	// まず、打かどうかで処理を分けます
	drop := mov_src_sq
	var piece string
	switch mov_src_sq {
	case DROP_R1:
		piece = PIECE_R1
	case DROP_B1:
		piece = PIECE_B1
	case DROP_G1:
		piece = PIECE_G1
	case DROP_S1:
		piece = PIECE_S1
	case DROP_N1:
		piece = PIECE_N1
	case DROP_L1:
		piece = PIECE_L1
	case DROP_P1:
		piece = PIECE_P1
	case DROP_R2:
		piece = PIECE_R2
	case DROP_B2:
		piece = PIECE_B2
	case DROP_G2:
		piece = PIECE_G2
	case DROP_S2:
		piece = PIECE_S2
	case DROP_N2:
		piece = PIECE_N2
	case DROP_L2:
		piece = PIECE_L2
	case DROP_P2:
		piece = PIECE_P2
	default:
		// Not drop
		drop = Square(0)
	}

	if drop != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands[drop-DROP_ORIGIN] -= 1

		// 行き先に駒を置きます
		pPos.Board[mov_dst_sq] = piece
		pPos.AddControl(mov_dst_sq, 1)
		mov_piece_type = What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します
		captured := pPos.Board[mov_dst_sq]
		if captured != PIECE_EMPTY {
			pPos.AddControl(mov_dst_sq, -1)
			cap_piece_type = What(captured)
			cap_src_sq = mov_dst_sq
		}

		// 元位置の駒を除去
		pPos.AddControl(mov_src_sq, -1)

		// 行き先の駒の配置
		pPos.Board[mov_dst_sq] = pPos.Board[mov_src_sq]
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.Board[mov_src_sq] = PIECE_EMPTY
		pPos.AddControl(mov_dst_sq, 1)

		drop := Square(0)
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1, PIECE_PR1:
			drop = DROP_R2
		case PIECE_B1, PIECE_PB1:
			drop = DROP_B2
		case PIECE_G1:
			drop = DROP_G2
		case PIECE_S1, PIECE_PS1:
			drop = DROP_S2
		case PIECE_N1, PIECE_PN1:
			drop = DROP_N2
		case PIECE_L1, PIECE_PL1:
			drop = DROP_L2
		case PIECE_P1, PIECE_PP1:
			drop = DROP_P2
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2, PIECE_PR2:
			drop = DROP_R1
		case PIECE_B2, PIECE_PB2:
			drop = DROP_B1
		case PIECE_G2:
			drop = DROP_G1
		case PIECE_S2, PIECE_PS2:
			drop = DROP_S1
		case PIECE_N2, PIECE_PN2:
			drop = DROP_N1
		case PIECE_L2, PIECE_PL2:
			drop = DROP_L1
		case PIECE_P2, PIECE_PP2:
			drop = DROP_P1
		default:
			fmt.Printf("Error: Unknown captured=[%s]", captured)
		}

		if drop != 0 {
			pPos.Hands[drop-DROP_ORIGIN] += 1
		}
	}

	pPos.Moves[pPos.OffsetMovesIndex] = move
	pPos.OffsetMovesIndex += 1
	pPos.Phase = pPos.Phase%2 + 1

	// 長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []interface{}{mov_piece_type, cap_piece_type}
	src_sq_list := []interface{}{mov_src_sq, cap_src_sq}
	for j, moving_piece_type := range piece_type_list {
		switch moving_piece_type {
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i, sq := range pPos.RookLocations {
				if sq == src_sq_list[j] {
					pPos.RookLocations[i] = sq
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i, sq := range pPos.BishopLocations {
				if sq == src_sq_list[j] {
					pPos.BishopLocations[i] = sq
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.LanceLocations {
				if sq == src_sq_list[j] {
					pPos.LanceLocations[i] = sq
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します
	// TODO pPos.AddControlAllSlidingPiece(1)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPos *Position) UndoMove() {
	if pPos.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	// 作業前に、長い利きの駒の利きを -1 します
	// TODO pPos.AddControlAllSlidingPiece(-1)

	pPos.OffsetMovesIndex -= 1
	pPos.Phase = pPos.Phase%2 + 1
	move := pPos.Moves[pPos.OffsetMovesIndex]
	captured := pPos.CapturedList[pPos.OffsetMovesIndex]

	mov_dst_sq := move.GetDestination()
	mov_src_sq := move.GetSource()
	var cap_dst_sq Square

	// 打かどうかで分けます
	switch mov_src_sq {
	case DROP_R1, DROP_B1, DROP_G1, DROP_S1, DROP_N1, DROP_L1, DROP_P1, DROP_R2, DROP_B2, DROP_G2, DROP_S2, DROP_N2, DROP_L2, DROP_P2:
		// 打なら
		drop := mov_src_sq
		// 盤上から駒を除去します
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.Board[mov_dst_sq] = PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands[drop-DROP_ORIGIN] += 1
		cap_dst_sq = 0
	default:
		// 打でないなら

		// 行き先の駒の除去
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.AddControl(mov_dst_sq, -1)
		// 移動元への駒の配置
		pPos.Board[mov_src_sq] = pPos.Board[mov_dst_sq]

		// あれば、取った駒は駒台から下ろします
		cap := Square(0)
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1, PIECE_PR1:
			cap = DROP_R2
		case PIECE_B1, PIECE_PB1:
			cap = DROP_B2
		case PIECE_G1:
			cap = DROP_G2
		case PIECE_S1, PIECE_PS1:
			cap = DROP_S2
		case PIECE_N1, PIECE_PN1:
			cap = DROP_N2
		case PIECE_L1, PIECE_PL1:
			cap = DROP_L2
		case PIECE_P1, PIECE_PP1:
			cap = DROP_P2
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2, PIECE_PR2:
			cap = DROP_R1
		case PIECE_B2, PIECE_PB2:
			cap = DROP_B1
		case PIECE_G2:
			cap = DROP_G1
		case PIECE_S2, PIECE_PS2:
			cap = DROP_S1
		case PIECE_N2, PIECE_PN2:
			cap = DROP_N1
		case PIECE_L2, PIECE_PL2:
			cap = DROP_L1
		case PIECE_P2, PIECE_PP2:
			cap = DROP_P1
		default:
			fmt.Printf("Error: Unknown captured=[%s]", captured)
		}

		if cap != 0 {
			cap_dst_sq = cap
			pPos.Hands[cap-DROP_ORIGIN] -= 1

			// 取った駒を行き先に戻します
			cap_piece_type = What(captured)
			pPos.Board[mov_dst_sq] = captured
			pPos.AddControl(mov_src_sq, 1)
			pPos.AddControl(mov_dst_sq, 1)
		}
	}

	// 長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []interface{}{mov_piece_type, cap_piece_type}
	dst_sq_list := []interface{}{mov_dst_sq, cap_dst_sq}
	for j, moving_piece_type := range piece_type_list {
		switch moving_piece_type {
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i, sq := range pPos.RookLocations {
				if sq == dst_sq_list[j] {
					pPos.RookLocations[i] = sq
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i, sq := range pPos.BishopLocations {
				if sq == dst_sq_list[j] {
					pPos.BishopLocations[i] = sq
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.LanceLocations {
				if sq == dst_sq_list[j] {
					pPos.LanceLocations[i] = sq
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します
	// TODO pPos.AddControlAllSlidingPiece(1)
}

// AddControlAllSlidingPiece - すべての長い利きの駒の利きを増減させます
func (pPos *Position) AddControlAllSlidingPiece(sign int8) {
	for _, from := range pPos.RookLocations {
		if from != SQUARE_EMPTY {
			pPos.AddControl(from, sign)
		}
	}
	for _, from := range pPos.BishopLocations {
		if from != SQUARE_EMPTY {
			pPos.AddControl(from, sign)
		}
	}
	for _, from := range pPos.LanceLocations {
		if from != SQUARE_EMPTY {
			pPos.AddControl(from, sign)
		}
	}
}

// AddControl - 盤上のマスを指定することで、そこにある駒の利きを増減させます
func (pPos *Position) AddControl(from Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(Who(piece)) - 1
	fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenControl(pPos, from)

	for _, to := range sq_list {
		fmt.Printf("Debug: to=%d\n", to)
		// 差分の方のテーブルを更新（＾～＾）
		pPos.ControlBoardsDiff[ph][to] += sign * 1
	}
}

// Homo - 手番と移動元の駒を持つプレイヤーが等しければ真。移動先が空なら偽
func (pPos *Position) Homo(to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase == Who(pPos.Board[to])
}

// Hetero - 手番と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase != Who(pPos.Board[to])
}

func (pPos *Position) IsEmptySq(sq Square) bool {
	return pPos.Board[sq] == PIECE_EMPTY
}
