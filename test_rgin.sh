#!/bin/bash
set -e

# 1. 编译 rgin 工具
rm -f rgin
echo "[1/5] Building rgin..."
go build -o rgin main.go

# 2. 查看帮助和版本
./rgin --help || { echo "[FAIL] rgin --help"; exit 1; }
./rgin --version || { echo "[FAIL] rgin --version"; exit 1; }

echo "[2/5] rgin CLI works."

# 3. 生成 Gin 项目
rm -rf testapp
echo "[3/5] Generating Gin project..."
./rgin init testapp
cd testapp

# 自动初始化 SQLite 数据库表结构（需已安装sqlite3）
if command -v sqlite3 >/dev/null 2>&1; then
  mkdir -p data
  sqlite3 ./data/app.db < models/create_table.sql
  echo "[3.5/5] SQLite 表结构已初始化。"
else
  echo "[WARN] 未检测到sqlite3命令，未自动初始化表结构。"
fi

# 4. 编译并运行生成的 Gin 项目
rm -f app
echo "[4/5] Building generated Gin project..."
go mod tidy
go build -o app main.go

# 5. 启动项目（后台运行，测试端口）
echo "[5/5] Running generated Gin project..."
./app &
APP_PID=$!
sleep 2

# 6. 测试端口（假设默认 8080 或 8888，可根据实际端口调整）
PORT=8888
if lsof -i :$PORT | grep LISTEN; then
  echo "[PASS] Gin app is running on port $PORT."
else
  echo "[FAIL] Gin app did not start on port $PORT."
  kill $APP_PID
  exit 1
fi

# 7. 清理
kill $APP_PID
cd ..
rm -rf testapp rgin

echo "[SUCCESS] All rgin tests passed!" 