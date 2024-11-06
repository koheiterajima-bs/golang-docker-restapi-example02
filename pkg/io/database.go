package io

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	errs "github.com/pkg/errors"
)

// この時間内に応答がない場合、エラーが発生する
const (
	queryTimeoutSec = 3
)

// MySQLデータベース接続の設定を取得するためのインターフェース
type MySQLSettings interface {
	DSN() string
	MaxOpenConns() int
	MaxIdleConns() int
	ConnsMaxLifetime() int
}

// データベース接続を管理するための構造体
type SQLDatabase struct {
	Database *sql.DB
}

// データベースの初期化
func NewDatabase(setting MySQLSettings) (*SQLDatabase, error) {
	db, err := sql.Open("mysql", setting.DSN())
	if err != nil {
		return nil, errs.WithStack(err)
	}

	// check config
	// 最大オープン接続数、最大アイドル接続数、接続の最大寿命の設定
	if setting.MaxOpenConns() <= 0 {
		return nil, errs.WithStack(errs.New("require set max open conns"))
	}
	if setting.MaxIdleConns() <= 0 {
		return nil, errs.WithStack(errs.New("require set max idle conns"))
	}
	if setting.ConnsMaxLifetime() <= 0 {
		return nil, errs.WithStack(errs.New("require set conns max lifetime"))
	}
	db.SetMaxOpenConns(setting.MaxOpenConns())
	db.SetMaxIdleConns(setting.MaxIdleConns())
	db.SetConnMaxLifetime(time.Duration(setting.ConnsMaxLifetime()) * time.Second)

	return &SQLDatabase{Database: db}, nil
}

// データベースへのPing(データベースに接続できるかを確認するためのメソッド)
func (d *SQLDatabase) Ping() error {
	return d.Database.Ping()
}

// トランザクションの開始
func (d *SQLDatabase) Begin() (*sql.Tx, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	tx, err := d.Database.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	return tx, cancel, err
}

// データベースのクローズ
func (d *SQLDatabase) Close() error {
	return d.Database.Close()
}

// ステートメントの準備(クエリの実行前にデータベースが存在しているかを確認)
func (d *SQLDatabase) Prepare(query string) (*sql.Stmt, error) {
	if d.Database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeoutSec*time.Second)
	defer cancel()
	stmt, err := d.Database.PrepareContext(ctx, query)

	return stmt, err
}

// クエリの実行
func (d *SQLDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.Database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeoutSec*time.Second)
	defer cancel()
	res, err := d.Database.ExecContext(ctx, query, args...)

	return res, err
}

// エラー処理
func errDoesNotDB() error {
	return errs.New("database does not exist. Please Open() first")
}
