#!/bin/bash

# 指定您的 Golang 源代码目录
SOURCE_DIR="./your_go_source_code"

# 指定要编译的 Go 文件
GO_FILE="main.go"

# 指定输出的可执行文件名
OUTPUT_FILE="your_executable_name"

# 进入源代码目录
cd "$SOURCE_DIR" || exit

# 执行编译命令
go build -o "$OUTPUT_FILE" "$GO_FILE"

if [ $? -eq 0 ]; then
    echo "Golang program compiled successfully."
else
    echo "Error: Failed to compile Golang program."
    exit 1
fi

