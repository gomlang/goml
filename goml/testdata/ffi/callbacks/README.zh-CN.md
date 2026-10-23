<a id="retained-and-concurrent-go-callbacks"></a>

# 保留回调与并发 Go 回调

[English](README.md) | [简体中文](README.zh-CN.md)

在仓库根目录运行：

```sh
just make
cd goml/testdata/ffi/callbacks
../../../../stage2/bin/goml run
```

此测试用例只使用 Go 标准库包和本地适配层。输出如下：

```text
deferred until release: true
concurrent retained callbacks: true
unregistered callbacks stopped: true
recursive reentry: true
callback can unregister itself: true
registration stress balanced: true
```

`retained` 创建一个 GoML 闭包，并返回持有转换后函数的 Go 订阅对象。创建闭包的函数不再持有该局部闭包。`Start` 在 goroutine 上将一批任务排队，并在任何回调执行前返回。`Release` 打开执行闸门；`Wait` 等待整批任务结束，并报告派发的回调次数。每个已启动的批次都必须释放并等待完成。Wait 只能使用一次；Release 是幂等的。测试在 Start 和 Release 之间强制执行一次 GC，不依赖终结器。

第一个批次从 64 个 goroutine 调用同一个被保留的回调。捕获的 Go `Counter` 使用原子操作，捕获的偏移量不可变。FFI 转换不会让普通 GoML `Ref` 单元变得线程安全；并发修改与其他共享状态一样需要同步。

`Unregister` 清除订阅对象对回调的引用，且是幂等的。工作协程在订阅对象的互斥锁保护下取得回调，释放锁后再调用。这允许回调注销自身，也允许 Go↔GoML 递归重入。工作协程已经选中的回调仍可能在 Unregister 之后开始或完成；释放这些回调使用的资源前，应等待活动批次结束。排队中的批次若在 Unregister 后释放，将不再派发任何回调，新批次也不会派发。

自行注销的回调读取一个在 Release 前初始化的 Ref。通道闸门提供所需的执行顺序；该 Ref 不会被并发修改。测试随后执行 128 轮注册、派发、注销循环，验证每次预期的计数递增，并检查注册计数是否平衡。适配层的全局计数器只存储数量，不存储引用或句柄。订阅对象直接持有函数，清空该字段即释放订阅对象对函数的引用。可达性与显式注册生命周期仍是两个独立问题。

在受支持的主机上，使用 Go 竞态检测器运行生成的程序：

```sh
go run -race _artifact/build/pkg/ffi_callbacks/goml_generated.go
```

驱动程序回归测试将这些已提交的源文件复制到隔离工作区，执行全新构建和缓存构建，检查生成的源文件是否稳定，并在启用 CGO 的情况下执行上述竞态检测命令。测试同时检查输出和进程是否成功，不依赖按实际时间休眠或终结器的执行时机。
