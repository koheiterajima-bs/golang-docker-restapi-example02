package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server *serverConfig
	DB     *databaseConfig
}

type serverConfig struct {
	Port int `env:"PORT, default=8080"` // サーバーのポート番号を設定
}

type databaseConfig struct {
	DSN              string `env:"MYSQL_DSN,default=root:password@tcp(localhost:33061)/example?charset=utf8&parseTime=true"` // データベース接続文字列を指定
	MaxOpenConns     int    `env:"MAX_OPEN_CONNS,default=100"`                                                               // データベース接続の最大接続数
	MaxIdleConns     int    `env:"MAX_IDLE_CONNS,default=100"`                                                               // データベースのアイドル接続数の上限を設定するフィールド
	ConnsMaxLifetime int    `env:"CONNS_MAX_LIFETIME,default=100"`                                                           // データベース接続の最大有効期間を設定するフィールド
}

// 設定を読み込み、Config構造体を生成して返す関数(環境変数を基にアプリケーション設定を初期化する)
func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	// 環境変数から設定を読み込み、cfg構造体に格納する
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// サーバーのアドレスをstring形式で返すメソッド
func (cfg *Config) Address() string {
	return fmt.Sprintf(":%d", cfg.Server.Port) // ポート番号を表示
}

/*
Go言語で環境変数を扱う
- 方法
  - osパッケージを使う
  - ライブラリを使う(今回はEnvconfig)

- Envconfig
  - 環境変数または任意の検索関数に基づいて構造体フィールドの値を設定する
*/
