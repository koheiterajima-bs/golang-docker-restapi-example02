package infrastructure

// import (
// 	"context"
// 	"database/sql"

// 	"github.com/pkg/errors"
// )

// type UserRepository struct {
// 	database *io.SQLDatabase
// }

// func NewUserRepository(db *io.SQLDatabase) *UserRepository {
// 	return &UserRepository{
// 		database: db,
// 	}
// }

// // userIDに対応するユーザーの情報をデータベースから取得
// func (r *UserRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
// 	// SQLクエリ(クエリ内の?はプレースホルダで後でuserIDがここにバインドされる)
// 	query := `
// 		SELECT
// 			id,
// 			name
// 		FROM
// 			users
// 		WHERE
// 			id = ?
// 	`

// 	// クエリの準備(SQLクエリをコンパイルし、実行できるステートメントを返す)
// 	stmtOut, err := r.database.Prepare(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmtOut.Close()

// 	// クエリ実行と結果の取得
// 	var user entity.User
// 	// コンテキストとuserIDを渡して、クエリを実行し、1行の結果を取得する
// 	err = stmtOut.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Name)
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			return nil, nil
// 		default:
// 			return nil, err
// 		}
// 	}

// 	// 結果の返却
// 	return &user, nil
// }

// // データベース内のすべてのユーザー情報を取得
// func (r *UserRepository) ListUsers(ctx context.Context) ([]*entity.User, error) {
// 	// SQLクエリの定義
// 	query := `
// 		SELECT
// 			id,
// 			name
// 		FROM
// 			users
// 	`

// 	// クエリの準備
// 	stmtOut, err := r.database.Prepare(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// クエリの実行
// 	rows, err := stmtOut.QueryContext(ctx)
// 	if err != nil {
// 		return []*entity.User{}, err
// 	}
// 	defer stmtOut.Close()

// 	// 結果の格納
// 	users := make([]*entity.User, 0)
// 	// 行の取得とスキャン
// 	for rows.Next() {
// 		var user entity.User
// 		err = rows.Scan(&user.ID, &user.Name)
// 		if err != nil {
// 			return []*entity.User{}, err
// 		}
// 		users = append(users, &user)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return []*entity.User{}, err
// 	}

// 	return users, nil
// }

// // 新しいユーザーをデータベースに挿入するための処理
// func (r *UserRepository) CreateUser(ctx context.Context, name string) (*int, error) {
// 	// SQLクエリの定義
// 	query := `
// 	INSERT INTO
// 		users (name)
// 	VALUE (?)
// 	`

// 	// クエリの準備
// 	stmtOut, err := r.database.Database.Prepare(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmtOut.Close()

// 	// クエリの実行
// 	result, err := stmtOut.ExecContext(ctx, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 挿入した行のIDを取得
// 	insertID, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// IDの型変換と返却
// 	id := int(insertID)

// 	return &id, nil
// }

// // 指定されたユーザーの名前を更新するための処理
// func (r *UserRepository) UpdateUser(ctx context.Context, userID string, name string) error {
// 	// SQLクエリの定義
// 	query := `
// 		UPDATE
// 			users
// 		SET
// 			name = ?
// 		WHERE
// 			id = ?
// 	`

// 	// クエリの準備
// 	stmtOut, err := r.database.Database.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmtOut.Close()

// 	// クエリの実行
// 	result, err := stmtOut.ExecContext(ctx, name, userID)
// 	if err != nil {
// 		return err
// 	}
// 	// 影響を受けた行数の取得
// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // 指定されたユーザーIDに基づいてデータベースからユーザーを削除する
// func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
// 	// SQLクエリの定義
// 	query := `
// 		DELETE FROM
// 			users
// 		WHERE
// 			id = ?
// 	`

// 	// クエリの準備
// 	stmtOut, err := r.database.Database.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmtOut.Close()

// 	// クエリの実行
// 	result, err := stmtOut.ExecContext(ctx, userID)
// 	if err != nil {
// 		return err
// 	}
// 	// 影響を受けた行数の取得
// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // 指定されたユーザーIDがデータベースに存在するかどうかを確認する
// func (r *UserRepository) TxExistUser(ctx context.Context, tx *sql.Tx, userID int) (bool, error) {
// 	// SQLクエリの定義
// 	query := `
// 		SELECT EXISTS (
// 			SELECT
// 				*
// 			FROM
// 				users
// 			WHERE
// 				id = ?
// 		);
// 	`
// 	// クエリの準備
// 	// トランザクションに関連づけられたクエリを事前にコンパイルする
// 	stmtOut, err := tx.PrepareContext(ctx, query)
// 	if err != nil {
// 		return false, errors.WithStack(err)
// 	}
// 	defer stmtOut.Close()

// 	// クエリの実行
// 	var b bool
// 	// 準備したステートメントを実行し、1行の結果を取得
// 	err = stmtOut.QueryRowContext(ctx, userID).Scan(&b)
// 	if err != nil {
// 		return false, errors.WithStack(err)
// 	}

// 	return b, nil
// }
