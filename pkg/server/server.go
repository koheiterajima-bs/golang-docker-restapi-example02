/*
このファイルでやっていること
- サーバーの初期化
- サーバーのルーティングやエンドポイントに対応するハンドラの設定
*/

package server

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"go.uber.org/zap"
// 	"honnef.co/go/tools/config"
// )

// type Server struct {
// 	Router  *chi.Mux         // ルーティング管理
// 	server  *http.Server     // HTTPサーバー自体
// 	handler *handler.Handler // ハンドラ管理
// 	log     *zap.Logger      // ログ管理
// }

// // サーバーインスタンスを作成するファクトリ関数
// // Server構造体のインスタンスが作成され、初期化が行われる
// func NewServer(registry *handler.Handler, cfg *Config, env *config.Config) *Server {
// 	// サーバー構造体の初期化
// 	s := &Server{
// 		Router:  chi.NewServer(),
// 		handler: registry,
// 	}

// 	// ログの設定
// 	if cfg != nil {
// 		if log := cfg.Log; log != nil {
// 			s.log = log
// 		}
// 	}

// 	// サーバーのルーティングやエンドポイントに対応するハンドラを登録
// 	s.registerHandler(env, cfg)
// 	return s
// }

// func (s *Server) registerHandler(env *config.Config, cnf *Config) {
// 	// HTTPリクエストのログを自動で出力する役割
// 	s.Router.Use(chiMiddleware.Logger)
// 	// APIのルーティング(どのURLに対してどのハンドラを実行するか)の設定
// 	s.Router.Route("/", func(r chi.Router) {
// 		// user
// 		r.Route("/user", func(r chi.Router) {
// 			// GetUserHandler()とDeleteUserHandler()は、同じパスだが、HTTPメソッド(GET,DELETE)が異なるため、異なるリクエストとして扱われる
// 			r.Get("/{userID}", s.handler.GetUserHandler())
// 			r.Get("/all", s.handler.ListUserHandler())
// 			r.Post("/", s.handler.PostUserHandler())
// 			r.Delete("/{userID}", s.handler.DeleteUserHandler())
// 		})
// 	})
// }
