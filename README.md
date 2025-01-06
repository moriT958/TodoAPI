# Just Do It

A simple todo api.

## endpoints

- `GET /todos`: todoリストを取得する.
- `PATCH /todos/{id}`: idで指定したtodoを完了状態にする.
- `DELETE /todos/{id}`: idで指定したtodoを削除する.
- `POST /todos`: 新しいtodoを登録する.

## directory

- `/migrations`: gooseで使用するDBのマイグレーションファイル
- `/server`: Handler関数やSever構造体、レスポンスのスキーマなどを記述
- `/todo`: Todoエンティティの定義とそのリポジトリのインターフェースの定義を記述
- `/store`: `/todo`で定義したリポジトリの実装. PostgreSQL用のクエリもこの中に記述.

## How to start

1. `docker compose up -d`: PostgreSQLサーバを起動.
2. `export DATABASE_URL=...`: DBのDSNを指定.
3. `goose up`: マイグレーションの実行
4. `go run main.go`: 実行
