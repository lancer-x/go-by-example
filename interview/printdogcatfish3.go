packagre main

func main() {
	names := []string{"fish", "cat", "dog"}
	chanMap := make(map[string]chan int)

	for k, name := range names {
		chanMap[name] = make(chan int)
	}
}