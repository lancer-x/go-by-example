package main
/*
在函数调用失败的时候，有三种方式可以将下游的 error 传递出去：

直接返回失败函数返回的 error。
使用 "pkg/errors".Wrap 增加更多的上下文信息，这种情况下可以使用 "pkg/errors".Cause 去提取原始的 error 信息。
如果调用者不需要检测和处理返回的 error 信息的话，可以直接使用 fmt.Errorf 将需要附加的信息进行格式化添加进去。
 */

/*
如果条件允许，最好增加上下文信息。
比如 "connection refused" 和 "call service foo: connection refused" 这两种错误信息在可读性上比较也是高下立判。
当增加上下文信息的时候，尽量保持简洁。比如像 "failed to" 这种极其明显的信息就没有必要写上去了。
 */

