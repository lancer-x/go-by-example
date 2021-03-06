package main

import (
	"bufio"
	"os"
)

func main()  {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			//fmt.Println("%d\t%s\n", n, line)
			print(line)
		}
	}
}