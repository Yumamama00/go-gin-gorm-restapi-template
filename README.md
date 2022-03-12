# go-gin-gorm-restapi-template

Go (Gin + Gorm) を使用した RestAPI 開発のスターターテンプレート
RestAPI Starter Template Using Golang (Gin + Gorm)

## 主要な依存ライブラリ

- Gin (Web フレームワーク)
  - REST API 対応、リクエスト/レスポンスのフィルターを行う Middleware 等。機能豊富なリッチ Web フレームワーク。
  - https://github.com/gin-gonic/gin
- Gorm (ORM)
  - Go の ORM のデファクトスタンダード的存在。機能豊富な ORM。
  - https://github.com/go-gorm/gorm
- Zap (ロギング)
  - 高速なロギングライブラリ。構造化ロギングにも対応。
  - https://github.com/uber-go/zap
- dig (Dependency Injection)
  - DI コンテナライブラリ
  - https://github.com/uber-go/dig
- godotenv (環境変数)
  - 環境変数ファイルを用いるためのライブラリ
  - https://github.com/joho/godotenv

## パッケージ構成/アーキテクチャ (DDD + クリーンアーキテクチャ + CQRS をイメージ)

```
ー go.mod
ー go.sum
ー .test_env (テスト用 環境変数ファイル)
ー .dev_env (開発用 環境変数ファイル)
ー .prod_env (本番用 環境変数ファイル)
ー main.go

ー interfaces (プレゼンテーション層)
　　ー api
　　　ー controller (リクエスト、レスポンスのハンドラー)
　　　ー middleware (リクエストFilter)
　　　ー httputils (http系のユーティリティ)
　　　ー routers  (Routing処理)

ー usecases (アプリケーション層。アプリケーションサービスのInterface、実装が入る)

ー query (Query専用 → ※サンプルでは作成していない)
　　ー services (Queryで用いるサービスのInterface)
　　ー response (DTO, ViewModel)

ー domain (ドメイン層)
　　ー model
     ー xxxx(例:User)
　　　　⇨　entity、value object、repository (Interface)
　　ー services (ドメインサービスの実装)

ー infrastructure (インフラストラクチャ層)
　　ー persistence (repositryの実装)
　　ー query (QueryService の実装)

ー test (テスト専用 → ※サンプルでは作成していない)
　　ー migrate (GORMのMigrate処理)
　　ー repository (Repository層の単体テスト)
　　ー usecases (Usecase層の単体テスト、モックのRepositoryを利用)

ー utils (ユーティリティ)
　　ー di (依存性注入関数を登録)
　　ー idutil (id生成)
　　ー logger (Zap Logger)
　　ー transaction (トランザクションのHelper。実装はインフラストラクチャ層へ)

ー tools (DBの自動生成等)
　ー migrate (GORMのMigrate処理)
```
