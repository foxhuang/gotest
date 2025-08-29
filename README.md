# Go HTTP GraphQL Service

本專案依據 [golang-standards/project-layout](https://github.com/golang-standards/project-layout) 標準架構，實作 Go HTTP + GraphQL 服務。

## 目錄結構

```
cmd/
  server/main.go         # 服務入口
internal/
  service/resolver.go    # GraphQL resolver
  gql/                   # GraphQL schema 與自動產生檔案
    schema.graphqls
    reader.graphqls
    muser.graphqls
    generated.go
    models_gen.go
pkg/                     # 可放共用函式庫
build-and-docker.sh      # 編譯與 Docker 啟動腳本
gqlgen.yml               # gqlgen 設定檔
Dockerfile               # Docker 設定
README.md                # 專案說明
```

## 啟動方式

1. 安裝依賴
   ```powershell
   go mod tidy
   ```
2. 產生 GraphQL 型別
   ```powershell
   go run github.com/99designs/gqlgen generate
   ```
3. 啟動服務
   ```powershell
   go run cmd/server/main.go
   ```

## Docker 部署

1. 編譯與啟動容器
   ```bash
   ./build-and-docker.sh
   ```
2. 或直接建置映像檔
   ```bash
   docker build -t go-service:latest .
   docker run -d -p 8080:8080 --name go-service go-service:latest
   ```

## GraphQL Playground

- 服務啟動後可於 [http://localhost:8080/graphiql](http://localhost:8080/graphiql) 測試 API

## gqlgen 使用

- schema 檔案分散於 `internal/gql/` 目錄
- 修改 schema 後，請重新執行 `gqlgen generate` 產生 Go 型別
- 主要 resolver 實作於 `internal/service/resolver.go`

---

如有問題請提出 issue 或聯絡維護者。
