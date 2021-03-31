package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	// Epoch is set to the twitter snowflake epoch of 2019-10-16 00:00:00 UTC+8 in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1571155200000

	// NodeBits holds the number of bits to use for Node
	// Remember, you have a total 22 bits to share between Node/Step
	NodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Node/Step
	StepBits uint8 = 12

	// DEPRECATED: the below four variables will be removed in a future release.
	mu        sync.Mutex
	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)


func main() {
    newNodeByIp()
}

func newNodeByIp() {
    var nodeID int64
	// generate 64/ms,  64000 /s per node
	initMask(16, 6)

	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println("Ip >>>>>>=", ipnet.IP.String())
					ip4 := ipnet.IP.To4()
                    fmt.Println(ip4[1])
					nodeID += int64(ip4[2]) << 8
					nodeID += int64(ip4[3])
					break
				}
			}
		}
	}

	fmt.Println("nodeID >>>>>>=", nodeID)
	//return NewNode(nodeID)
}


func initMask(nodeBits, stepBits uint8) {
	nodeMax = -1 ^ (-1 << nodeBits)
	nodeMask = nodeMax << stepBits
	stepMask = -1 ^ (-1 << stepBits)
	timeShift = nodeBits + stepBits
	nodeShift = stepBits
}