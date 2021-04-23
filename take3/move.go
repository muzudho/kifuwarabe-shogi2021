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
)

// Move - 指し手
type Move struct {
	// 移動元
	// 持ち駒は仕方ないから 100～113 を使おうぜ（＾～＾）
	Source byte
	// 移動先
	Destination byte
	// 成
	Promotion bool
}
