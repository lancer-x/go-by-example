package main

import "log"

func main1()  {
	log.Println("haha")

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("new haha")
}
