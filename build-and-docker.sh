#!/bin/sh
# 編譯 Go 專案並重新上傳到 Docker
set -e

echo "[1] 編譯 Go 專案..."     
go build ./...
go build -o server main.go

echo "[2] 建置 Docker 映像檔..."
docker build -t go-http-service .

echo "[3] 移除舊容器..."
docker rm -f go-http-service || true

echo "[4] 啟動新容器..."
docker run -d -p 8080:8080 --network mynet   --name go-http-service go-http-service

echo "完成！服務已啟動於 http://localhost:8080"
