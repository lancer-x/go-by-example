package main

type CommonError struct {
	msg string
	Offset int64
}

func (e *CommonError) Error() string {
	return e.msg
}

func main()  {

}
