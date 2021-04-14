package take1

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	l "github.com/muzudho/go-logger"
	g "github.com/muzudho/kifuwarabe-shogi2021/take1/global"
)

// MainLoop - 開始。
func MainLoop() {
	// Working directory
	dwd, err := os.Getwd()
	if err != nil {
		// ここでは、ログはまだ設定できてない
		panic(fmt.Sprintf("...Engine DefaultWorkingDirectory=%s", dwd))
	}

	// コマンドライン引数登録
	workdir := flag.String("workdir", dwd, "Working directory path.")

	// グローバル変数の作成
	g.G = *new(g.Variables)

	tracePath := filepath.Join(*workdir, "output/trace.log")
	debugPath := filepath.Join(*workdir, "output/debug.log")
	infoPath := filepath.Join(*workdir, "output/info.log")
	noticePath := filepath.Join(*workdir, "output/notice.log")
	warnPath := filepath.Join(*workdir, "output/warn.log")
	errorPath := filepath.Join(*workdir, "output/error.log")
	fatalPath := filepath.Join(*workdir, "output/fatal.log")
	printPath := filepath.Join(*workdir, "output/print.log")

	// ロガーの作成。
	// TODO ディレクトリが存在しなければ、強制終了します。
	g.G.Log = *l.NewLogger(
		tracePath,
		debugPath,
		infoPath,
		noticePath,
		warnPath,
		errorPath,
		fatalPath,
		printPath)

	// 既存のログ・ファイルを削除
	g.G.Log.RemoveAllOldLogs()

	// ログ・ファイルの開閉
	err = g.G.Log.OpenAllLogs()
	if err != nil {
		// ログ・ファイルを開くのに失敗したのだから、ログ・ファイルへは書き込めません
		panic(err)
	}
	defer g.G.Log.CloseAllLogs()

	g.G.Log.Trace("Start Take1\n")

	// 何か標準入力しろだぜ☆（＾～＾）
	scanner := bufio.NewScanner(os.Stdin)

MainLoop:
	for scanner.Scan() {
		g.G.Log.FlushAllLogs()

		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
		case "isready":
		case "usinewgame":
		case "position":
		case "go":
		case "quit":
			break MainLoop
		}
	}

	g.G.Log.Trace("Finished\n")
	g.G.Log.FlushAllLogs()
}
