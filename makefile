# Makefile for generating .pb.go files from .proto files

# 定义Proto文件的路径和文件名
PROTO_SRC_DIR := path/to/proto
PROTO_FILE := your_proto_file.proto

# 定义生成的Go代码的路径
GO_OUT_DIR := path/to/output

# protoc命令路径，如果没有配置在PATH中，则需要指定完整路径
PROTOC := protoc

# 生成.pb.go文件的目标
.PHONY: proto
proto:
    $(PROTOC) --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative $(PROTO_SRC_DIR)/$(PROTO_FILE)
