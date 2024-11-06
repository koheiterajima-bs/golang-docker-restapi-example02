package handler

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"go.uber.org/zap"
// )

// // データベースからユーザー情報を取得し、結果をJSON形式で返すための関数
// func (h *Handler) GetUserHandler() http.HandleFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// URLからuserIDを取得
// 		userID := chi.URLParam(r, "userID")
// 		// データベースからユーザーを取得する処理
// 		user, err := h.repo.UserRepository.GetUser(context.Background(), userID)
// 		if err != nil {
// 			h.logger.Error("failed to get user", zap.Error(err))
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// JSON形式に変換
// 		b, err := json.Marshal(user)
// 		if err != nil {
// 			h.logger.Error("marshal error", zap.Error(err))
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		// HTTPレスポンスの作成
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(b)
// 	}
// }

// // 複数のユーザー情報を取得し、クライアントにJSON形式で返す
// func (h *Handler) ListUsersHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// HTTPリクエストのコンテキストを取得
// 		c := r.Context()
// 		// ユーザー情報の取得
// 		users, err := h.repo.UserRepository.ListUsers(c)
// 		if err != nil {
// 			h.logger.Error("failed to get user", zap.Error(err))
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// JSON形式に変換
// 		b, err := json.Marshal(users)
// 		if err != nil {
// 			h.logger.Error("marshal error", zap.Error(err))
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		// HTTPレスポンスの作成
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(b)
// 	}
// }

// // 新しいユーザーを作成するためのHTTPリクエストを処理する
// func (h *Handler) PostUserHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		// JSONデータのデコード(リクエストボディから送信されたJSONデータをGoのinput.User構造に変換している)
// 		var user input.User
// 		decode := json.NewDecoder(req.Body)
// 		err := decode.Decode(&user)
// 		if err != nil {
// 			log.Printf("failed to decode (error:%s)", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		// ユーザーの作成
// 		id, err := h.repo.UserRepository.CreateUser(context.Background(), user.Name)
// 		if err != nil {
// 			log.Printf("failed to create user (error:%s)", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		// ユーザー重複？の処理
// 		if *id == 0 {
// 			log.Printf("user is conflict")
// 			w.WriteHeader(http.StatusConflict)
// 			return
// 		}

// 		// ユーザーデータをJSON形式に変換
// 		byte, err := json.Marshal(user)
// 		if err != nil {
// 			log.Printf("failed to marshal user (error:%s)", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		// HTTPレスポンスの作成
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(byte)
// 	}
// }

// // HTTPリクエストに応じて指定されたuserIDに対応するユーザーをデータベースから削除するハンドラを実装
// func (h *Handler) DeleteUserHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		// ユーザーIDの取得
// 		userID := chi.URLParam(req, "userID")
// 		// ユーザーの削除
// 		err := h.repo.UserRepository.DeleteUser(context.Background(), userID)
// 		if err != nil {
// 			log.Printf("failed to create user (error:%s)", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusNoContent)
// 	}
// }
