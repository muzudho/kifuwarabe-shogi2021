package take3

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	DROP_R1 = iota + 100
	DROP_B1
	DROP_G1
	DROP_S1
	DROP_N1
	DROP_L1
	DROP_P1
	DROP_R2
	DROP_B2
	DROP_G2
	DROP_S2
	DROP_N2
	DROP_L2
	DROP_P2
	DROP_ORIGIN = DROP_R1
)

// Move - 指し手
type Move struct {
	// [0]移動元 [1]移動先
	// 持ち駒は仕方ないから 100～113 を使おうぜ（＾～＾）
	Squares []byte
	// 成
	Promotion bool
}

func NewMove() *Move {
	move := new(Move)
	move.Squares = []byte{0, 0}
	return move
}

// ToCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move *Move) ToCode() string {
	str := make([]byte, 0, 5)
	count := 0

	switch move.Squares[0] {
	case DROP_R1, DROP_R2:
		str = append(str, 'R')
		count = 1
	case DROP_B1, DROP_B2:
		str = append(str, 'B')
		count = 1
	case DROP_G1, DROP_G2:
		str = append(str, 'G')
		count = 1
	case DROP_S1, DROP_S2:
		str = append(str, 'S')
		count = 1
	case DROP_N1, DROP_N2:
		str = append(str, 'N')
		count = 1
	case DROP_L1, DROP_L2:
		str = append(str, 'L')
		count = 1
	case DROP_P1, DROP_P2:
		str = append(str, 'P')
		count = 1
	default:
		// Ignored
	}

	if count == 1 {
		str = append(str, '+')
	}

	for count < 2 {
		// 正常時は必ず２桁（＾～＾）
		file := move.Squares[count] / 10
		rank := move.Squares[count] % 10
		// ASCII Code
		// '0'=48, '9'=57, 'a'=97, 'i'=105
		str = append(str, file+48)
		str = append(str, rank+96)
		// fmt.Printf("Debug: file=%d rank=%d\n", file, rank)
		count += 1
	}

	return string(str)
}
