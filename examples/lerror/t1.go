package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

type MyError struct {
	Inner error
	Message string
	StrackTrace string
	Misc map[string]interface{}
}

type LowLevelErr struct {
	error
}

func wrapError(err error, msg string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:       err,
		Message:     fmt.Sprintf(msg, msgArgs),
		StrackTrace: string(debug.Stack()),
		Misc:        make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

func isGloabllyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}


type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const josBinPath = "/med/log/asfaljal"
	isExecutable,err := isGloabllyExec(josBinPath)
	if err != nil {
		return IntermediateErr{wrapError(err, "cannot run job %q: requisite binaries not available", id)}
	} else if isExecutable == false {
		return wrapError(nil, "job binary is not executable")
	}
	return exec.Command(josBinPath, "--id="+id).Run()
}

func handleError(key int, err error, msg string) {
	log.SetPrefix(fmt.Sprintf("[logId]: %v: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, msg)
}

func main()  {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		msg := "eeeeeeeeer"
		if _,ok := err.(IntermediateErr); ok {
		msg = err.Error()
		handleError(1, err, msg)
		}
	}
}


