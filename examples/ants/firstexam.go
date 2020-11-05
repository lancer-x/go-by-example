package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{})  {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Println("run with%d\n", n)
}

func demoFunc()  {
	time.Sleep(10 * time.Second)
	fmt.Println("Hello World!")
}



func main()  {
	defer ants.Release()

	runTimes := 1000

	var wg sync.WaitGroup
	syncTask := func() {
		demoFunc()
		wg.Done()
	}


	for i:=0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncTask)
	}
	wg.Wait()

	p,_ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	for i:=0; i<runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
}