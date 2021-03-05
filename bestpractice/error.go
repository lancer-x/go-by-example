package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
在 Go 语言中声明 error 可以有多种方式：

errors.New 声明包含简单静态字符串的 error
fmt.Errorf 格式化 error string
其他自定义类型使用了 Error() 方法
使用 "pkg/errors".Wrap
当要把 error 作为返回值的时候，可以考虑如下的处理方式

是不是不需要额外信息，如果是，errors.New 就足够了。
client 需要检测和处理返回的 error 吗？如果是，最好使用实现了 Error() 方法的自定义类型，这样可以包含更多的信息。
error 是不是从下游函数传递过来的？如果是，考虑一下 error wrap，参考：section on error wrapping.
其他情况，fmt.Errorf 一般足够了。
 */

func main5()  {

}


/*
1、对于 client 需要检测和处理 error 的情况，这里详细说一下。如果你要通过 errors.New 声明一个简单的 error，那么可以使用一个变量声明：var ErrCouldNotOpen = errors.New("Could not open")
 */
//bad use
func open1() error {
	return errors.New("could not open")
}
func badUse1()  {
	if err := open1(); err != nil {
		if err.Error() == "could not open" {
			//handle
		} else {
			panic("unknown error")
		}
	}
}

//good use
var ErrCouldNotOpen = errors.New("could not open")

func open2() error {
	return ErrCouldNotOpen
}

func goodUse1()  {
	if err := open2(); err !=nil {
		if err == ErrCouldNotOpen {
			//handle
		} else {
			panic("unkown error")
		}
	}
}

/*
2、如果需要 error 中包含更多的信息，而不仅仅类型原生 error 的这种简单字符串，那么最好使用一个自定义类型。
 */
//bad
func open3() error {
	return fmt.Errorf("file not found")
}
func badUse2()  {
	if err := open3(); err != nil {
		if strings.Contains(err.Error(), "not found") {
			//
		} else {
			panic("aa")
		}
	}
}

//good
type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open4(file string) error {
	return errNotFound{file: file}
}
func goodUse2()  {
	if err := open4("aa"); err != nil {
		if _,ok := err.(errNotFound); ok {
			//
		} else {
			panic("aa")
		}
	}
}


/*
3、在直接暴露自定义的 error 类型的时候，最好 export 配套的检测自定义 error 类型的函数。
在2基础上新增检测函数
 */
func IsNotFoundError(err error) bool {
	_,ok := err.(errNotFound)
	return ok
}

func goodUse3()  {
	if err := open4("aaa"); err != nil {
		if IsNotFoundError(err) {
			//
		} else {
			panic("bb")
		}
	}
}

