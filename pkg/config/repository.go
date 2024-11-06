package config

type SQLDBSettings struct {
	SqlDsn              string
	SqlMaxOpenConns     int
	SqlMaxIdleConns     int
	SqlConnsMaxLifetime int
}

// データベース接続文字列
func (s *SQLDBSettings) DSN() string {
	return s.SqlDsn
}

// 最大オープン接続数
func (s *SQLDBSettings) MaxOpenConns() int {
	return s.SqlMaxOpenConns
}

// アイドル接続数
func (s *SQLDBSettings) MaxIdleConns() int {
	return s.SqlMaxIdleConns
}

// 接続の最大有効期間
func (s *SQLDBSettings) ConnsMaxLifetime() int {
	return s.SqlConnsMaxLifetime
}

/*
データベース接続設定を格納する構造体を定義し、その設定値にアクセスするためのメソッドを提供している
*/
