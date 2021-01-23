package main

import (
	"flag"
	"fmt"
	"time"
)

func main()  {
	period := flag.Duration("p", 1 * time.Second, "sleep period")
	flag.Parse()

	fmt.Println("sleeping for ", *period)
	time.Sleep(*period)

}
