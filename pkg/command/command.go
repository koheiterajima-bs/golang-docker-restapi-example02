package command

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/koheiterajima-bs/golang-docker-restapi-example02/pkg/config"
	"github.com/koheiterajima-bs/golang-docker-restapi-example02/pkg/io"
	"go.uber.org/zap"
	"honnef.co/go/tools/lintcmd/version"
)

// プログラムの実行を終了時の状態
const (
	exitOK  = 0 // 正常終了
	exitErr = 1 // 異常終了
)

func Run() {
	// 親のコンテキストがない場合に使われる空のコンテキスト
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	// Logger
	// Loggerの初期化
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err) // os.Stderrは標準エラー出力であり、ターミナルに直接表示される
		return exitErr
	}
	// ログ出力のバッファをフラッシュして全てのログが確実に出力されるようにする役割がある
	defer logger.Sync()
	// ログの各エントリーに共通の情報を追加できる、バージョンがログごとに記録されるので、異なるバージョン間でログを管理する際に役立つ
	logger = logger.With(zap.String("version", version.Version))

	// Config
	// 設定の初期化を呼び出し
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return exitErr
	}

	// init listener
	// net.Listenでサーバー作成
	listener, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		logger.Error("failed to listen port", zap.Int("port", cfg.Server.Port), zap.Error(err))
		return exitErr
	}
	logger.Info("server start listening", zap.Int("port", cfg.Server.Port))

	// サーバー全体を安全に停止させるためのキャンセル機能
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// listenerが使われていないことの一時的処置
	fmt.Println(listener)

	// DB
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}
	// データベースの初期化
	db, err := io.NewDatabase(sqlSetting)
	if err != nil {
		logger.Error("failed to connect db", zap.Error(err))
		return exitErr
	} else {
		logger.Info("successed to connect db")
	}
	// データベースへのPing(データベースに接続できるかを確認するためのメソッド)
	if err = db.Ping(); err != nil {
		logger.Error("failed to ping mysql db", zap.Error(err))
		return exitErr
	}

	return exitOK
}
