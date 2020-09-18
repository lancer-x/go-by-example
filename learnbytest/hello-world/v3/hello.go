package main

func main()  {

}

func Hello(name string) string {
	if name == "" {
		return "hello,world"
	}
	return "hello," + name
}