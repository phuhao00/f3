
在编译 Golang 代码时，您可以使用一些编译器标志和选项来进行优化。以下是一些常见的优化选项：

编译时优化：

-gcflags: 允许您传递标志给 Go 编译器，控制代码生成、优化和垃圾回收等。例如，使用 -gcflags="-l -N" 可以禁用函数内联和优化。
-ldflags: 允许您传递链接器标志，可以设置一些额外的链接信息。例如，-ldflags="-s -w" 可以去除调试信息和符号表，减小可执行文件的大小。
交叉编译：

-buildmode: 用于指定编译器构建模式，例如将代码编译为静态库（-buildmode=c-archive）、动态库（-buildmode=c-shared）或者插件（-buildmode=plugin）。
CPU 架构和平台：

-target: 用于指定目标平台和架构，例如使用 GOOS=linux GOARCH=amd64 go build 可以编译适用于 Linux x86-64 架构的程序。
优化等级：

-gc: 允许您设置垃圾回收器的调优级别，可以通过 -gc=1 或 -gc=off 来调整。
-trimpath: 剔除构建路径信息，以减小可执行文件的大小。



# 使用-gcflags优化，禁用函数内联和优化
go build -gcflags="-l -N" -o output_file main.go

# 设置链接标志，去除调试信息和符号表
go build -ldflags="-s -w" -o output_file main.go

# 交叉编译为 Linux x86-64 架构
GOOS=linux GOARCH=amd64 go build -o output_file main.go
