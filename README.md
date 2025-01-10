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

requirements: Docker desktop

1. `make db-up`: DBの起動
2. `make app-build`: Dockerイメージのビルド
3. `make app-run`: Dockerコンテナの起動

- `make app-down`: apiの終了
- `make db-down`: DBの終了
- `make db-migrate`: DB初回起動後に実行(要goose)
