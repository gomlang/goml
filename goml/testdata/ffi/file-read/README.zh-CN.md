<a id="file-reading-across-the-go-boundary"></a>

# 跨越 Go 边界读取文件

[English](README.md) | [简体中文](README.zh-CN.md)

在仓库根目录构建工具链并运行此测试用例：

```sh
just make
cd goml/testdata/ffi/file-read
../../../../stage2/bin/goml run
```

此测试用例只使用 Go 标准库和本地 `shim` 包。请从本目录运行，以便 `data.txt` 解析到随附的三字节文件。使用已安装的工具链时，在此目录运行 `goml run` 即可。

```text
read: 3
partial data: abc
Go round-trip preserved partial read: true
read error: unexpected EOF
explicit close succeeded: true
original handle is closed: true
```

`os.Open` 接收显式转换后的原始 Go 字符串。返回的 `*os.File` 在传入 Go 并传回后保持不变。`io.ReadFull` 尝试从三字节文件读取四字节，因此同时返回三个有效字节和 `io.ErrUnexpectedEOF`。`ffi::Outcome` 保留这两个结果。Go 适配层接收传回的结果，并检查原始错误的身份和字节数。

Go 将数据写入 GoML 缓冲区的共享切片。测试只复制返回的字节范围，然后检查 UTF-8。无效文本会被报告，而不会替换字节；`Bytes` 值仍包含原始数据。如果将读取结果转换为出错时丢弃值的 Result，就会丢失这次部分读取的数据。

显式 `Close` 紧接读取发生，在解释读取结果之前完成。关闭 `EchoFile` 返回的指针也会关闭原始句柄，最后一次读取验证了这一点。关闭错误与读取错误分别保留。适配器不会创建句柄注册表，也不会安排自动 Close。测试不依赖 Go 终结器运行。

这只是有界读取示例，不是完整的文件加载 API。较长文件会成功填满四字节缓冲区；空文件会报告 EOF 并返回零字节。文件不存在时，程序报告打开错误并返回，不会尝试读取或关闭 nil 指针。`IsPartialRead` 适配函数有意只识别随附的三字节场景。

驱动程序回归测试在隔离工作区中运行这些已提交的源文件，涵盖缓冲区填满、空文件、无效 UTF-8 和缺失文件的情况。文本无效时仍保留部分读取的字节数和原始读取错误，且句柄在解码前已关闭。发布冒烟测试只复制源文件，使用解压后的工具链运行此用例两次，检查全新构建、缓存构建及生成源文件的摘要，不访问第三方网络。
