# take10

玉は逃げてくれてるようだし、アルファベータ探索入れてみよかな（＾～＾）  
その前に駒得評価関数入れてみるかな（＾～＾）？  
駒得評価を入れてみたけど、きふわらべがどう評価しているか見えないぜ（＾～＾）
評価値出してくれだぜ（＾～＾）  

## Test

```plain
# 強制終了した（＾～＾） -> 配列のサイズが変わっていたので length 再計算（＾～＾）
position startpos moves 7g7f 3a4b 6i7h 6c6d 7i6h 5a6b 8h5e 4b5a 5e6d 2c2d 6h7g 6b6c 6d4f 8b3b 4f2d 7a8b 3i4h 1c1d 5i6h 5a6b 7g8h 4a5a 8g8f 2b3a 8h8g 6a5b 7h7i 3b2b 2d4f
go btime 75000 wtime 74000 binc 2000 winc 2000

# 飛打されてなぜか即投了（＾～＾）->利きボードの内容がおかしい
position startpos moves 7g7f 3c3d 2g2f 2b8h 7i8h 1c1d B*4e B*7h 6i7h 8c8d 4e6c+ 3d3e 6c5c 8b6b 5c6b 5a6b R*6d
go btime 69000 wtime 68000 binc 2000 winc 2000
```
