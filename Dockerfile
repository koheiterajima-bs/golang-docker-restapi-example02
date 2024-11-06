# Goアプリケーションのビルド
FROM golang:1.23.0-buster AS build

# ビルド時に渡せる変数
ARG GITHUB_TOKEN=local
ARG VERSION=local

# .netrcファイルにGitHubのログイン情報を保存し、プライベートリポジトリにアクセスできるようにする
RUN echo "machine github.com login ${GITHUB_TOKEN}" > ~/.netrc

# コマンドを実行する作業ディレクトリ
WORKDIR /project

# 現在のディレクトリの全てのファイルを/projectにコピー
COPY . .

# go.modとgo.sumに指定された依存関係をダウンロード
RUN go mod download

# cmd/api/ディレクトリ内のGoサーバーアプリケーションをビルドし、実行ファイルを./bin/serverに生成する
RUN go build -o ./bin/server ./cmd/api/

# 実行用の軽量なイメージ作成
FROM debian:buster

# buildステージからビルドした実行ファイル(/project/bin/server)をコピーし、最終イメージの/bin/serverに配置する
COPY --from=build /project/bin/server /bin/server

# コンテナが実行されるときに/bin/serverが起動される
CMD ["/bin/server"]