<a id="unicode-data"></a>

# Unicode 数据

[English](DATA.md) | [简体中文](DATA.zh-CN.md)

所有公开表均固定使用 Unicode 15.0.0。七个 tables_*.goml 文件由 tools/generate_unicode_tables.go 根据 Go 1.26 的 unicode 包生成。生成器会拒绝其他 Unicode 版本，对具名表排序，并通过 stage2/bin/gomlfmt（可由 GOMLFMT 覆盖）输出规范格式的 GoML。

在仓库根目录运行，确保 PATH 中可找到 Go 1.26：

```sh
go run tools/generate_unicode_tables.go --check
go run tools/generate_unicode_tables.go --write
```

检查模式比较生成并格式化后的输出，不修改源文件，并在脚本 CI 检查中运行。写入模式重新生成七个派生文件。现有的完整 casefold.goml 改用 tools/generate_unicode_casefold.py 生成；它要求 Python 的 Unicode 数据版本为 15.0.0，并支持多标量折叠。两个生成器都支持不修改文件的 `--check`、规范格式化和 GOMLFMT 覆盖；两项检查均在脚本 CI 中运行。使用 `python3 tools/generate_unicode_casefold.py --write` 重新生成完整折叠，或使用 `--check` 验证。不带参数调用此 Python 脚本时，保留其历史上的写入模式。

范围记录使用小写十六进制：含端点的下界和上界各六位，后跟八位非零 u32 步长，共 20 个字符。大小写记录包含一个码点及其大写、小写和标题大小写映射，每项六位，共 24 位。简单折叠记录包含一个码点及其后继，共 12 位。记录有序排列；查询使用二分查找，不调用原生函数。该编码属于私有实现，不是受支持的序列化格式。

Unicode 包不依赖其他标准库包。其私有输出辅助函数将已验证的字符串片段或十六进制数字组装到内建字节向量中，再使用现有的 UTF-8 字符串边界原语完成转换。这使 text -> unicode 保持无环，无需移动公开的 Unicode 类型或复制带版本的表。完整大小写折叠生成器使用相同的辅助函数；其映射数据保持不变。

外部标准算法使用方会将每个有效标量的分类和简单大小写映射与 Go 对比，并检查所有具名范围表。Go 调用仅发生在生成和测试阶段。上游 Go 的版权和许可声明见 ../THIRD_PARTY_NOTICES。
