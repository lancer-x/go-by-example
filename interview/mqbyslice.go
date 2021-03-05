package main

import (
	"fmt"
	"sync"
)

type lmq struct {
	q []int
	sync.Mutex
}

func newLmq() *lmq {
	return &lmq{q: make([]int, 0)}
}

func (q *lmq) push(n int) {
	q.Lock()
	defer q.Unlock()
	q.q = append(q.q, n)
	fmt.Println(q.q)
}

func (q *lmq) pop() (ret int, err error) {
	q.Lock()
	defer q.Unlock()
	if len(q.q) > 0 {
		ret = q.q[0]
		q.q = q.q[1:]
		return ret, nil
	}
	return 0, fmt.Errorf("empty q")
}

func main1() {
	myQ := newLmq()
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go testPush(myQ, i, &wg)
	}
	wg.Wait()
	testPop(myQ)
}

func testPush(q *lmq, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := n * 10; i < (n+1)*10; i++ {
		q.push(i)
	}
}

func testPop(q *lmq) {
	for len(q.q) > 0 {
		num, _ := q.pop()
		fmt.Println(num)
	}

}
