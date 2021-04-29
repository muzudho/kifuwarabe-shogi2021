# take11

駒得評価関数は入れたしな（＾～＾）  
きふわらべがどう評価しているか、infoも入れたし（＾～＾）  
利きの差分表も 飛、角、香 を分けた（＾～＾）  
テスト用のランダム局面を増やすために、シャッフルとプレイアウトも入れた（＾～＾）  
王手回避のときに打をして、間駒もするはずだぜ（＾～＾）  

玉は逃げてくれてるようだし、アルファベータ探索入れてみよかな（＾～＾）  
探索中に盤面が元に戻ってないことがあるなあ（＾～＾）盤面２つに増やすかな（＾～＾）？  

## Test

```plain
# 強制終了した（＾～＾）
position startpos moves 1g1f 4a3b 1f1e 6a5b 2g2f 2c2d 3g3f 5a4b 3f3e 4c4d 4g4f 5b4c 5g5f 7c7d 5f5e 7d7e 6g6f 9c9d 6f6e 9d9e 8g8f 8c8d 2h4h 5c5d 1e1d 1c1d 3e3d 3c3d 5e5d 4c5d 6e6d 5d6d 4h6h 8d8e 2f2e 2d2e 4f4e 8e8f 4e4d 8f8g+ 4d4c 3b4c 7g7f 8g8h 7f7e 8h9i 7e7d 8b8i+ 7d7c N*5g 7c7b 2b7g+
go btime 86000 wtime 83000 binc 2000 winc 2000
Error: 59 square is empty
quit

# 盤をコピー（＾～＾）
board copy

# 盤[1]を表示（＾～＾）
pos 1
```
