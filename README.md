# errors

基于 `github.com/pkg/errors` 包，增加对 `error code` 的支持，完全兼容 `github.com/pkg/errors`。

性能跟 `github.com/pkg/errors` 基本持平。

该 errors 包匹配的错误码设计请参考：[marmotedu/sample-code](https://github.com/marmotedu/sample-code/blob/master/README.md)

错误包需具备的功能：
- 支持错误堆栈：极大提高问题定位的效率，降低定位的难度。
- 支持不同的打印格式：比如 %+v、%v、%s 等格式，可以根据需要打印不同丰富度的错误信息。
- 支持 Wrap/Unwrap 功能：在已有的错误上，追加新的信息。
  - 不同的错误类型，Wrap 函数的逻辑也可以不同。
  - 在调用 Wrap 时也会生成一个错误堆栈节点。
- 支持 Is 函数：由于存在 wrapping error，无法直接通过 == 或 != 判断错误。
- 支持 As 函数：对于多层嵌套的 error，无法直接使用类型断言判断。
- 支持两种错误创建方式：非格式化创建（errors.New）和格式化创建（errors.Errorf）。