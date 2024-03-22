package pms

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 17:11
 * @Desc:
 */

//go:generate kitex.exe -module github.com/yuanyp8/wxcommerce/pms -type protobuf  -I ./idl  ./idl/product.proto
//go:generate kitex.exe -module github.com/yuanyp8/wxcommerce/pms -type protobuf  -I ./idl  ./idl/user.proto

// 进入对应的rpc目录生成文件
// cd cmd/xxx
//go:generate kitex.exe -module onetech/example/easy_note -service easy_note.note -use onetech/example/easy_note/kitex_gen -type protobuf  -I ../../idl  ../../idl/note.proto
//go:generate kitex.exe -module onetech/example/easy_note -service easy_note.user -use onetech/example/easy_note/kitex_gen -type protobuf  -I ../../idl  ../../idl/user.proto

// 生成uml类图
//go:generate protodot -src ./idl/note.proto -output ./doc/note.svg

// 运行之前要更新下： go get -u github.com/golang/protobuf
// go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0
